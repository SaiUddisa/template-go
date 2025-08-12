package controllers

import (
	"net/http"
	"strconv"
	"template-go/models"
	"template-go/services"

	"github.com/gin-gonic/gin"
)

// Set up the user routes
func SetupUserRoutes(r *gin.Engine, service services.UserService) {
	r.GET("/users", getUsers(service))
	r.GET("/users/:id", getUser(service))
	r.POST("/users", createUser(service))
	r.PUT("/users/:id", updateUser(service))
	r.DELETE("/users/:id", deleteUser(service))
}

func getUsers(service services.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		limitParam := c.DefaultQuery("limit", "10") // Default limit is 10
		limit, err := strconv.Atoi(limitParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid limit"})
			return
		}

		users, err := service.GetUsers(limit)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error fetching users"})
			return
		}
		c.JSON(http.StatusOK, users)
	}
}

func getUser(service services.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
			return
		}
		user, err := service.GetUserByID(uint(id))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
			return
		}
		c.JSON(http.StatusOK, user)
	}
}

func createUser(service services.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.Users
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input"})
			return
		}
		createdUser, err := service.CreateUser(user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error creating user"})
			return
		}
		c.JSON(http.StatusCreated, createdUser)
	}
}

func updateUser(service services.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
			return
		}
		var user models.Users
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input"})
			return
		}
		user.Id = uint(id)
		updatedUser, err := service.UpdateUser(user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error updating user"})
			return
		}
		c.JSON(http.StatusOK, updatedUser)
	}
}

func deleteUser(service services.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
			return
		}
		err = service.DeleteUser(uint(id))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
			return
		}
		c.JSON(http.StatusNoContent, nil)
	}
}
