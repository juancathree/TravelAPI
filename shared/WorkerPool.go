package shared

func init() {
	Pool = WorkerPool{
		MaxWorkers: 10,
		QueueTasks: make(chan func(), 10),
	}
	Pool.Run()
}

var Pool WorkerPool

type WorkerPool struct {
	MaxWorkers int
	QueueTasks chan func()
}

func (w *WorkerPool) Run() {
	for i := 0; i < w.MaxWorkers; i++ {
		go func() {
			for {
				task := <-w.QueueTasks
				task()
			}
		}()
	}
}

func (w *WorkerPool) AddTask(task func()) {
	w.QueueTasks <- task
}
