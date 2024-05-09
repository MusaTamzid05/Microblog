package database

import (
    "gorm.io/gorm"
    "gorm.io/driver/sqlite"
    "fmt"
)

type Database struct {
    Path string
    DB* gorm.DB
}

var RootDatabase Database


func InitDatabase(path string) {
    var err error
    RootDatabase = Database{}
    RootDatabase.Path = path

    RootDatabase.DB, err = gorm.Open(sqlite.Open(path), &gorm.Config{})

    if err != nil {
        errMessage := fmt.Sprintf("Error opening Database => %s\n",err.Error() )
        panic(errMessage)
    }

    fmt.Println("Database init ", path)

}



