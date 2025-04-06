CREATE DATABASE users;
CREATE DATABASE todolist_database;

\connect users;

CREATE TABLE users (
    id SERIAL,
    username VARCHAR(30) UNIQUE NOT NULL,
    password_hash VARCHAR(256) NOT NULL,
    PRIMARY KEY(id)
);

INSERT INTO users VALUES(1, 'admin', '8c6976e5b5410415bde908bd4dee15dfb167a9c873fc4bb8a81f6f2ab448a918');

\connect todolist_database;

CREATE TABLE tasks (
    task_id SERIAL,
    title VARCHAR(50) NOT NULL,
    description VARCHAR(200) NOT NULL,
    owner_id INT,
    PRIMARY KEY(task_id)
);

INSERT INTO tasks VALUES(1, 'test title', 'test description', 1);
