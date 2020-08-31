package handlers

import (
	"encoding/json"
	"foodGateway/domain"
	"log"
	"net/http"
)

type Router struct {
	domain *domain.Domain
}

func NewRouter(domain *domain.Domain) *Router {
	return &Router{domain: domain}
}


func (r *Router) HomeHandler(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)

	var t string
	if err := decoder.Decode(&t); err != nil {
		log.Printf("Error to decode body: %s\n", err)
		return
	}

	response, err := r.domain.Predict(t)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response.Response); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Can not transform predict result"))
	}
}
