package domain

import (
	"encoding/base64"
	"errors"
	"fmt"
	"foodGateway/model"
	"foodGateway/predict_api"
	"io/ioutil"
	"log"
	"path"
	"sync"
)

type Counter struct {
	i int
	mux sync.Mutex
}

func (c *Counter) GetIndex() int {
	c.mux.Lock()
	defer c.mux.Unlock()
	retI := c.i
	c.i += 1
	return retI
}

var counter = Counter{}

type Domain struct {
	predictor *predict_api.Predictor
}

func NewDomain(predictor *predict_api.Predictor) *Domain {
	return &Domain{predictor: predictor}
}

func (d *Domain) Predict(body string) (model.HTTPPredictResponse, error) {
	data, err := base64.StdEncoding.DecodeString(body)
	if err != nil {
		log.Printf("Error to decode data: %s\n", err)
		return model.HTTPPredictResponse{}, err
	}

	i := counter.GetIndex()
	fileName := path.Join("images", fmt.Sprintf("%v_input_image.png", i))
	err = ioutil.WriteFile(fileName, data, 0644)
	if err != nil {
		log.Printf("Error to write file: %s\n", err)
		return model.HTTPPredictResponse{}, errors.New("Error in model request")
	}

	response := d.predictor.Predict()
	return response, nil
}