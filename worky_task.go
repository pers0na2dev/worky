package worky

type task func() Result

type Task struct {
	fn task
}

type Result struct {
	Meta   Meta
	Result interface{}
	Err    error
}

type Meta map[string]interface{}
