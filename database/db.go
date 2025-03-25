package db

import (
    "database/sql"
    _ "github.com/lib/pq"
)

func Connect_to_database() (*sql.DB, error) {
    return sql.Open("postgres", "host=localhost port=5432 user=postgres password=1234 dbname=todolist_database sslmode=disable");
}

func Add_task_to_table(database *sql.DB, task_to_add string) (sql.Result, error) {
    return database.Exec("INSERT INTO tasks(task) VALUES($1) RETURNING task_id", task_to_add);
}
