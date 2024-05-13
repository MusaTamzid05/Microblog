package users

import (
    "github.com/gin-gonic/gin"
    "net/http"
)


type UserHandler struct {
}

func MakeUserHandle() UserHandler {
    return UserHandler{}
}


func (u *UserHandler) SignUpHandler(c *gin.Context) {
    var user UserModel
    err := c.BindJSON(&user)
    
    if err != nil {
        c.JSON(
            http.StatusBadRequest, gin.H {
                "status" : "failed",
                "message" : err.Error(),
            },
        )

        return
    }

    if user.NameExists()  {
        c.JSON(
            http.StatusBadRequest, gin.H {
                "status" : "failed",
                "message" : "user with the same name exists!!",
            },
        )


        return
    }

    err = user.Save()

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



func (u *UserHandler) GetUsersHandler(c *gin.Context) {
    var user UserModel
    c.JSON(
        http.StatusOK, gin.H {
            "users" : user.GetAll(),
        },
    )

}



func (u *UserHandler) LoginHandler(c *gin.Context) {
    var user UserModel
    err := c.BindJSON(&user)
    
    if err != nil {
        c.JSON(
            http.StatusBadRequest, gin.H {
                "status" : "failed",
                "message" : err.Error(),
            },
        )

        return
    }

    if user.Exists()  {
        c.JSON(
            http.StatusOK, gin.H {
                "status" : "success",
            },
        )


        return
    }


    c.JSON(
        http.StatusBadRequest, gin.H {
            "status" : "failed",
        },
    )


}
