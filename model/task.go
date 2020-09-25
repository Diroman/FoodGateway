package model

import "foodGateway/grpc_api/api"

type Task struct {
	Id int64
	Data *api.Frames
	Output chan *api.Prediction
}
