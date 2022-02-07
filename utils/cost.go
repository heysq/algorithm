package utils

import (
	"fmt"
	"time"
)

func Cost(t time.Time){
	fmt.Printf("func cost: %d ms\n", time.Since(t).Milliseconds())
}