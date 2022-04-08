package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type polar struct {
	radius float64
	angle  float64
}

type cartesian struct {
	x float64
	y float64
}

func main() {
	question := make(chan polar)
	defer close(question)
	answer := createResolver(question)
	defer close(answer)
	interact(question, answer)
}

func createResolver(question chan polar) chan cartesian {
	answer := make(chan cartesian)
	go func() {
		for {
			ques := <-question
			ang := ques.angle * math.Pi / 180.0
			answer <- cartesian{x: ques.radius * math.Cos(ang), y: ques.radius * math.Sin(ang)}
		}
	}()
	return answer
}

func interact(question chan polar, answer chan cartesian) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("radius and angle:")
		in, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		var rad, ang float64
		if _, err = fmt.Sscanf(in, "%f %f", &rad, &ang); err != nil {
			fmt.Println("invalid input")
			break
		}
		question <- polar{rad, ang}
		ans := <-answer
		fmt.Printf("%.02f %.02f -> %.02f %.02f", rad, ang, ans.x, ans.y)
	}
}
