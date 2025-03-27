package main

import (
    "fmt"
    "ToDoList/database"
    _ "ToDoList/models"
)

const (
        to_perform = "To perform"
        in_process = "In process"
        done = "Done"
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

    all_tasks, err2 := db.GetAllTasks(database)
    if err2 != nil {
        fmt.Println("Error in selection:", err2)
        return
    }

    for i := 0; i < len(all_tasks); i++ {
        fmt.Printf("%v %v %v %v %v\n", all_tasks[i].Id, all_tasks[i].Title, all_tasks[i].Description, all_tasks[i].Status, all_tasks[i].UserId)
    }
}
