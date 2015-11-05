package main 

import (
	"fmt"
	"runtime"
	"math/rand"
	"time"
)

func main() {
	
	userOs := runtime.GOOS

	switch userOs {
	case "darwin":
		fmt.Println("Hello apple guy")
	case "windows":
		fmt.Println("What are you even doing here? ")
	case "linux":
		fmt.Println( "Ahoy! A linux pirat!")
	default:
		fmt.Println("\n Lol where do you come from ? ")
	}

	switch tmpNum := random(); tmpNum {
	case 0, 2, 4, 6, 8:
		fmt.Println("Even Number - ", tmpNum)
	case 1, 3, 5, 7, 9:
		fmt.Println("Odd Number - ", tmpNum)
	}
}

func random() int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(10)
}