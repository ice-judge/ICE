package main

import ("fmt"
    "time")

func main() {
    fmt.Println("hello world")
    for ;; {
	var num1 int
	fmt.Scanf("%d", &num1)
	time.Sleep(500*time.Millisecond)
	fmt.Printf("%d\n", num1)

    }
}
