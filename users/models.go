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

func (u UserModel) GetAll() []UserModel {
    userModels := []UserModel{}
    database.RootDatabase.DB.Find(&userModels)

    return userModels


}


func (u UserModel) NameExists() bool {
    var tempUser UserModel
    result := database.RootDatabase.DB.Where("name=?", u.Name).First(&tempUser)

    if result.Error != nil {
        if result.Error == gorm.ErrRecordNotFound {
            return false

        }

        // @TODO: Do not panic here !, send the error message
        panic(result.Error.Error())
    }


    return true


}

func (u UserModel) Exists() bool {
    var tempUser UserModel
    result := database.RootDatabase.DB.Where("name=? AND password=?", u.Name, u.Password).First(&tempUser)

    if result.Error != nil {
        if result.Error == gorm.ErrRecordNotFound {
            return false

        }

        // @TODO: Do not panic here !, send the error message
        panic(result.Error.Error())
    }


    return true


}
