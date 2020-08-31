package main

import (
	"fmt"
	domainSt "foodGateway/domain"
	"foodGateway/handlers"
	"foodGateway/predict_api"
	"log"
	"net/http"
)

func main() {
	predictor := predict_api.NewPredictor("127.0.0.1", "8080")
	domain := domainSt.NewDomain(predictor)
	router := handlers.NewRouter(domain)

	http.HandleFunc("/", router.HomeHandler)

	fmt.Println("Start server on http://127.0.0.1")
	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Printf("%s\n", err)
	}
}