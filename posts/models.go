package posts


import (
    "gorm.io/gorm"
    "microblog/database"
    "fmt"
)


type PostModel struct {
    gorm.Model

    Text string  `json:"text"`
    Like int `json:"like"`
    Dislike int `json:"dislike"`

    UserId int `json:"user"`
}

func MakePostModel() PostModel {
    return PostModel{}
}

func (p *PostModel) Migrate() {
    database.RootDatabase.DB.AutoMigrate(&PostModel{})
    fmt.Println("post migrate")

}


func (p *PostModel) Save() error  {
    result := database.RootDatabase.DB.Create(p)

    if result.Error != nil {
        return result.Error
    }

    return nil

}

