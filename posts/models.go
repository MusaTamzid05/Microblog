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


func (p *PostModel) Create() error  {
    result := database.RootDatabase.DB.Create(p)

    if result.Error != nil {
        return result.Error
    }

    return nil

}


func (p *PostModel) Update() error  {
    result := database.RootDatabase.DB.Save(p)

    if result.Error != nil {
        return result.Error
    }

    return nil

}





func (p *PostModel) GetAll() []PostModel {
    posts := []PostModel{}
    database.RootDatabase.DB.Find(&posts)

    return posts 

}


func (p PostModel) FindByID(id int) (PostModel, bool) {
    var postModel PostModel
    result := database.RootDatabase.DB.First(&postModel, id)

    if result.Error != nil {
        if result.Error == gorm.ErrRecordNotFound {
            return postModel, false 

        }

        // @TODO: Do not panic here !, send the error message
        panic(result.Error.Error())
    }


    return postModel, true


}


type CommentModel struct {
    gorm.Model

    Text string  `json:"text"`
    Like int `json:"like"`
    Dislike int `json:"dislike"`

    CommentID int `json:"comment"`
}

func MakeCommentModel() CommentModel {
    return CommentModel{}
}

func (c *CommentModel) Migrate() {
    database.RootDatabase.DB.AutoMigrate(&CommentModel{})
    fmt.Println("comment migrate")

}
