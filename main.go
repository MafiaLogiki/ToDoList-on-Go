package main

import (
    "fmt"
    "flag"

    "ToDoList/database"
    "ToDoList/handlers"
    _ "ToDoList/models"
)

const (
        ToPerform = "To perform"
        InProcess = "In process"
        Done = "Done"
)

func main () {
    database, err := db.ConnectToDatabase();
    
    if err != nil {
        fmt.Println("Error:", err)
        return
    } 
    defer database.Close();
    

    _, insert_error := db.AddTaskToTable(database, "Clean home", "Need to wash the dish", "To perform", 1);
    
    if insert_error != nil {
        fmt.Println("Error in insertion:", insert_error);
    } else {
        fmt.Println("Insert complete successfully");
    }


    addr := flag.String("addr", ":8080", "localhost")

    TaskHandler.CreateAndRunServer(database, *addr)
}
