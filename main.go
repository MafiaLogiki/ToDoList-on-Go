package main

import (
    "fmt"
    "ToDoList/database"
)

func main () {
    database, err := db.Connect_to_database();
    if err != nil {
        fmt.Println("Error:", err);
    } else {
        fmt.Println("Ok.");
        
        _, insert_error := db.Add_task_to_table(database, "clean home");
        if insert_error != nil {
            fmt.Println("Error:", insert_error);
        } else {
            fmt.Println("Insert complete successfully");
        }

        database.Close();
    }
}
