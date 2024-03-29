package main

type State int

const (
	Running State = iota
	Stopped
	Rebooting
	Terminated
)

func (s State) String() string {
	//switch s {
	//case Running:
	//	return "Runnning"
	//case Rebooting:
	//	return "Rebooting"
	//case Stopped:
	//	return "Stopped"
	//case Terminated:
	//	return "Terminated"
	//default:
	//	return "unknown"
	//}
	return []string{"Running", "Stopped", "Rebooting", "Terminated"}[s]

}

//kb mb gb定义
type ByteSize float64

const (
	_           = iota             // ignore first value by assigning to blank identifier
	KB ByteSize = 1 << (10 * iota) // 1 << (10*1)
	MB                             // 1 << (10*2)
	GB                             // 1 << (10*3)
	TB                             // 1 << (10*4)
	PB                             // 1 << (10*5)
	EB                             // 1 << (10*6)
	ZB                             // 1 << (10*7)
	YB                             // 1 << (10*8)
)

func test() ByteSize {
	return MB
}

/*
func main() {
	fmt.Println(Running)
	r := test()
	if r == KB {
		fmt.Println(r)
	}
	fmt.Println(r)
}

*/
