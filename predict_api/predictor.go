package predict_api

import (
	"context"
	"foodGateway/grpc_api/api"
	"google.golang.org/grpc"
)

type PredictServer struct {
	address string
}

func NewPredictServer(address string) *PredictServer {
	return &PredictServer{address: address}
}

func (ps *PredictServer) Predict(request *api.Frames) (*api.Prediction, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithBlock())
	conn, err := grpc.Dial(ps.address, opts...)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := api.NewFoodDetectorClient(conn)

	response, err := client.Predict(context.Background(), request)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (ps *PredictServer) SetClasses(classes *api.Classes) (*api.Classes, error) {
	return &api.Classes{}, nil
}
