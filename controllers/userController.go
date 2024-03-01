package controllers

import (
	"finaltask-pbi-btpn/app"
	initializers "finaltask-pbi-btpn/database"
	"finaltask-pbi-btpn/helpers"
	"finaltask-pbi-btpn/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Register(c *gin.Context) {
	// Get Username, Email and password of req body
	var body app.User

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	// Initialize validator
	v := validator.New()

	err := v.Struct(body)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	// hash the password
	hash, err := helpers.GeneratePasswordHash(body.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	// create user
	user := models.User{Username: body.Username, Email: body.Email, Password: string(hash)}

	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})
		return
	}

	// respond
	c.JSON(http.StatusOK, gin.H{
		"message": "Success register user",
		"user":    user})
}

func Login(c *gin.Context) {
	// Get Email and password of req body
	var body app.User

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	// Look up requested user
	var user models.User
	initializers.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	// Compare sent in pass with saved user pass hash
	err := helpers.ComparePasswords([]byte(user.Password), body.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid password",
		})
		return
	}

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := helpers.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid get token",
		})
		return
	}

	// Respond
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("token", tokenString, 3600*24*30, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}

func UpdateUser(c *gin.Context) {

	// Mendapatkan data yang ingin diperbarui dari permintaan HTTP
	var body app.User
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	// Initialize validator
	v := validator.New()

	err := v.Struct(body)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	// Temukan pengguna yang ingin diperbarui dalam database
	var user models.User
	userId := c.Param("id")
	if err := initializers.DB.First(&user, userId).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})
		return
	}

	// Perbarui bidang pengguna dengan data baru
	user.Username = body.Username
	user.Email = body.Email

	// Simpan perubahan ke database
	if err := initializers.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update user",
		})
		return
	}

	// Berikan respons yang sesuai kepada pengguna
	c.JSON(http.StatusOK, gin.H{
		"message": "User updated successfully",
		"user":    user,
	})
}

func DeleteUser(c *gin.Context) {
	// Mendapatkan id pengguna dari parameter rute
	userId := c.Param("id")

	// Temukan pengguna yang ingin dihapus dalam database
	var user models.User
	if err := initializers.DB.First(&user, userId).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})
		return
	}

	// Hapus pengguna dari database
	if err := initializers.DB.Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete user",
		})
		return
	}

	// Berikan respons yang sesuai kepada pengguna
	c.JSON(http.StatusOK, gin.H{
		"message": "User deleted successfully",
	})
}

func FindAllUser(c *gin.Context) {
	// Query untuk mendapatkan semua foto dari database
	var user []models.User
	if err := initializers.DB.Find(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve user",
		})
		return
	}

	// Kirim respons dengan daftar semua foto
	c.JSON(http.StatusOK, gin.H{
		"photos": user,
	})
}

func Validate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Message": "I'm Logged in",
	})
}
