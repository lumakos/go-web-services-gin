# RESTful web service API with Go and Gin.

## Run
>go run main.go

## GET endpoint /cars
>curl localhost:8088/cars --header "Content-Type: application/json" --request "GET"

## GET endpoint /cars/:id
>curl localhost:8088/cars/1 --header "Content-Type: application/json" --request "GET"

## POST /cars
 > curl localhost:8088/cars --include --header "Content-Type: application/json" --request "POST" --data '{"id": "4", "make": "Audi", "model": "A1", "price": 165000, "currency": "EUR"}'