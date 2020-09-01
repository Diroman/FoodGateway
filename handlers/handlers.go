package handlers

import (
	"encoding/json"
	"foodGateway/domain"
	"foodGateway/model"
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

	request := model.ToHTTPRequest(t)
	response, err := r.domain.Predict(request)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Can not transform predict result"))
	}
}
