package users

import (
    "github.com/gin-gonic/gin"
    "net/http"
)


type UserHandler struct {
    model UserModel
}

func MakeUserHandle() UserHandler {
    return UserHandler{model: UserModel{}}
}


func (u *UserHandler) SignUpHandler(c *gin.Context) {
    err := c.BindJSON(&u.model)
    
    if err != nil {
        c.JSON(
            http.StatusBadRequest, gin.H {
                "status" : "failed",
                "message" : err.Error(),
            },
        )

        return
    }

    err = u.model.Save()

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
        http.StatusBadRequest, gin.H {
            "status" : "success",
            "message" : "users created",
        },
    )

}



