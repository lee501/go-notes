package main

type State int

const (
	Running State = iota
	Stopped
	Rebooting
	Terminated
)

func (s State) String() string {
	switch s {
	case Running:
		return "Runnning"
	case Rebooting:
		return "Rebooting"
	case Stopped:
		return "Stopped"
	case Terminated:
		return "Terminated"
	default:
		return "unknown"
	}
}