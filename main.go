package main

import(
  "./data"
  "./controller"
  "github.com/gin-gonic/gin"
  "github.com/gin-gonic/contrib/sessions"
  _"github.com/go-sql-driver/mysql"
)

func main() {
  data.CreateDB()
  server := gin.Default()

  store := sessions.NewCookieStore([]byte("secret"))
  server.Use(sessions.Sessions("SessionName", store))

  server.LoadHTMLGlob("templates/*")
  server.Static("/assets", "./assets")
  server.Static("/css","../assets/*")
  server.GET("/", controller.IndexRouter)
  server.POST("/post", controller.PostRouter)
  server.GET("/signup", controller.SignupRouter)
  server.POST("/user/signup", controller.UserSignupRouter)
  server.GET("signin", controller.SigninRouter)
  server.POST("/user/signin", controller.UserSigninRouter)
  server.GET("/user/mypage", controller.UserMyPageRouter)
  server.GET("/message/:id", controller.MessageShow)
  server.POST("/message/:id/comments", controller.PostCommentRouter)
  server.Run()
}
