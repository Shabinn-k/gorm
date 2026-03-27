package controllers

import (
	"golang/config"
	"golang/models"
	"golang/utils"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context){
	var user models.User
	if err:=c.ShouldBindJSON(&user);err!=nil{
		c.JSON(400,gin.H{"error":err.Error()})
		return
	}
	user.Password=utils.HashPassword(user.Password)
	if err:=config.DB.Create(&user).Error;err!=nil{
		c.JSON(500,gin.H{"error":"user creation failed"})
		return 
	}
	c.JSON(201,user)
}

func Login(c *gin.Context){
	var input models.User
	var user models.User

	if err:=c.ShouldBindJSON(&input);err!=nil{
		c.JSON(400,gin.H{"error":err.Error()})
		return
	}
	config.DB.Where("email = ?",input.Email).First(&user)
	if !utils.ComparePassword(user.Password,input.Password){
		c.JSON(401,gin.H{"error":"Invalid Credentials"})
		return
	}
	c.JSON(200,gin.H{"message":"Login successful"})
}

func GetAllUsers(c *gin.Context){
	var users []models.User
	config.DB.Preload("Profile").Find(&users)
	c.JSON(200,users)
}

func GetUser(c *gin.Context){
	id:=c.Param("id")
	var user models.User
	if err:=config.DB.Preload("Profile").First(&user,id).Error;err!=nil{
		c.JSON(404,gin.H{"error":"User Not Found"})
		return 
	}
	c.JSON(200,user)
}

func UpdateUser(c *gin.Context){
	id:=c.Param("id")
var user models.User
if err:=config.DB.First(&user,id).Error;err!=nil{
	c.JSON(404,gin.H{"error":"User Not Found"})
	return
}
if err:=c.ShouldBindJSON(&user);err!=nil{
	c.JSON(400,gin.H{"error":err.Error()})
	return
}
config.DB.Save(&user)
c.JSON(200,user)
}

func DeleteUser(c *gin.Context){
	id:=c.Param("id")
	if err:=config.DB.Delete(&models.User{},id).Error;err!=nil{
		c.JSON(500,gin.H{"error":"Delete failed"})
		return
	}
	c.JSON(200,gin.H{"message":"User Deleted"})
}