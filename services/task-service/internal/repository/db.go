package db

import (
    "fmt"
    "database/sql"
    
    "github.com/MafiaLogiki/common/domain"

    _ "github.com/lib/pq"
)

var database *sql.DB

func ConnectToDatabase() (error) {
    var err error
    database, err = sql.Open("postgres", "host=localhost port=5432 user=postgres password=1234 dbname=todolist_database sslmode=disable");
    return err
}

func CloseConnection() {
    database.Close()
}

func AddTaskToTable(Task domain.Task) (sql.Result, error) {
    execQuery := fmt.Sprintf(
        `INSERT INTO tasks(title, description, status, user_id) VALUES(%v, %v, %v, %v)`,
                                                                       Task.Title,
                                                                       Task.Description,
                                                                       Task.Status,
                                                                       Task.UserId)
    return database.Exec(execQuery);
}

func extractTasksFromRows(rows *sql.Rows) ([]domain.Task, error) {
    var tasks []domain.Task
    for rows.Next() {
        var new_task domain.Task
        if err := rows.Scan(&new_task.Id, &new_task.Title, &new_task.Description, &new_task.Status, &new_task.UserId); err != nil {
            return nil, err 
        }
        tasks = append(tasks, new_task)
    }
    return tasks, nil
}

func GetAllTasksByUserId(userID int) ([]domain.Task, error) {
    rows, err := database.Query("SELECT * FROM tasks WHERE user_id = $1", userID)
    if err != nil {
        return nil, fmt.Errorf("Error in query: %w", err)
    }

    return extractTasksFromRows(rows)
}

func GetTaskById(taskID int) (domain.Task, error) {
    row := database.QueryRow("SELECT * FROM tasks WHERE id = $1", taskID)
    var task domain.Task
    if err := row.Scan(&task.Id, &task.Title, &task.Description, &task.Status, &task.UserId); err != nil {
        return task, err 
    }

    return task, nil    
}
