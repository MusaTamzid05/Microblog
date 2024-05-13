package posts


import (
    "github.com/gin-gonic/gin"
    "net/http"
    "microblog/users"
    "fmt"
)


type PostHandler struct {
}

type PostLikeRequest struct {
    PostID int `json:"post"`
    Type string  `json:"type"`
}


func MakePostHandler() PostHandler {
    return PostHandler{}
}

func (p *PostHandler) CreatePostHandler(c *gin.Context) {
    var post PostModel
    err := c.BindJSON(&post)
    
    if err != nil {
        c.JSON(
            http.StatusBadRequest, gin.H {
                "status" : "failed",
                "message" : err.Error(),
            },
        )

        return
    }


    userId := post.UserId
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


    err = post.Create()

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
    var post PostModel
    posts := post.GetAll()

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


type CommentHandler struct {
}


func MakeCommentHandler() CommentHandler {
    return CommentHandler{}
}

func (co *CommentHandler) CreateCommentHandler(c *gin.Context) {
    var comment CommentModel
    err := c.BindJSON(&comment)
    
    if err != nil {
        c.JSON(
            http.StatusBadRequest, gin.H {
                "status" : "failed",
                "message" : err.Error(),
            },
        )

        return
    }

    post := MakePostModel()

    _, postFound := post.FindByID(comment.PostID)

    if !postFound {
        c.JSON(
            http.StatusNotFound, gin.H {
                "status" : "failed",
                "message" : "comment parent post not found",
            },
        )

        return

    }

    err = comment.Create()

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


func (co *CommentHandler) GetCommentsHandler(c *gin.Context) {
    var comment CommentModel
    comments := comment.GetAll()

    c.JSON(
        http.StatusOK, gin.H {
            "comments" : comments,
            
        },
    )
}




