package UserRegistration

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"go.mongodb.org/mongo-driver/bson"
	"SKILLARCADE/DatabaseConnection"
)

// UserLogin function will handles user login
func UserLogin(c *gin.Context) {
	// Extract username & password from JSON request body
	var loginData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	// Access MongoDB collection
	userDetailsCollection := DatabaseConnection.GetCollection("SkillArcade","UserDetails")

	// check if user exists in DB
	var userExists bson.M
	err := userDetailsCollection.FindOne(c, bson.M{"username": loginData.Username}).Decode(&userExists)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Email not found"})
		return
	}

	// Check if the passwords match
	if userExists["password"] != loginData.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}

	// Successful login
	c.JSON(http.StatusOK, gin.H{"message": "Login successful!"})

}

func RegisterUser(c *gin.Context) {
	// Extract username, email, and password from JSON request body
	var userData struct {
		FirstName string `json:"firstname"`
		LastName string `json:"lastname"`
		Email    string `json:"email"`
		DOB 	 string `json:"dob"`
		Gender	 string `json:"gender"`
		Username string `json:"username"`
		Password string `json:"password"`
	}

	// Bind JSON request to userData struct
	if err := c.ShouldBindJSON(&userData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	userDetailsCollection := DatabaseConnection.GetCollection("SkillArcade","UserDetails")
	//Check if email already exists in DB
	var existingUser bson.M
	err := userDetailsCollection.FindOne(c, bson.M{"email": userData.Email}).Decode(&existingUser)
	if err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Email already exists"})
		return
	}

	//create user object to insert in DB
	user := bson.M{
		"firstname": userData.FirstName,
		"lastname": userData.LastName,
		"username": userData.Username,
		"password": userData.Password,
		"email":    userData.Email,
		"dob": userData.DOB,
		"gender": userData.Gender,
	}

	// Insert user into the MongoDB collection
	_, err = userDetailsCollection.InsertOne(c, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating user"})
		return
	}

	// Success response
	c.JSON(http.StatusCreated, gin.H{
		"message":  "User registered successfully!",
		"firstname": userData.FirstName,
		"lastname": userData.LastName,
		"username": userData.Username,
		"email":    userData.Email,
		"dob": userData.DOB,
		"gender": userData.Gender,
	})
}

// RegisterUserRoutes sets up user-related routes
func RegisterAndLoginUserRoutes(router *gin.Engine) {
	router.POST("/signin", UserLogin)
	router.POST("/signup",RegisterUser)
}
