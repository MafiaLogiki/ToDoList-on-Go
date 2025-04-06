package db

import (
	"database/sql"
	"fmt"

	"auth-service/internal/config"
    "github.com/MafiaLogiki/common/domain"

	_ "github.com/lib/pq"
)

var database *sql.DB

func ConnectToDatabase(cfg *config.Config) (error)  {
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

func GetIdByUserData (user domain.User) int {
    var id int
    err := database.QueryRow("SELECT id FROM users WHERE username = $1 AND password_hash = $2", user.Username, user.Password).Scan(&id)
    if err != nil {
        return 0
    }
    return id
}
