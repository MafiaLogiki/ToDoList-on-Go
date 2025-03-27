package db

import (
    "ToDoList/models"

    "fmt"
    "database/sql"
    _ "github.com/lib/pq"
)

func ConnectToDatabase() (*sql.DB, error) {
    return sql.Open("postgres", "host=localhost port=5432 user=postgres password=1234 dbname=todolist_database sslmode=disable");
}

func AddTaskToTable(database *sql.DB, Title, Description, Status string, UserId int) (sql.Result, error) {
    return database.Exec("INSERT INTO tasks(title, description, status, user_id) VALUES($1, $2, $3, $4)", Title, Description, Status, UserId);
}

func GetAllTasks(database *sql.DB) ([]task.Task, error) {
    rows, err := database.Query("SELECT * FROM tasks")
    if err != nil {
        return nil, fmt.Errorf("Error in query: %w", err)
    }

    var tasks []task.Task
    for rows.Next() {
        var new_task task.Task
        if err := rows.Scan(&new_task.Id, &new_task.Title, &new_task.Description, &new_task.Status, &new_task.UserId); err != nil {
            return nil, err 
        }
        tasks = append(tasks, new_task)
    }
    return tasks, nil
}
