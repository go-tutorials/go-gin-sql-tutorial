package services

import (
	"context"
	"database/sql"
	"fmt"
	s "github.com/core-go/sql"
	"github.com/lib/pq"
	"reflect"
	"strings"

	. "go-service/internal/models"
)

type SqlUserService struct {
	DB *sql.DB
}

func NewUserService(db *sql.DB) *SqlUserService {
	return &SqlUserService{DB: db}
}

func (m *SqlUserService) GetAll(ctx context.Context) (*[]User, error) {
	query := "select id, username, email, phone, date_of_birth, interests, skills, achievements, settings from users"
	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	var result []User
	for rows.Next() {
		var user User
		err = rows.Scan(&user.Id, &user.Username, &user.Phone, &user.Email, &user.DateOfBirth, pq.Array(&user.Interests), pq.Array(&user.Skills), pq.Array(&user.Achievements), &user.Settings)
		result = append(result, user)
	}
	return &result, nil
}

func (m *SqlUserService) Load(ctx context.Context, id string) (*User, error) {
	var user User
	query := "select id, username, email, phone, date_of_birth, interests, skills, achievements, settings from users where id = $1"
	err := m.DB.QueryRowContext(ctx, query, id).Scan(&user.Id, &user.Username, &user.Phone, &user.Email, &user.DateOfBirth, pq.Array(&user.Interests), pq.Array(&user.Skills), pq.Array(&user.Achievements), &user.Settings)
	if err != nil {
		errMsg := err.Error()
		if strings.Compare(fmt.Sprintf(errMsg), "0 row(s) returned") == 0 {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return &user, nil
}

func (m *SqlUserService) Insert(ctx context.Context, user *User) (int64, error) {
	query := "insert into users (id, username, email, phone, date_of_birth, interests, skills, achievements, settings) values ($1, $2, $3, $4, $5, $6, $7, $8, $9)"
	stmt, er0 := m.DB.Prepare(query)
	if er0 != nil {
		return -1, nil
	}
	result, er1 := stmt.ExecContext(ctx, user.Id, user.Username, user.Email, user.Phone, user.DateOfBirth, pq.Array(user.Interests), pq.Array(user.Skills), pq.Array(user.Achievements), user.Settings)
	if er1 != nil {
		return -1, nil
	}
	return result.RowsAffected()
}

func (m *SqlUserService) Update(ctx context.Context, user *User) (int64, error) {
	query := "update users set username = $2, email = $3, phone = $4, date_of_birth = $5, interests = $6, skills = $7, achievements = $8, settings = $9 where id = $1"
	stmt, er0 := m.DB.Prepare(query)
	if er0 != nil {
		return -1, nil
	}
	result, er1 := stmt.ExecContext(ctx, user.Id, user.Username, user.Email, user.Phone, user.DateOfBirth, pq.Array(user.Interests), pq.Array(user.Skills), pq.Array(user.Achievements), user.Settings)
	if er1 != nil {
		return -1, er1
	}
	return result.RowsAffected()
}

func (m *SqlUserService) Patch(ctx context.Context, user map[string]interface{}) (int64, error) {
	userType := reflect.TypeOf(User{})
	result, err := s.Patch(ctx, m.DB, "users", user, userType)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (m *SqlUserService) Delete(ctx context.Context, id string) (int64, error) {
	query := "delete from users where id = $1"
	stmt, er0 := m.DB.Prepare(query)
	if er0 != nil {
		return -1, nil
	}
	result, er1 := stmt.ExecContext(ctx, id)
	if er1 != nil {
		return -1, er1
	}
	rowAffect, er2 := result.RowsAffected()
	if er2 != nil {
		return 0, er2
	}
	return rowAffect, nil
}

func (m *SqlUserService) Search(ctx context.Context, user User) (*[]User, error) {
	query := "select * from users where"
	var whereArray []string
	var agrs []interface{}
	var orWhereSkill []string
	var orWhereAchievements []string
	i := 1
	if user.Interests != nil && len(user.Interests) > 0 {
		whereArray = append(whereArray, fmt.Sprintf(` interests && $%d`, i))
		agrs = append(agrs, pq.Array(user.Interests))
		i++
		fmt.Println("Interests: ", user.Interests)
	}

	//select * from users where settings -> 'language' ? 'Spanish';
	//SELECT * FROM users WHERE settings @> '{"language":"Spanish"}';

	if user.Settings != nil && len(user.Settings.Language) > 0 {
		whereArray = append(whereArray, fmt.Sprintf(` settings -> 'language' ? $%d `, i))
		agrs = append(agrs, user.Settings.Language)
		i++
		fmt.Println("Language:", user.Settings.Language)
	}

	if user.Skills != nil && len(user.Skills) > 0 {
		for j := 0; j < len(user.Skills); j++ {
			fmt.Println("len: ", len(user.Skills))
			agrs = append(agrs, user.Skills[j])
			orWhereSkill = append(orWhereSkill, fmt.Sprintf(` $%d <@ ANY(skills)`, i))
			i++
			fmt.Println("Skills:", user.Skills)
		}
	}

	if user.Achievements != nil && len(user.Achievements) > 0 {
		for j := 0; j < len(user.Achievements); j++ {
			fmt.Println("len: ", len(user.Achievements))
			agrs = append(agrs, user.Achievements[j])
			orWhereAchievements = append(orWhereAchievements, fmt.Sprintf(` $%d <@ ANY(achievements) `, i))
			i++
			fmt.Println("Achievements:", user.Achievements)
		}
	}
	if len(whereArray) > 0 {
		query = query + strings.Join(whereArray, " and")
	}

	if len(orWhereSkill) > 0 {
		query = query + `and` + strings.Join(orWhereSkill, " or")
	}

	if len(orWhereAchievements) > 0 {
		query = query + `and` + strings.Join(orWhereAchievements, " or")
	}

	fmt.Println(query)
	rows, err := m.DB.QueryContext(ctx, query, agrs...)
	if err != nil {
		return nil, err
	}

	var result []User
	for rows.Next() {
		var user User
		err = rows.Scan(&user.Id, &user.Username, &user.Phone, &user.Email, &user.DateOfBirth, pq.Array(&user.Interests), &user.Settings, pq.Array(&user.Skills), pq.Array(&user.Achievements))
		result = append(result, user)
	}
	return &result, nil
}
