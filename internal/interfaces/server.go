package interfaces

import (
	"test/internal/domain/repository"
	"test/internal/interfaces/handler"
	"test/internal/interfaces/infra"
	"test/internal/interfaces/usecase/book"

	"github.com/labstack/echo/v4"
)

var (
	//setup DB
	db = infra.SetupDB()
	//setup Repository
	bookRepository = repository.InitBookRepository(db)
	//setup Usecase
	bookService = book.NewService(db, bookRepository)
	//setup controller
	bookController = handler.InitBookHandler(bookService)
)

func SetupRouter(app *echo.Echo) {
	app.POST("/book", bookController.CreateBook)
	app.GET("/book", bookController.GetAllBook)
}
