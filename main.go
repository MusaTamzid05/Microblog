package main

import (
    "microblog/database"
    "microblog/users"
    "microblog/posts"
    "github.com/gin-gonic/gin"

)

func init() {
    database.InitDatabase("./test.db")

}


func main() {
    userHandler := users.MakeUserHandle()
    postHandler := posts.MakePostHandler()

    router := gin.Default()
    router.POST("/users/signup", userHandler.SignUpHandler)
    router.POST("/users/login", userHandler.LoginHandler)
    router.GET("/users/", userHandler.GetUsersHandler)


    router.POST("/posts/create", postHandler.CreatePostHandler)
    router.Run()


}
