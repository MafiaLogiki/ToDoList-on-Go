package main

import (
    "fmt"
    "flag"

    _ "github.com/swaggo/files"         // Swagger UI files
    "github.com/swaggo/http-swagger"    // Swagger middleware
    "ToDoList/database"
    "ToDoList/handlers"
)

const (
        ToPerform = "To perform"
        InProcess = "In process"
        Done = "Done"
)

// 65e84be33532fb784c48129675f9eff3a682b27168c0ea744b2cf58ee02337c5 - sha256 for qwerty

func main () {
    database, err := db.ConnectToDatabase();
    
    if err != nil {
        fmt.Println("Error:", err)
        return
    } 
    defer database.Close(); 
    addr := flag.String("addr", ":8080", "localhost")
    
    fmt.Printf("%v\n", db.GetIdOfNewUser(database))

    TaskHandler.CreateAndRunServer(database, *addr)
}
