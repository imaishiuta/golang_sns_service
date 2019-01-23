package controller

import (
  "github.com/gin-gonic/gin"
  "../data"
  //"fmt"
  "../session"
)

func IndexRouter(c *gin.Context) {
  messages := data.Get_All_Messages()
  c.HTML(200, "index.html", gin.H{
    "messages": messages,
    })
}

func PostRouter(c *gin.Context) {
  text := c.PostForm("text")
  data.Create_text(c, text)
  c.JSON(200, gin.H{
    "Text": text,
  })
}

func SignupRouter(c *gin.Context){
  c.HTML(200, "signup.html", nil)
}

func UserSignupRouter(c *gin.Context) {
  name := c.PostForm("name")
  email := c.PostForm("email")
  password := c.PostForm("password")
  confirm_password := c.PostForm("confirm_password")
  data.Create_user(name, email, password, confirm_password)
}

func SigninRouter(c *gin.Context) {
  c.HTML(200, "signin.html", nil)
}

func UserSigninRouter(c *gin.Context) {
  email := c.PostForm("email")
  password := c.PostForm("password")
  user := data.Find_user(email, password)
  session.Login(c, user)
  messages := data.Get_All_Messages()
  c.HTML(200, "index.html", gin.H{
    "messages": messages,
    })
}

func UserMyPageRouter(c *gin.Context) {
  user := data.Get_Current_User(c)
  c.HTML(200, "mypage.html", gin.H{
    "User": user,
    })
}
