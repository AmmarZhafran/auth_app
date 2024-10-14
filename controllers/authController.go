// package controllers

// import (
// 	"auth-app/models"
// 	"auth-app/utils"
// 	"errors"
// 	"log"
// 	"net/http"
//
//
//
//
//
//
//
//
//
//
//
//
//
//

// 	"github.com/gin-gonic/gin"
// 	"golang.org/x/crypto/bcrypt"
// 	"gorm.io/gorm"

// )

// func Register(db *gorm.DB) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var user models.User
// 		if err := c.ShouldBindJSON(&user); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		// Cek apakah email sudah ada
// 		var existingUser models.User
// 		err := db.Where("email = ?", user.Email).First(&existingUser).Error
// 		if err != nil && err != gorm.ErrRecordNotFound {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to check existing user", "details": err.Error()})
// 			return
// 		}

// 		if err == nil {
// 			c.JSON(http.StatusConflict, gin.H{"error": "Email already exists"})
// 			return
// 		}

// 		// Hash password
// 		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
// 			return
// 		}
// 		user.PasswordHash = string(hashedPassword)
// 		user.OTP = utils.GenerateOTP()

// 		// Kirim OTP terlebih dahulu
// 		if err := utils.SendOTP(user.Email, user.OTP); err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to send OTP"})
// 			return
// 		}

// 		// Jika OTP berhasil dikirim, simpan user ke database
// 		user.IsVerified = false // Set IsVerified ke false sebelum menyimpan
// 		if err := db.Create(&user).Error; err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create user"})
// 			return
// 		}

// 		c.JSON(http.StatusOK, gin.H{"message": "User created, please verify your email"})
// 	}
// }

// // func Register(db *gorm.DB) gin.HandlerFunc {
// // 	return func(c *gin.Context) {
// // 		var user models.User
// // 		if err := c.ShouldBindJSON(&user); err != nil {
// // 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// // 			return
// // 		}

// // 		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
// // 		user.PasswordHash = string(hashedPassword)
// // 		user.OTP = utils.GenerateOTP()

// // 		if err := db.Create(&user).Error; err != nil {
// // 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create user"})
// // 			return
// // 		}

// // 		if err := utils.SendOTP(user.Email, user.OTP); err != nil {
// // 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to send OTP"})
// // 			return
// // 		}

// // 		c.JSON(http.StatusOK, gin.H{"message": "User created, please verify your email"})
// // 	}
// // }

// func Login(db *gorm.DB) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var user models.User
// 		var input models.User

// 		if err := c.ShouldBindJSON(&input); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		if err := db.Where("email = ?", input.Email).First(&user).Error; err != nil {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
// 			return
// 		}

// 		if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(input.PasswordHash)); err != nil {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
// 			return
// 		}

// 		token, _ := utils.GenerateToken(user.Email)
// 		c.JSON(http.StatusOK, gin.H{"token": token})
// 	}
// }

// // func VerifyOTP(db *gorm.DB) gin.HandlerFunc {
// //     return func(c *gin.Context) {
// //         var user models.User
// //         var input struct {
// //             Email string `json:"email"`
// //             OTP   string `json:"otp"`
// //         }

// //         if err := c.ShouldBindJSON(&input); err != nil {
// //             c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// //             return
// //         }

// //         // Mencari pengguna berdasarkan email
// //         if err := db.Where("email = ?", input.Email).First(&user).Error; err != nil {
// //             if errors.Is(err, gorm.ErrRecordNotFound) {
// //                 c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
// //             } else {
// //                 c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
// //             }
// //             return
// //         }

// //         // Memeriksa apakah OTP valid
// //         if user.OTP != input.OTP {
// //             c.JSON(http.StatusForbidden, gin.H{"error": "Invalid OTP"})
// //             return
// //         }

// //         // Mengatur user sebagai terverifikasi dan menghapus OTP
// //         user.IsVerified = true
// //         user.OTP = ""
// //         if err := db.Save(&user).Error; err != nil {
// //             c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user verification status"})
// //             return
// //         }

// //         c.JSON(http.StatusOK, gin.H{"message": "User verified successfully"})
// //     }
// // }

// func VerifyOTP(db *gorm.DB) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var user models.User
// 		var input struct {
// 			Email string `json:"email"`
// 			OTP   string `json:"otp"`
// 		}

// 		// Mengambil data input
// 		if err := c.ShouldBindJSON(&input); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		// Mencari pengguna berdasarkan email
// 		if err := db.Where("email = ?", input.Email).First(&user).Error; err != nil {
// 			if errors.Is(err, gorm.ErrRecordNotFound) {
// 				c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
// 			} else {
// 				c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
// 			}
// 			return
// 		}

// 		// Log untuk debugging
// 		log.Printf("Verifying OTP for user %s: input OTP: %s, stored OTP: %s", input.Email, input.OTP, user.OTP)

// 		// Memeriksa apakah OTP valid
// 		if user.OTP != input.OTP {
// 			c.JSON(http.StatusForbidden, gin.H{"error": "Invalid OTP"})
// 			return
// 		}

// 		// Jika OTP valid, atur pengguna sebagai terverifikasi
// 		user.IsVerified = true
// 		user.OTP = "" // Reset OTP
// 		if err := db.Save(&user).Error; err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user verification status"})
// 			return
// 		}

//			c.JSON(http.StatusOK, gin.H{"message": "User verified successfully"})
//		}
//	}
package controllers

import (
	"auth-app/models"
	"auth-app/utils"
	_ "errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var tempUsers = make(map[string]models.User) // Temporary storage for users

func Login(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		var input models.User

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.Where("email = ?", input.Email).First(&user).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(input.PasswordHash)); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}

		token, _ := utils.GenerateToken(user.Email)
		c.JSON(http.StatusOK, gin.H{"token": token})
	}
}

// Register handles user registration and sends OTP
func Register(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Cek apakah email sudah ada
		var existingUser models.User
		err := db.Where("email = ?", user.Email).First(&existingUser).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to check existing user", "details": err.Error()})
			return
		}

		if err == nil {
			c.JSON(http.StatusConflict, gin.H{"error": "Email already exists"})
			return
		}

		// Hash password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
			return
		}

		user.PasswordHash = string(hashedPassword)
		user.OTP = utils.GenerateOTP()

		// Simpan pengguna sementara
		tempUsers[user.Email] = user

		// Kirim OTP
		if err := utils.SendOTP(user.Email, user.OTP); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to send OTP"})
			delete(tempUsers, user.Email) // Remove temporary user on failure
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "OTP sent, please verify your email"})
	}
}

// VerifyOTP verifies the OTP and saves the user if successful
func VerifyOTP(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input struct {
			Email string `json:"email"`
			OTP   string `json:"otp"`
		}

		// Mengambil data input
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Mencari pengguna sementara berdasarkan email
		user, exists := tempUsers[input.Email]
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found or OTP expired"})
			return
		}

		// Log untuk debugging
		log.Printf("Verifying OTP for user %s: input OTP: %s, stored OTP: %s", input.Email, input.OTP, user.OTP)

		// Memeriksa apakah OTP valid
		if user.OTP != input.OTP {
			c.JSON(http.StatusForbidden, gin.H{"error": "Invalid OTP"})
			return
		}

		// Jika OTP valid, simpan pengguna ke dalam database
		user.IsVerified = true
		user.OTP = "" // Reset OTP

		if err := db.Create(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save user"})
			return
		}

		// Hapus pengguna sementara setelah berhasil disimpan
		delete(tempUsers, input.Email)

		c.JSON(http.StatusOK, gin.H{"message": "User verified and created successfully"})
	}
}
