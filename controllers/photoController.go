package controllers

import (
	"finaltask-pbi-btpn/app"
	initializers "finaltask-pbi-btpn/database"
	"finaltask-pbi-btpn/helpers"
	"finaltask-pbi-btpn/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreatePhotoHandler untuk membuat foto baru
func CreatePhoto(c *gin.Context) {
	// Dapatkan informasi pengguna yang saat ini login dari konteks Gin
	currentUser, _ := c.Get("user")
	user := currentUser.(models.User)

	// Buat objek Photo dengan informasi dari body
	var body app.Photo

	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read body: " + err.Error()})
		return
	}

	file, header, err := c.Request.FormFile("photo")
	if err != nil {
		c.JSON(http.StatusBadRequest, "No File Uploads")
		return
	}
	defer file.Close()

	if header.Header.Get("Content-Type") != "image/jpeg" && header.Header.Get("Content-Type") != "image/png" {
		c.JSON(http.StatusBadRequest, "Invalid file type, Only JPG and PNG are allowed")
		return
	}

	if header.Size > (5 << 20) {
		c.JSON(http.StatusBadRequest, "file size too large")
		return
	}

	newNameFile := helpers.NewFileName(header.Filename)

	err = c.SaveUploadedFile(header, "uploads/"+newNameFile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unknown Error " + err.Error()})
		return
	}

	photoUrl := "uploads/" + newNameFile

	var photo models.Photo

	photo.Title = body.Title
	photo.Caption = body.Caption
	photo.PhotoUrl = photoUrl
	photo.UserId = user.ID

	// Buat foto baru
	result := initializers.DB.Create(&photo)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create photo"})
		return
	}

	// Kirim respons setelah foto berhasil dibuat
	c.JSON(http.StatusCreated, gin.H{
		"message": "Success create photo",
		"photo":   photo,
	})
}

func FindAllPhoto(c *gin.Context) {
	// Query untuk mendapatkan semua foto dari database
	var photos []models.Photo
	if err := initializers.DB.Find(&photos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve photos",
		})
		return
	}

	// Kirim respons dengan daftar semua foto
	c.JSON(http.StatusOK, gin.H{
		"photos": photos,
	})
}

func UpdatePhoto(c *gin.Context) {

	currentUser, _ := c.Get("user")
	user := currentUser.(models.User)

	// Temukan pengguna yang ingin diperbarui dalam database
	var photo models.Photo
	photoId := c.Param("id")
	if err := initializers.DB.First(&photo, photoId).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Photo not found",
		})
		return
	}

	// Mendapatkan data yang ingin diperbarui dari permintaan HTTP
	var body app.Photo

	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read body: " + err.Error()})
		return
	}

	file, header, err := c.Request.FormFile("photo")
	if err != nil {
		c.JSON(http.StatusBadRequest, "No File Uploads")
		return
	}
	defer file.Close()

	if header.Header.Get("Content-Type") != "image/jpeg" && header.Header.Get("Content-Type") != "image/png" {
		c.JSON(http.StatusBadRequest, "Invalid file type, Only JPG and PNG are allowed")
		return
	}

	if header.Size > (5 << 20) {
		c.JSON(http.StatusBadRequest, "file size too large")
		return
	}

	newNameFile := helpers.NewFileName(header.Filename)

	err = c.SaveUploadedFile(header, "uploads/"+newNameFile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unknown Error " + err.Error()})
		return
	}

	photoUrl := "uploads/" + newNameFile

	// Perbarui bidang pengguna dengan data baru
	photo.Title = body.Title
	photo.Caption = body.Caption
	photo.PhotoUrl = photoUrl
	photo.UserId = user.ID

	// Simpan perubahan ke database
	if err := initializers.DB.Save(&photo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update photo",
		})
		return
	}

	// Berikan respons yang sesuai kepada pengguna
	c.JSON(http.StatusOK, gin.H{
		"message": "photo updated successfully",
		"photo":   photo,
	})

}

func DeletePhoto(c *gin.Context) {
	// Mendapatkan id pengguna dari parameter rute
	photoId := c.Param("id")

	// Temukan pengguna yang ingin dihapus dalam database
	var photo models.Photo
	if err := initializers.DB.First(&photo, photoId).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Photo not found",
		})
		return
	}

	// Hapus pengguna dari database
	if err := initializers.DB.Delete(&photo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete photo",
		})
		return
	}

	// Berikan respons yang sesuai kepada pengguna
	c.JSON(http.StatusOK, gin.H{
		"message": "photo deleted successfully",
	})
}
