package users

import (
	"errors"
	"fmt"
	"time"

	"gobackend/database"
	"gobackend/pkg"
	"gobackend/pkg/entities"
	"gobackend/storage"

	"github.com/gofiber/fiber/v2"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// users without paasword
// type Users struct {
// 	database.DefaultModel
// 	Username string `json:"username"`
// 	Email    string `json:"email"`
// 	Role     string `json:"role"`
// 	Image    string `json:"image" gorm:"type:text"`
// }

func GetUsers(c *fiber.Ctx) error {
	db := database.DB
	var users []entities.Users
	db.Find(&users)
	return c.Status(200).JSON(users)
}

func Getuser(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var user entities.Users
	err := db.Where("id = ?", id).First(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return pkg.EntityNotFound("record Not Found")
	} else if err != nil {
		return pkg.Unexpected(err.Error())
	}

	return c.Status(200).JSON(user)
}
func Getyou(c *fiber.Ctx) error {
	user := c.Locals("datauser").(entities.Users)
	return c.Status(200).JSON(user)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func NewUser(c *fiber.Ctx) error {
	type user struct {
		database.DefaultModel
		Username string `json:"username" gorm:"uniqueIndex"`
		Email    string `json:"email" gorm:"uniqueIndex"`
		Password string `json:"password" gorm:"type:text"`
		// Issuperadmin bool   `json:"issuperadmin"`
		Image string `json:"image" gorm:"type:text"`
	}
	db := database.DB
	newuser := new(user)
	if err := c.BodyParser(newuser); err != nil {
		return pkg.BadRequest("Invalid params")
	}

	var existuser entities.Users
	err := db.Where("username = ?", newuser.Username).First(&existuser).Error
	if !(errors.Is(err, gorm.ErrRecordNotFound)) {
		return c.JSON(fiber.Map{
			"message": "user telah ada",
		})
	}

	hash, _ := HashPassword(newuser.Password)
	newuser.Password = hash
	newuser.ID = uuid.New()
	dberr := db.Create(&newuser).Error

	if dberr != nil {
		return pkg.BadRequest("Failet to save user")
	}
	return c.JSON(newuser)
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB

	var user entities.Users
	err := db.First(&user, id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return pkg.EntityNotFound("No user found")
	} else if err != nil {
		return pkg.Unexpected(err.Error())
	}
	db.Delete(&user)
	return c.SendStatus(204)
}

func UpdateUser(c *fiber.Ctx) error {
	//update Repository
	type user struct {
		database.DefaultModel
		Username string `json:"username" gorm:"uniqueIndex"`
		Email    string `json:"email" gorm:"uniqueIndex"`
		Password string `json:"-" gorm:"type:text"`
		Image    string `json:"image" gorm:"type:text"`
	}
	uservalidate := new(user)
	var modeluser entities.Users

	db := database.DB
	id := c.Params("id")
	err := db.First(&modeluser, id).Error
	// return c.JSON(user)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return pkg.EntityNotFound("No book found")
	} else if err != nil {
		// return pkg.Unexpected(err.Error())
		return pkg.Unexpected("i dont know")
	}

	if err := c.BodyParser(uservalidate); err != nil {
		return pkg.BadRequest("invalid params")
	}

	err2 := db.Model(&modeluser).Updates(&entities.Users{Email: uservalidate.Email, Username: uservalidate.Username}).Error
	if err2 == nil {
		return pkg.BadRequest("Failed To Save User")
	}

	return c.JSON(uservalidate)

}

func UploadAvatar(c *fiber.Ctx) error {
	if form, err := c.MultipartForm(); err == nil {
		files := form.File["image"]
		if len(files) != 0 {
			for _, file := range files {
				maxFileSize := int64(1 * 1024 * 1024)
				if file.Size > maxFileSize {
					return c.Status(413).JSON(pkg.BadRequest("file not more than 1 mb"))
				}
				path := fmt.Sprintf("avatar/%02d%s", time.Now().Nanosecond(), file.Filename)
				fmt.Println("sebelum save file")
				if err := c.SaveFileToStorage(file, path, storage.Storage()); err != nil {
					return err
				}
				fmt.Println("sesudah save file")
				user := c.Locals("datauser").(entities.Users)
				db := database.DB
				db.Model(user).Update("image", path)
				c.Status(200).JSON(fiber.Map{
					"file": user.Image,
				})
			}
		} else {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		return err
	} else {
		return c.SendStatus(fiber.StatusBadRequest)
	}
}
