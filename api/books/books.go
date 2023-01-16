package books

import (
	"errors"

	"gobackend/database"
	"gobackend/pkg"
	"gobackend/pkg/entities"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetBooks(c *fiber.Ctx) error {
	db := database.DB
	var books []entities.Books
	db.Find(&books)
	return c.JSON(books)
}

func GetBook(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var book entities.Books
	err := db.First(&book, id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return pkg.EntityNotFound("No book found")
	} else if err != nil {
		return pkg.Unexpected(err.Error())
	}

	return c.JSON(book)
}

func NewBook(c *fiber.Ctx) error {
	user := c.Locals("datauser").(entities.Users)
	db := database.DB
	book := new(entities.Books)
	if err := c.BodyParser(book); err != nil {
		return pkg.BadRequest("Invalid params")
	}
	book.Created_by = int64(user.ID)
	db.Create(&book)
	return c.JSON(book)
}

func UpdateBook(c *fiber.Ctx) error {

	id := c.Params("id")
	db := database.DB
	var book entities.Books
	err := db.First(&book, id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return pkg.EntityNotFound("No book found")
	} else if err != nil {
		return pkg.Unexpected(err.Error())
	}

	updatedBook := new(entities.Books)

	if err := c.BodyParser(updatedBook); err != nil {
		return pkg.BadRequest("Invalid params")
	}

	updatedBook = &entities.Books{Title: updatedBook.Title, Author: updatedBook.Author, Rating: updatedBook.Rating}

	if err = db.Model(&book).Updates(updatedBook).Error; err != nil {
		return pkg.Unexpected(err.Error())
	}

	return c.SendStatus(204)
}

func DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB

	var book entities.Books
	err := db.First(&book, id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return pkg.EntityNotFound("No book found")
	} else if err != nil {
		return pkg.Unexpected(err.Error())
	}

	db.Delete(&book)
	return c.SendStatus(204)
}
