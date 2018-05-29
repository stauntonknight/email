package main

import (
	"fmt"
	"math/rand"
	"time"
	"math"
	"math/cmplx"
)

var x = cmplx.Sqrt(-1)

func add(x int, y int) int {
	return x + y
}

const Big = 1 << 100

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println("Hello world")
	fmt.Println("Hello world ", rand.Intn(10))
	fmt.Println("Hello world ", math.Sqrt(10))
	fmt.Println("Hello world ", math.Pi)
	a := add(2,3)
	fmt.Println(x)
	fmt.Println(a)
	fmt.Println(float64(Big))

	var sum int
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
	sum = 1
	kount := 1
	for ; sum < 1000; {
		sum += sum
		kount += 1
	}
	fmt.Println(kount, sum)
	target := 273573.0
	ans := float64(rand.Intn(100))
	kount = 1
	for ; math.Abs(ans * ans  - target) > .0005; {
		ans -= (ans * ans - target) / (2 * ans)
		kount += 1
	}
	fmt.Println(kount, ans, math.Abs(ans * ans - target))
	t := time.Now()
	switch {
		case t.Hour() < 12:
			fmt.Println("Good morning!")
		case t.Hour() < 17:
			fmt.Println("Good afternoon.")
		default:
		fmt.Println("Good evening.")
	}

	for  i := 0 ; i < 10 ; i++ {
		defer fmt.Println(i)
	}
	v := Vertex {1,2}
	p := &v
	fmt.Println(p.X)
	defer fmt.Println("Done")
	fmt.Println(Vertex{1,2})

	array := [4] int {1,2,3,4}
	b := array[1:4]
	fmt.Println(b)
}

