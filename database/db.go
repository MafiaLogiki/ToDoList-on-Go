package db

import (
    "ToDoList/models/task"

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

func extractTasksFromRows(rows *sql.Rows) ([]task.Task, error) {
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

func GetAllTasks(database *sql.DB) ([]task.Task, error) {
    rows, err := database.Query("SELECT * FROM tasks")
    if err != nil {
        return nil, fmt.Errorf("Error in query: %w", err)
    }
    return extractTasksFromRows(rows)
}

func GetAllTasksByUserId(database *sql.DB, userID int) ([]task.Task, error) {
    rows, err := database.Query("SELECT * FROM tasks WHERE user_id = $1", userID)
    if err != nil {
        return nil, fmt.Errorf("Error in query: %w", err)
    }

    return extractTasksFromRows(rows)
}

func GetTaskById(database *sql.DB, taskID int) (task.Task, error) {
    row := database.QueryRow("SELECT * FROM tasks WHERE id = $1", taskID)
    var task task.Task
    if err := row.Scan(&task.Id, &task.Title, &task.Description, &task.Status, &task.UserId); err != nil {
        return task, err 
    }

    return task, nil    
}

func GetUserToken (database *sql.DB, username, password string) (string) {
    var token string
    
    row := database.QueryRow("SELECT token FROM users WHERE username = $1 AND password_hash = $2", username, password)
    
    if err := row.Scan(&token); err != nil {
        return ""
    }

    return token
}
