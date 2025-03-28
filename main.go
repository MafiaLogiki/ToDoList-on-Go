package main

import (
    "fmt"
    "flag"

    "ToDoList/database"
    "ToDoList/handlers"
    "ToDoList/middleware"
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
    
    token, _ := auth.CreateToken("test")
    fmt.Printf("Token for user test: %v\n", token)

    TaskHandler.CreateAndRunServer(database, *addr)
}
