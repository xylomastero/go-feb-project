package main

import (
	"fmt"
)

type Car struct {
	Make  string
	Model string
	Year  int
}

func (c Car) DisplayInfo() {
	fmt.Printf("This is a %d %s %s.\n", c.Year, c.Make, c.Model)
}
func (c *Car) UpdateYear(newYear int) {
	c.Year = newYear
}
func printCarOwners(carOwners map[string]string) {
	for car, owner := range carOwners {
		fmt.Printf("%s is owned by %s.\n", car, owner)
	}
}
func createCar(make string, model string, year int) *Car {
	return &Car{Make: make, Model: model, Year: year}
}

func MapsMethods() {
	tesla := Car{Make: "Tesla", Model: "Model 3", Year: 2021}
	tesla.DisplayInfo()

	tesla.UpdateYear(2022)
	fmt.Printf("After update, the Tesla is a %d model.\n", tesla.Year)

	ford := createCar("Ford", "Mustang", 2020)
	ford.DisplayInfo()

	carOwners := map[string]string{
		"Tesla Model 3": "Alice",
		"Ford Mustang":  "Bob",
		"Toyota Camry":  "Carol",
	}

	printCarOwners(carOwners)
}
