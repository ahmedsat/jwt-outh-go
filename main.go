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
	if user.Username != user.Username || user.Password != user.Password {
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
	if user.Username == user.Username {
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

type Product struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description"`
}

func productsHandler(c *gin.Context) {
	products := []Product{
		Product{100, "BassTune Headset 2.0", 200, "A headphone with a inbuilt high-quality microphone"},
		Product{101, "Fastlane Toy Car", 100, "A toy car that comes with a free HD camera"},
		Product{101, "ATV Gear Mouse", 75, "A high-quality mouse for office work and gaming"},
	}
	c.JSON(200, gin.H{
		"products": products,
	})
}
func main() {
	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/products", productsHandler)
	r.POST("/login", Login)
	r.POST("/register", Register)
	r.Run(":5000")
}
