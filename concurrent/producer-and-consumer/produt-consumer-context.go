package main

import "context"


func Process(ctx context.Context) {
	//控制product线程
	n := 3
	for i := 0; i<n; i++ {
		go producter()
	}
}
