package interfaces

type JobInterface interface {
	Init(data interface{})
	Run()
	Close()
	Error(err error)
}
