package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 15; i++ {
			fmt.Printf("%v, ", <-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}

// func sum(s []int, c chan int) {
// 	sum := 0
// 	for _, v := range s {
// 		sum += v
// 	}
// 	c <- sum
// }

// sum arrays, channels testing etc

// func main() {
// 	s := []int{7, 2, 8, -9, -2, 0}
// 	c := make(chan int)

// 	go sum(s[:len(s)/2], c)
// 	go sum(s[len(s)/2:], c)

// 	x, y := <-c, <-c

// 	fmt.Println(x, y, x+y)

// 	go buff.Buffer()
// 	str := make(chan string)
// 	go mnstring.MyString(str)

// 	res := <-str
// 	fmt.Println(sel)

// }

// 1 1 2 3 5 8 13 21 34 55
