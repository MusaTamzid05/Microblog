package main

import (
    "fmt"
    "microblog/database"
    "microblog/users"
)

func init() {
    database.InitDatabase("./test.db")

}


func main() {
    users.Migrate()
    fmt.Println("Entry point!!")
}
