package book

type Service interface {
	FindAll() ([]Book, error)
	findByID(ID int) (Book, error)
	Create(bookRequest BookRequest) (Book, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Book, error) {
	books, err := s.repository.FindAll()
	return books, err
	// return s.repository.FindAll()
}

func (s *service) FindByID(ID int) (Book, error) {
	book, err := s.repository.FindByID(ID)
	return book, err
}

func (s *service) Create(bookRequest BookRequest) (Book, error) {
	price, _ := bookRequest.Price.(int64) // convert interface to float64 dari sebelumnya price.Int64
	// problemnya di 2:16:29 source https://www.youtube.com/watch?v=GjI0GSvmcSU&t=3261s
	book := Book{
		Title: bookRequest.Title,
		Price: int(price),
	}
	newBook, err := s.repository.Create(book)
	return newBook, err

}