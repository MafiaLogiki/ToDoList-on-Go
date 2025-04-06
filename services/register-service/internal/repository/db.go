package db

import (
	"database/sql"
	"fmt"
	"register-service/internal/config"

	"github.com/MafiaLogiki/common/domain"

	_ "github.com/lib/pq"
)

var database *sql.DB

func ConnectToDatabase(cfg *config.Config) (error)  {
    databaseInfo := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v", 
                                cfg.Postgres.Host,
                                cfg.Postgres.Port,
                                cfg.Postgres.HostName,
                                cfg.Postgres.Password,
                                cfg.Postgres.DBName)
    
    var err error
    database, err = sql.Open("postgres", databaseInfo)

    return err
}

func CloseConnection() {
    database.Close()
}

func checkIfUsernameValid(username string) (bool) {
    var result_of_searching string
    err := database.QueryRow("SELECT username FROM users WHERE username = $1", username).Scan(&result_of_searching)

    return err != nil
}

func CreateNewUser(newUser domain.User) (int, error) {
    
    if !checkIfUsernameValid(newUser.Username) {
        fmt.Printf("%v", newUser.Username)
        return 0, fmt.Errorf("Error. Username is busy")
    }

    query := fmt.Sprintf("INSERT INTO users(username, password_hash) VALUES('%v', '%v') RETURNING id", 
                        newUser.Username,
                        newUser.Password)
    
    var id int
    err := database.QueryRow(query).Scan(&id)
    
    if err != nil {
        fmt.Printf("testttt\n")
        return 0, err
    }

    return id, nil
}

/*
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

    if err != nil {
        return 0
    }
    return newId
}
*/
