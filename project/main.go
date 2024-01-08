package project

import (
	"fmt"
	"time"
)

// Estructura de la tareas de procesar
type Job struct {
	Name   string        //Nombre de la tarea
	Delay  time.Duration //Tiempo de espera
	Number int           // Numero a procesar
}

type Worker struct {
	Id         int           // id del Worker
	JobQueue   chan Job      // Canal de tareas del worker
	WorkerPool chan chan Job //Canal de canales de tareas, este canal se comparte entre todos los workers
	QuitChan   chan bool     //Canal para parar al worker
}

type Dispatcher struct {
	WorkerPool chan chan Job //Canal de canales de tareas, este se les pasa a cada worker nuevo
	MaxWorkers int           //cantidad maxima de workers
	JobQueue   chan Job      //Canal de tareas, se puede ver como un canal global de tareas que despues se reparten entre workers
}

func NewWorker(id int, workerPool chan chan Job) *Worker {
	return &Worker{
		Id:         id,              //Se asigna un id
		WorkerPool: workerPool,      //Se le indica el canal donde tiene quie agregar su canal de tareas
		JobQueue:   make(chan Job),  //Canal de tareas del worker
		QuitChan:   make(chan bool), //Canal para parar al worker
	}
}

func (w Worker) Start() {

	//Se inicia de manera concurrente un ciclo sin fin
	go func() {
		for {

			//Al worker pool se manda el canal de worker, este se manda cada vez iteracion, es decir cuando el worker termino de hacer un jobs
			w.WorkerPool <- w.JobQueue

			//Se multiplexean los canales del worker
			select {
			case job := <-w.JobQueue:
				//Si se recibe un job en el canal de tareas del worker se ejecuta
				fmt.Printf("Worker with id %d Started\n", w.Id)
				fib := Fibonacci(job.Number)
				time.Sleep(job.Delay)
				fmt.Printf("Worker with id %d Finishes with result %d\n", w.Id, fib)

			case <-w.QuitChan:
				//Si se recibe un job en el canal de salida se para el worker (lo sca del ciclo)
				fmt.Printf("Worker with id %d Stopped\n", w.Id)
				return
			}

		}
	}()
}

// La funcion stop manda un true al canl de salida del worker
func (w Worker) Stop() {
	go func() {
		w.QuitChan <- true
	}()
}

//El dispatcher cuenta con el el canal global de jobs y un canal de todos los canales de los workers

func NewDispatcher(jobQueue chan Job, maxWorkers int) *Dispatcher {

	worker := make(chan chan Job, maxWorkers)
	return &Dispatcher{
		JobQueue:   jobQueue,
		MaxWorkers: maxWorkers,
		WorkerPool: worker,
	}
}

func (d *Dispatcher) Dispatch() {

	//Inicia de manera indefinidad a mandar jobs a los canales que se van recibiendo en el canal de caneles de jobs
	for {
		select {
		case job := <-d.JobQueue:
			go func() {
				workerJobQueue := <-d.WorkerPool
				workerJobQueue <- job
			}()
		}
	}
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	return Fibonacci(n-1) + Fibonacci(n-2)
}
