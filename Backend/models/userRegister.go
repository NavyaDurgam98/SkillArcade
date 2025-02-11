package models

type UserRegister struct{
	FirstName string `json:"firstname" binding:"required"`
	LastName string `json:"lastname" binding:"required"`
	Email    string `json:"email" binding:"required"`
	DOB 	 string `json:"dob"`
	Gender	 string `json:"gender"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
}