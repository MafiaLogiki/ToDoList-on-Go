package db

import (
	"database/sql"
	"fmt"
	"task-service/internal/config"

	"github.com/MafiaLogiki/common/domain"

	_ "github.com/lib/pq"
)

var database *sql.DB

func ConnectToDatabase(cfg *config.Config) (error) {
    databaseInfo := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable",
        cfg.Postgres.Host,
        cfg.Postgres.Port,
        cfg.Postgres.HostName,
        cfg.Postgres.Password,
        cfg.Postgres.DBName,
    )
    
    var err error
    database, err = sql.Open("postgres", databaseInfo)

    if database.Ping() != nil {
        return database.Ping()
    }

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
        if err := rows.Scan(&new_task.Id, &new_task.Title, &new_task.Description, &new_task.UserId); err != nil {
            return nil, err 
        }
        tasks = append(tasks, new_task)
    }
    return tasks, nil
}

func GetAllTasksByUserId(userID int) ([]domain.Task, error) {
    rows, err := database.Query("SELECT * FROM tasks WHERE owner_id = $1", userID)
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
