package language

import (
	"context"
	"fmt"
	"time"
)

func Break() {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
	_ = cancelFunc
LOOP:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("done")
			break LOOP
		case <-time.After(1 * time.Second):
			fmt.Println(1)
		}
	}
}
