package server

import (
	"database/sql"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"library/internal"
	"library/internal/db"
	"strconv"
)

const (
	Err         = "Error"
	Status      = "Status:"
	Success     = "Success"
	Unsuccess   = "Unsuccess"
	Books       = "books"
	Book        = "book"
	ErrSaveBook = "Dont save book in db"
	ErrBindJSON = "Dont bind tag"
	ErrConvert  = "Cant convert ascii to int"
)

func SetupRouter(Database *sql.DB) *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:63342"}, // Укажи URL фронтенда
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type"},
		AllowCredentials: true,
	}))

	router.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Load status": "Successfull",
		})
	})

	router.GET("/books", func(c *gin.Context) {
		books, err := db.AllBooks(Database)
		if err != nil {
			c.JSON(500, gin.H{
				Status: Unsuccess,
				Err:    err,
			})
			return
		}

		c.JSON(200, gin.H{
			Status: Success,
			Books:  books,
		})
	})

	router.POST("/book", func(c *gin.Context) {
		var book internal.Book

		if err := c.ShouldBindJSON(&book); err != nil {
			c.JSON(500, gin.H{
				Status: Unsuccess,
				Err:    ErrBindJSON,
			})
			return
		}

		if err := db.InsertBook(Database, book); err != nil {
			c.JSON(500, gin.H{
				Status: Unsuccess,
				Err:    ErrSaveBook,
			})
			return
		}

		c.JSON(200, gin.H{
			Status: Success,
			Book:   book,
		})

	})

	router.GET("/book/:id", func(c *gin.Context) {
		query := c.Param("id")

		id, err := strconv.Atoi(query)
		if err != nil {
			c.JSON(400, gin.H{
				Status: Unsuccess,
				Err:    ErrConvert,
			})
			return
		}

		InfoForBook, err := db.InfoForID(Database, id)
		if err != nil {
			c.JSON(500, gin.H{
				Status: Unsuccess,
				Err:    "Book with this id is absent",
			})
			return
		}

		c.JSON(200, gin.H{
			Status: Success,
			"Info": InfoForBook,
		})

	})

	router.DELETE("/book/:id", func(c *gin.Context) {
		query := c.Param("id")

		id, err := strconv.Atoi(query)
		if err != nil {
			c.JSON(400, gin.H{
				Status: Unsuccess,
				Err:    ErrConvert,
			})
			return
		}

		if err := db.DeleteFromID(Database, id); err != nil {
			c.JSON(500, gin.H{
				Status: Unsuccess,
				Err:    "Book with this id is absent",
			})
			return
		}

		c.JSON(200, gin.H{
			Status:    Success,
			"message": "Boook with this id is deleted",
		})

	})

	return router
}
