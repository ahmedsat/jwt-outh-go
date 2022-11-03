package main

import (
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type User struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// Mock data
var user = User{
	ID:       3,
	Username: "test",
	Password: "pass",
}

// login
func Login(c *gin.Context) {
	var u User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}
	if u.Username != user.Username || u.Password != user.Password {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "Invalid Credentials",
		})
		return
	}
	t, err := generateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Logged in successfully",
		"token":   t,
	})

}

// generate token
func generateToken(userId uint64) (string, error) {
	var err error

	os.Setenv("ACCESS_SECRET", "tushsshss") //this should be in an env file
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userId
	atClaims["exp"] = time.Now().Add(time.Minute * 60).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	t, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return t, nil
}

// register token
func Register(c *gin.Context) {
	var u User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}
	if u.Username == user.Username {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "User already exists",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "User registered successfully",
	})
}

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	r.POST("/login", Login)
	r.POST("/register", Register)
	r.Run(":5000")
}
