package main

import (
	"fmt"
)

type BizType int

const (
	SearchFlight BizType = iota
	CalendarFlight
	BookingFlight
)

//func (b BizType) ValueString() string {
//	return strconv.Itoa(int(b))
//}

func (b BizType) String() string {
	return [...]string{"search flight", "calendar flight", "booking flight"}[b]
}

func main() {
	var search BizType = SearchFlight
	fmt.Println(search)
}