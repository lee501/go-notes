package goroutinePool

type signal struct {}

type f func() error

type Pool struct {
	capacity int32
	running  int32
	freeSignal chan signal
}