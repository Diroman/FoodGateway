package queue
//
//import (
//	"foodGateway/grpc_api/api"
//	"foodGateway/model"
//)
//
//const KILL = 0
//
//type Queue struct {
//	Task chan *model.Task
//	Workers []*Worker
//}
//
//func NewQueue(taskCount, workerCount int) *Queue {
//	var workers []*Worker
//	var taskQueue = make(chan *model.Task, taskCount)
//
//	for i := 0; i < workerCount; i++ {
//		worker := Worker{
//			Task:   taskQueue,
//			Signal: make(chan int),
//		}
//		go worker.Listen()
//
//		workers = append(workers, &worker)
//	}
//
//	return &Queue{Task: taskQueue, Workers: workers}
//}
//
//func (q *Queue) Predict(task *model.Task) *api.Prediction {
//	output := make(chan *api.Prediction)
//	task.Output = output
//
//	q.Task <- task
//	return <- output
//}
//
//func (q *Queue) Close() {
//	for _, worker := range q.Workers {
//		worker.Signal <- KILL
//	}
//}
