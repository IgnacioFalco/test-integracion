package products

type Service interface {
	GetAll() ([]Product, error)
	Store(name, productType string, count int, price float64) (Product, error)
	Update(id int, name, productType string, count int, price float64) (Product, error)
	UpdateName(id int, name string) (Product, error)
	Sum(prices ...float64) float64
	Delete(id int) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll() ([]Product, error) {
	ps, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return ps, nil
}

func (s *service) Store(nombre, tipo string, cantidad int, precio float64) (Product, error) {
	lastID, _ := s.repository.LastID()

	lastID++

	producto, err := s.repository.Store(lastID, nombre, tipo, cantidad, precio)
	if err != nil {
		return Product{}, err
	}

	return producto, nil
}

func (s *service) Update(id int, name, productType string, count int, price float64) (Product, error) {

	return s.repository.Update(id, name, productType, count, price)
}

func (s *service) UpdateName(id int, name string) (Product, error) {
	return s.repository.UpdateName(id, name)
}

func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}

func (s *service) Sum(prices ...float64) float64 {
	var price float64
	for _, p := range prices {
		price += p
	}
	return price
}
