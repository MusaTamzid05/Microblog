package main

import (
    "microblog/database"
    "microblog/users"
    "github.com/gin-gonic/gin"

)

func init() {
    database.InitDatabase("./test.db")

}


func main() {
    userHandler := users.MakeUserHandle()

    router := gin.Default()
    router.POST("/users/signup", userHandler.SignUpHandler)
    router.GET("/users/", userHandler.GetUsersHandler)

    router.Run()


}
