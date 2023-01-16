package book

type Service interface {
	CreateBook(req BookRequest) (res DefaultResponse, err error)
	GetAllBook(limit, page string) (res GetAllBookResponse, err error)
}
