package users

import (
    "gorm.io/gorm"
    "microblog/database"
    "fmt"
)

type UserModel struct {
    gorm.Model

    Name string  `json:"name"`
    Password string `json:"password"` // @TODO: Must be hash data !! Build POC for now !!
    Followers int `json:"followers"`

}


func Migrate() {
    database.RootDatabase.DB.AutoMigrate(&UserModel{})
    fmt.Println("users migrate")

}

func (u *UserModel) Save() error  {
    result := database.RootDatabase.DB.Create(u)

    if result.Error != nil {
        return result.Error
    }

    return nil

}
