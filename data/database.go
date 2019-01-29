package data

import (
  "fmt"
  "github.com/gin-gonic/gin"
  "github.com/gin-gonic/contrib/sessions"
  "github.com/jinzhu/gorm"
)

type User struct {
  gorm.Model
  Name string
  Email string
  Password string
  Confirm_Password string
  Messages []Message
}

type Message struct {
  gorm.Model
  Text string `json:"name"`
  UserId uint
  User User
}

type Comment struct {
  gorm.Model
  Text string `json:"name"`
  UserId uint `json:"userid"`
  MessageId string `json:"messageid"`
  User User
  Message Message
}

func CreateDB() {
  db, err := gorm.Open("mysql", "root@/chatapp?charset=utf8&parseTime=True&loc=Local")
  if err != nil {
    fmt.Println(err)
  }
  defer db.Close()
  db.AutoMigrate(&Message{}, &User{}, &Comment{})
}

func Create_text(c *gin.Context, text string) {
  db, err := gorm.Open("mysql", "root@/chatapp?charset=utf8&parseTime=True&loc=Local")
  if err != nil {
    fmt.Println(err)
  }
  defer db.Close()
  session := sessions.Default(c)
  user_id := session.Get("userID")
  msg := Message{Text: text, UserId: user_id.(uint)}
  db.Create(&msg)
}

func Create_user(name string, email string, password string, confirm_password string) {
  db, err := gorm.Open("mysql", "root@/chatapp?charset=utf8&parseTime=True&loc=Local")
  if err != nil {
    fmt.Println(err)
  }
  defer db.Close()
  user := User{Name: name,
  Email: email,
  Password: password,
  Confirm_Password: confirm_password,
  }
  db.Create(&user)
}

func Find_user(email string, password string) User{
  db, err := gorm.Open("mysql", "root@/chatapp?charset=utf8&parseTime=True&loc=Local")
  if err != nil {
    fmt.Println(err)
  }
  defer db.Close()
  var user User
  db.Where("email = ? AND password = ?", email, password).Find(&user)
  return user
}

func Get_All_Messages() []Message{
  db, err := gorm.Open("mysql", "root@/chatapp?charset=utf8&parseTime=True&loc=Local")
  if err != nil {
    fmt.Println(err)
  }
  defer db.Close()
  var messages []Message
  var user User
  db.Find(&messages)
  db.Model(&user).Related(&messages)
  db.Preload("User").Find(&messages)
  return messages
}

func Get_Current_User(c *gin.Context) User {
  db, err := gorm.Open("mysql", "root@/chatapp?charset=utf8&parseTime=True&loc=Local")
  if err != nil {
    fmt.Println(err)
  }
  defer db.Close()
  session := sessions.Default(c)
  user_id := session.Get("userID")
  var user User
  var messages []Message
  db.Where("id = ?", user_id).Find(&user)
  db.Model(&user).Related(&messages)
  db.Find(&user).Related(&user.Messages)
  fmt.Println(user)
  return user
}

func Get_Message(id string) Message {
  db, err := gorm.Open("mysql", "root@/chatapp?charset=utf8&parseTime=True&loc=Local")
  if err != nil {
    fmt.Println(err)
  }
  defer db.Close()
  var message Message
  db.First(&message, id).Related(&message.User)
  fmt.Println(message)
  return message
}

func Post_Comment(c *gin.Context ,id string, text string) {
  db, err := gorm.Open("mysql", "root@/chatapp?charset=utf8&parseTime=True&loc=Local")
  if err != nil {
    fmt.Println(err)
  }
  defer db.Close()
  session := sessions.Default(c)
  user_id := session.Get("userID")
  comment := Comment{Text: text, UserId: user_id.(uint), MessageId: id}
  db.Create(&comment)
}
