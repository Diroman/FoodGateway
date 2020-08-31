package predict_api

import "foodGateway/model"

type Predictor struct {
	host string
	port string
}

func NewPredictor(host, port string) *Predictor {
	return &Predictor{
		host: host,
		port: port,
	}
}

func (p *Predictor) Predict() model.HTTPPredictResponse {
	box := [][]int{
		{100, 100}, {200, 100},
		{200, 200}, {100, 200},
	}

	response := map[string]interface{}{
		"boxes" : box,
		"labels" : []string{"meat"},
	}

	return model.HTTPPredictResponse{Response: response}
}
