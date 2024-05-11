package posts


import (
    "github.com/gin-gonic/gin"
    "net/http"
    "microblog/users"
)


type PostHandler struct {
    model PostModel
}


func MakePostHandler() PostHandler {
    return PostHandler{}
}

func (p *PostHandler) CreatePostHandler(c *gin.Context) {
    err := c.BindJSON(&p.model)
    
    if err != nil {
        c.JSON(
            http.StatusBadRequest, gin.H {
                "status" : "failed",
                "message" : err.Error(),
            },
        )

        return
    }


    userId := p.model.UserId
    var userModel users.UserModel

    _ , found := userModel.FindByID(userId)

    if found == false {

        c.JSON(
            http.StatusBadRequest, gin.H {
                "status" : "failed",
                "message" : "user id for the post not found",
            },
        )

        return

    }


    err = p.model.Save()

    if err != nil {
        c.JSON(
            http.StatusBadRequest, gin.H {
                "status" : "failed",
                "message" : err.Error(),
            },
        )

        return


    }

    c.JSON(
        http.StatusCreated, gin.H {
            "status" : "success",
        },
    )

}


func (p *PostHandler) GetPostsHandler(c *gin.Context) {
    posts := p.model.GetAll()

    c.JSON(
        http.StatusOK, gin.H {
            "posts" : posts,
            
        },
    )


}
