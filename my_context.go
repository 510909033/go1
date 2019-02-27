package main

import (
	"context"
	"time"
)

func my_context() {

	parent_ctx := context.Background()
	ctx, cancle_func := context.WithCancel(parent_ctx)

	t, b := ctx.Deadline()
	Info.Println("deadline: t=%v,b=%b", t, b)

	index := 0
	for index < 20 {

		go do1(ctx, index)
		index++
	}

	time.Sleep(time.Second * 2)
	cancle_func()

	time.Sleep(time.Second * 10)
}

func do1(ctx context.Context, index int) {
	if index%3 == 0 {
		go do1_sub(ctx, index)
	}
	select {
	case <-ctx.Done():
		Info.Println("done ,%d", index)
		break
	}
}
func do1_sub(ctx context.Context, index int) {
	if index == 18 {
		time.Sleep(time.Second * 5)
	}
	select {
	case <-ctx.Done():

		Info.Println("done ,sub - %d , err=&s", index, ctx.Err())
		break
	}
}
