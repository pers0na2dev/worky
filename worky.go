package worky

// Worky is library for creating and managing goroutine working pool with a fixed number of workers.
// Easy to use and simple to understand.
type Worky struct {
	// Number of workers
	workers int

	// Channel for tasks
	tasks chan Task

	// Channel for results
	result chan Result
}

// New creates a new Worky instance with a fixed number of workers.
func New(workers int) *Worky {
	return &Worky{
		workers: workers,
		tasks:   make(chan Task, workers*2),
		result:  make(chan Result, workers*2),
	}
}

// Start starts the Worky instance and its workers.
func (w *Worky) Start() {
	for i := 0; i < w.workers; i++ {
		go func() {
			for {
				select {
				case task := <-w.tasks:
					w.result <- task.fn()
				}
			}
		}()
	}
}

// Add adds a task to the Worky instance without waiting for an acknowledgment.
func (w *Worky) Add(fn task) {
	w.tasks <- Task{fn: fn}
}

// Results returns a channel for receiving results from the Worky instance.
func (w *Worky) Results() <-chan Result {
	return w.result
}
