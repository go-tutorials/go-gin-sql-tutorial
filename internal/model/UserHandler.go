package model

func GetAllUser(user *[]Users) (err error) {
	if err = ConnectDB().Find(user).Error; err != nil {
		return err
	} else {
		return nil
	}
}

func GetByIdUser(user *Users, id string) (err error) {
	if err := ConnectDB().Where("id = ?", id).First(user).Error; err != nil {
		return err
	} else {
		return nil
	}
}

func InsertUser(user *Users) (err error) {
	if err := ConnectDB().Create(user).Error; err != nil {
		return err
	} else {
		return nil
	}
}

func UpdateUser(user *Users, id string) (err error) {
	ConnectDB().Save(user)
	return nil
}

func DeleteUser(user *Users, id string) (err error) {
	ConnectDB().Where("id = ?", id).Delete(user)
	return nil
}
