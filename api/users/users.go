package users

import (
	"errors"

	"gobackend/database"
	"gobackend/pkg"
	"gobackend/pkg/entities"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// users without paasword
type Users struct {
	database.DefaultModel
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	Image    string `json:"image" gorm:"type:text"`
}

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
	err := db.First(&user, id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return pkg.EntityNotFound("record Not Found")
	} else if err != nil {
		return pkg.Unexpected(err.Error())
	}

	return c.Status(200).JSON(user)
}
func Getyou(c *fiber.Ctx) error {
	user := c.Locals("datauser").(entities.Users)
	// fmt.Println(user)
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
		Image    string `json:"image" gorm:"type:text"`
	}
	db := database.DB
	newuser := new(user)
	if err := c.BodyParser(newuser); err != nil {
		return pkg.BadRequest("Invalid params")
	}

	var existuser Users
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
