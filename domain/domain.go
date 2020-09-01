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
	i   int
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
	predictor *predict_api.PredictServer
}

func NewDomain(predictor *predict_api.PredictServer) *Domain {
	return &Domain{predictor: predictor}
}

func (d *Domain) Predict(rq model.HTTPPredictRequest) (model.HTTPPredictResponse, error) {
	if err := d.saveImage(rq); err != nil {
		log.Println(fmt.Sprintf("Error to save image: %s", err))
	}

	request := model.ToGRPCRequest(rq)
	response, err := d.predictor.Predict(request)
	httpResponse := model.ToHTTPResponse(response)

	if err != nil {
		log.Println(fmt.Sprintf("Error to get predict: %s", err))
		return model.HTTPPredictResponse{}, err
	}

	return httpResponse, nil
}

func (d Domain) saveImage(request model.HTTPPredictRequest) error {
	data, err := base64.StdEncoding.DecodeString(request.Data)
	if err != nil {
		log.Printf("Error to decode data: %s\n", err)
		return err
	}

	i := counter.GetIndex()
	fileName := path.Join("images", fmt.Sprintf("%v_input_image.png", i))
	err = ioutil.WriteFile(fileName, data, 0644)
	if err != nil {
		log.Printf("Error to write file: %s\n", err)
		return errors.New("Error in model request\n")
	}

	return nil
}
