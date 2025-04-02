package db

import (
    "fmt"
    "database/sql"
    _ "github.com/lib/pq"
    "github.com/MafiaLogiki/common/domain"
)

var database *sql.DB
var dbConfig *domain.DBConfig


func init() {
    dbConfig =  &domain.DBConfig {
        Host: "localhost",
        Port: "5432",
        HostName: "postgres",
        Password: "1234",
        DBName: "users",
    }
}

func ConnectToDatabase() (error)  {
    databaseInfo := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v",
        dbConfig.Host,
        dbConfig.Port,
        dbConfig.HostName,
        dbConfig.Password,
        dbConfig.DBName,
    )
    
    var err error
    database, err = sql.Open("postgres", databaseInfo)
    return err
}

func CloseConnection() {
    database.Close()
}

func GetUserToken (username, password string) (string) {
    var token string
    
    row := database.QueryRow("SELECT token FROM users WHERE username = $1 AND password_hash = $2", username, password)
    
    if err := row.Scan(&token); err != nil {
        return ""
    }

    return token
}

func GetIdOfNewUser () (int) {
    var newId int
    err := database.QueryRow(`   
            WITH num_seq AS (
                SELECT ROW_NUMBER () OVER () FROM users
            )
            SELECT num_seq.row_number
            FROM num_seq
            WHERE num_seq.row_number NOT IN (SELECT id FROM users)
            LIMIT 1
        `).Scan(&newId)
    if err == sql.ErrNoRows {
        err2 := database.QueryRow(`
            SELECT MAX(id) FROM users;
        `).Scan(&newId)
        if err2 != nil {
            return 0
        }
        return newId + 1
    }
    return newId
}
