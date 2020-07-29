package goroutinePool

type Work struct {
	pool *Pool
	task chan f
}
