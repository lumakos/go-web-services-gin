package main

import (
  "log"
  "os"
  "github.com/joho/godotenv"
  "net/http"
  "github.com/gin-gonic/gin"
)

// car represents data
type car struct {
	ID string `json:"id"`
	Make string `json:"make"`
	Model string `json:"model"`
	Price int `json:"price"`
	Currency string `json:"currency"`
}

// add data to json
var cars = []car{
	{ID: "1", Make: "BMW", Model: "M3", Price: 65010, Currency: "EUR"},
	{ID: "2", Make: "BMW", Model: "Z3", Price: 44010, Currency: "EUR"},
	{ID: "3", Make: "Audi", Model: "R8", Price: 165000, Currency: "EUR"},
}


// use godot package to load/read the .env file and
// return the value of the key
func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
	log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}


func main() {
	router := gin.Default()
	router.GET("cars", getCars)
	router.GET("cars/:id", getCarById)
	router.POST("cars", postCars)

	// godotenv package
	dotenv := goDotEnvVariable("HOST_URL")

	// Run server
	router.Run(dotenv)
}

// Get all cars
func getCars(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, cars)

}

// Post a new car
func postCars(c *gin.Context) {
	var newCar car

	if err := c.BindJSON(&newCar); err != nil {
		return
	}

	// Add the new car
	cars = append(cars, newCar)
	c.IndentedJSON(http.StatusCreated, newCar)
}

func getCarById(c *gin.Context) {
	id := c.Param("id")

	for _, a:= range cars {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "car not found"})
}