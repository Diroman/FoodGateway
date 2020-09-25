package queue
//
//import (
//	"log"
//
//	"foodGateway/model"
//)
//
//type Worker struct {
//	Task chan *model.Task
//	Signal chan int
//}
//
//func (w *Worker) Listen() {
//	for {
//		select {
//		case <- w.Signal:
//			log.Println("Worker stopped")
//			break
//		case task := <- w.Task:
//			w.DoWork(task)
//		}
//	}
//}
//
//func (w *Worker) DoWork(task *model.Task) {
//	response, _ := task.Func(task.Data)
//	task.Output <- response
//}
