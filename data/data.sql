create table if not exists users (
    id varchar(40) not null,
    username varchar(120),
    email varchar(120),
    phone varchar(45),
    date_of_birth timestamp with time zone,
    interests varchar[],
    skills json[],
    achievements json[],
    settings json,
    primary key (id)
    );

insert into users (id, username, email, phone, date_of_birth, interests, skills, settings) values ('001','ironman', 'tony.stark@gmail.com', '0987654321', '1963-03-25T16:59:59Z', '{"Play game","Foot ball","Basket ball"}', '[{"skill":"Writing fast","hirable":true}]'', ''{"userid":"001","language":"English","dateFormat":"dd/mm/yyyy","dateTimeFormat":"dd-mm-yyyy:hh:mm","timeFormat":"hh:mm:ss","notification":true}')  RETURNING id;
insert into users (id, username, email, phone, date_of_birth, interests, skills, settings) values ('002','spiderman', 'peter.parker@gmail.com', '0987654321', '1962-08-25T16:59:59Z', '{"Play game","Foot ball","Basket ball"}', '[{"skill":"Writing fast","hirable":true}]'', ''{"userid":"002","language":"English","dateFormat":"dd/mm/yyyy","dateTimeFormat":"dd-mm-yyyy:hh:mm","timeFormat":"hh:mm:ss","notification":true}')  RETURNING id;
insert into users (id, username, email, phone, date_of_birth, interests, skills, settings) values ('003','wolverine', 'james.howlett@gmail.com', '0987654321', '1974-11-16T16:59:59Z', '{"Play game","Foot ball","Basket ball"}', '[{"skill":"Writing fast","hirable":true}]'', ''{"userid":"003","language":"English","dateFormat":"dd/mm/yyyy","dateTimeFormat":"dd-mm-yyyy:hh:mm","timeFormat":"hh:mm:ss","notification":true}')  RETURNING id;
insert into users (id, username, email, phone, date_of_birth, interests, skills, settings) values ('004', 'Radhika', 'radhika_sharma.123@gmail.com', '0900012345', '1974-11-16T16:59:59Z', '{"Play game","Foot ball","Basket ball"}', '[{"skill":"Writing fast","hirable":true}]', '{"userid":"004","language":"English","dateFormat":"dd/mm/yyyy","dateTimeFormat":"dd-mm-yyyy:hh:mm","timeFormat":"hh:mm:ss","notification":true}') RETURNING id;
insert into users (id, username, email, phone, date_of_birth, interests, skills, settings) values ('002', 'Kisiman', 'bob_kisiman_sky@gmail.com', '0900349845', '1982-11-14T16:59:59Z', '{"Tennis","Volley ball","Basket ball"}', '[{"skill":"Writing fast","hirable":true}]', '{"userid":"002","language":"France","dateFormat":"dd-mm-yyyy","dateTimeFormat":"dd-mm-yyyy:hh:mm","timeFormat":"hh:mm","notification":true}') RETURNING id

CREATE INDEX interests_index ON users (interests);
CREATE INDEX skills_index ON users (skills);
CREATE INDEX achievements_index ON users (achievements);
CREATE INDEX settings_index ON users (settings);


