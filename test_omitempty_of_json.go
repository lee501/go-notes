package main

import (
	"encoding/json"
	"fmt"
)

type Message struct {
	FinalTestingRes FinalTestingRes `json:"final_testing_res,omitempty"`
}

type FinalTestingRes struct {
	Results       []TestingRes `json:",omitempty"`
	FinalAction   string       `json:",omitempty"`
	FinalRiskName string       `json:",omitempty"`
}

type TestingRes struct {
	Action   string
	RiskName string
	Reason   string
	IsOK     bool `json:"-"`
}

func testEmitemptyOfJson() {
	mess := Message{
		FinalTestingRes: FinalTestingRes{
			[]TestingRes{},
			"block",
			"",
		},
	}
	b, _ := json.Marshal(mess)
	fmt.Println(string(b))
}
