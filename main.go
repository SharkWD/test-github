package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello World from Golang!")
	t := time.Now()
	hour := t.Hour()
	minutes := t.Minute()
	fmt.Printf("Настоящее время:%d часов %v минут!", hour, minutes)
}
