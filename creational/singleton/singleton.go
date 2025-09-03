package creational

type Singleton struct {
	count int
}

var instance *Singleton

func GetInstance() *Singleton {
	if instance == nil {
		instance = &Singleton{}
		
	}

	return instance
}

func (s *Singleton) AddOne() int {
	s.count++
	return s.count
}