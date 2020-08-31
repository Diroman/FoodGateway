package model

type HTTPPredictRequest struct {
	data string
}

type HTTPPredictResponse struct {
	Response map[string]interface{}
}

func RequestToPredictRequest() HTTPPredictRequest {
	return HTTPPredictRequest{}
}
