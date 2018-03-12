package main

import (
	"context"
	"fmt"
	"time"
)

// :show start
func longMathOp(ctx context.Context, n int) (int, error) {
	res := n
	for i := 0; i < 100; i++ {
		select {
		case <-ctx.Done():
			return 0, ctx.Err()
		default:
			res += i
			// simulate long operation by sleeping
			time.Sleep(time.Millisecond)
		}
	}
	return res, nil
}

func main() {
	ctx, _ := context.WithTimeout(context.Background(), time.Millisecond*200)
	res, err := longMathOp(ctx, 5)
	fmt.Printf("Called longMathOp() with 200ms timeout. res; %d, err: %v\n", res, err)

	ctx, _ = context.WithTimeout(context.Background(), time.Millisecond*10)
	res, err = longMathOp(ctx, 5)
	fmt.Printf("Called longMathOp() with 10ms timeout. res: %d, err: %v\n", res, err)
}

// :show end
