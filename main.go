package main

import "fmt"

func fuelGuage(fuel int) {
  fmt.Println(fuel)
}

func calculateFuel(planet string) int {
  var fuel int
  switch planet {
  case "Venus":
    fuel = 300000
  case "Mercury":
    fuel = 500000
  case "Mars":
    fuel = 700000
  }
  return fuel
}

func greetPlanet(planet string) {
  fmt.Printf("Welcome! You are traveling to %v \n", planet)
}

func cantFly() {
  fmt.Println("We do not have the available fuel to fly there.")
}

func flyToPlanet(planet string, fuel int) int {
  var fuelRemaining, fuelCost int
  fuelRemaining = fuel
  fuelCost = calculateFuel(planet)
  if fuelRemaining >= fuelCost {
    greetPlanet(planet)
    fuelRemaining -= fuelCost
  } else {
    cantFly()
  }
  return fuelRemaining
}

func main() {
  var planetChoice string
  fuel := 10000
  question := "Which planet would you like to visit?"
  fmt.Println(question)
  fmt.Scan(&planetChoice)
  
  fuel = flyToPlanet(planetChoice, fuel)
  fuelGuage(fuel)
}