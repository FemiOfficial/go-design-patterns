package composite

import "fmt"

type Animal struct {
}

func (a *Animal) Eat() {
	fmt.Println("Animal ate")
}

func Swim() string {
	fmt.Println("swam")
	return "swam"
}

type AnimalSwimmer struct {
	Animal
	Swim func() string
}

type Athlete struct {
}

func (a *Athlete) Train() {
	fmt.Println("Athlete trained")
}

type AthleteSwimmer struct {
	Athlete
	Swim func() string
}