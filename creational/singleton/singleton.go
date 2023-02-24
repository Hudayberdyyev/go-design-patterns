package singleton

type singleton struct {
	counter int
}

var instance *singleton

func (s *singleton) AddOne() int {
	s.counter++
	return s.counter
}

func GetInstance() *singleton {
	if instance == nil {
		instance = new(singleton)
	}
	return instance
}

func (s *singleton) GetCount() int {
	return s.counter
}
