package main

import (
	"fmt"
	"log"
	"net/http"

	domainSt "foodGateway/domain"
	"foodGateway/handlers"
	"foodGateway/predict_api"
)

func main() {
	predictor := predict_api.NewPredictServer("192.168.0.100:50051")
	//station := queue.NewQueue(1000, 5)
	domain := domainSt.NewDomain(predictor)
	router := handlers.NewRouter(domain)

	http.HandleFunc("/", router.HomeHandler)

	fmt.Println("Start server on http://127.0.0.1")
	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Printf("%s\n", err)
	}

}