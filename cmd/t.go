package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	t, err := time.Parse(time.DateOnly, "1970-01-01")
	if err != nil {
		fmt.Println("error", err)
		os.Exit(1)
	}
	fmt.Println(t.Format(time.DateOnly))
}
