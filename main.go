package main

import(
  "fmt"
  "./data"
  "./controller"
  "github.com/gin-gonic/gin"
  "github.com/jinzhu/gorm"
  "github.com/gin-gonic/contrib/sessions"
  _"github.com/go-sql-driver/mysql"
)

func create_text(text string) {
  db, err := gorm.Open("mysql", "root@/chatapp?charset=utf8&parseTime=True&loc=Local")
  if err != nil {
    fmt.Println(err)
  }
  defer db.Close()
  msg := data.Message{Text: text}
  db.Create(&msg)
}

func get_all() []data.Message{
  db, err := gorm.Open("mysql", "root@/chatapp?charset=utf8&parseTime=True&loc=Local")
  if err != nil {
    fmt.Println(err)
  }
  defer db.Close()
  var messages []data.Message
  db.Find(&messages)
  return messages
}

func main() {
  data.CreateDB()
  server := gin.Default()

  store := sessions.NewCookieStore([]byte("secret"))
  server.Use(sessions.Sessions("SessionName", store))

  server.LoadHTMLGlob("templates/*")
  server.GET("/", controller.IndexRouter)
  server.POST("/post", controller.PostRouter)
  server.GET("/signup", controller.SignupRouter)
  server.POST("/user/signup", controller.UserSignupRouter)
  server.GET("signin", controller.SigninRouter)
  server.POST("/user/signin", controller.UserSigninRouter)
  server.GET("/user/mypage", controller.UserMyPageRouter)
  server.Run()
}
