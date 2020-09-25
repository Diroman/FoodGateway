package domain

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"path"
	"sync"

	"foodGateway/grpc_api/api"
	"foodGateway/model"
	"foodGateway/predict_api"
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
	queue chan *model.Task
}

func NewDomain(predictor *predict_api.PredictServer) *Domain {
	queue := make(chan *model.Task)
	domain := &Domain{
		predictor: predictor,
		queue: queue,
	}

	for i := 0; i < 5; i++ {
		go domain.ListenQueue()
	}

	return domain
}

func (d *Domain) Predict(rq model.HTTPPredictRequest) (model.HTTPPredictResponse, error) {
	if err := d.saveImage(rq); err != nil {
		log.Println(fmt.Sprintf("Error to save image: %s", err))
	}

	request := model.ToGRPCRequest(rq)
	task := model.Task{
		Id:     0,
		Data:   request,
		Output: make(chan *api.Prediction),
	}
	d.queue <- &task
	response := <- task.Output
	httpResponse := model.ToHTTPResponse(response)

	//if err != nil {
	//	log.Println(fmt.Sprintf("Error to get predict: %s", err))
	//	return model.HTTPPredictResponse{}, err
	//}

	return httpResponse, nil
}

func (d *Domain) saveImage(request model.HTTPPredictRequest) error {
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

func (d *Domain) ListenQueue()  {
	for {
		task := <- d.queue
		response, _ := d.predictor.Predict(task.Data)
		task.Output <- response
	}
}
