package posts


import (
    "github.com/gin-gonic/gin"
    "net/http"
    "microblog/users"
    "fmt"
)


type PostHandler struct {
    model PostModel
}

type PostLikeRequest struct {
    PostID int `json:"post"`
    Type string  `json:"type"`
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


    err = p.model.Create()

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

func (p *PostHandler) LikeUpdateHandler(c *gin.Context)  {
    var likeUpdateRequest PostLikeRequest

    err := c.BindJSON(&likeUpdateRequest)
    
    if err != nil {
        c.JSON(
            http.StatusBadRequest, gin.H {
                "status" : "failed",
                "message" : err.Error(),
            },
        )

        return
    }

    postID := likeUpdateRequest.PostID
    var postModel PostModel
    targetPost, found :=  postModel.FindByID(postID)

    if !found {

        c.JSON(
            http.StatusNotFound, gin.H {
                "status" : "failed",
                "message" : "post not found",
            },
        )

        return

    }

    fmt.Println("Post ID : ", postID)
    fmt.Println("Type : ", likeUpdateRequest.Type)

    if likeUpdateRequest.Type == "increase" {
        targetPost.Like += 1
    } else {

        if targetPost.Like > 0  {
        targetPost.Like -= 1

        }
    }

    err = targetPost.Update()

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
        http.StatusNotFound, gin.H {
            "status" : "success",
            "message" : "post was successfully updated",
        },
    )

}
