package repository

type Repository[M any] interface {
	GetAll(filter interface{}) (interface{}, error)
	GetByID(id string) (interface{}, error)
	Update(model *M) (interface{}, error)
	Store(model *M) (interface{}, error)
	Delete(filter interface{}) (bool, error)
}