package model

import "foodGateway/grpc_api/api"

func ToGRPCRequest(rq HTTPPredictRequest) *api.Frames {
	return &api.Frames{Frames: []string{rq.Data}}
}

func ToHTTPResponse(response *api.Prediction) HTTPPredictResponse {
	var frames []Frame

	for _, frame := range response.Result {
		var boxes []Box
		for _, box := range frame.Boxes {
			newBox := Box{Box: box.Box}
			boxes = append(boxes, newBox)
		}
		newFrames := Frame{Boxes: boxes}
		frames = append(frames, newFrames)
	}

	httpResponse := HTTPPredictResponse{Response: frames}
	return httpResponse
}
