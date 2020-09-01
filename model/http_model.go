package model

type HTTPPredictRequest struct {
	Data string
}

type Box struct {
	Box map[string]float32
}

type Frame struct {
	Boxes []Box
}

type HTTPPredictResponse struct {
	Response []Frame
}

func ToHTTPRequest(rq string) HTTPPredictRequest {
	return HTTPPredictRequest{Data: rq}
}
