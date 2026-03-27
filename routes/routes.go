package routes

import (
	"github.com/gin-gonic/gin"
	"golang/controllers"
)

func SetRoutes(r *gin.Engine) {
	r.POST("/register",controllers.Register)
	r.POST("/login",controllers.Login)

	r.GET("/users",controllers.GetAllUsers)
	r.GET("/users/:id",controllers.GetUser)

	r.PUT("/users/:id",controllers.UpdateUser)
	r.DELETE("users/:id",controllers.DeleteUser)
}