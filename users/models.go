package users

import (
    "gorm.io/gorm"
    "microblog/database"
    "fmt"
)

type UserModel struct {
    gorm.Model

    Name string
    Password string
    Followers int

}


func Migrate() {
    database.RootDatabase.DB.AutoMigrate(&UserModel{})
    fmt.Println("users migrate")

}
