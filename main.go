package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Wrong number of arguments")
		return
	}
	argument := os.Args[1]
	file, err := ioutil.ReadFile(argument)
	if err != nil {
		fmt.Println("Invalid name of the file.\nPlease, try [go run . data.txt]")
		return
	}
	hello := strings.Fields(string(file))
	var second []float64
	for i := 0; i < len(hello); i++ {
		file, _ := strconv.Atoi(hello[i])
		second = append(second, float64(file))
	}
	Average(second)
	Median(second)
	Variance(second)
	StandardDeviation(second)
}

func Average(numbers []float64) {
	sum := 0.0
	length := len(numbers)
	for i := 0; i < len(numbers); i++ {
		sum = float64(sum + float64(numbers[i]))
	}
	average := float64(sum / float64(length))
	fmt.Printf("Average: %.0f\n", average)
}

func Median(numbers []float64) {
	for i := 0; i < len(numbers); i++ {
		for k := 0; k < len(numbers)-i-1; k++ {
			if numbers[k] > numbers[k+1] {
				tmp := numbers[k]
				numbers[k] = numbers[k+1]
				numbers[k+1] = tmp
			}
		}
	}
	length := len(numbers)
	even := (float64(numbers[length/2]) + float64(numbers[length/2-1])) / float64(2)
	odd := float64(numbers[length/2])
	if length%2 == 0 {
		fmt.Printf("Median: %d\n", int(math.Round(even)))
		return
	}
	fmt.Printf("Median: %d\n", int(math.Round(odd)))
	return
}

func Variance(numbers []float64) {
	sum := 0.0
	var length float64
	for i := 0; i < len(numbers); i++ {
		length++
	}
	for i := 0; i < len(numbers); i++ {
		sum = sum + float64(numbers[i])
	}
	mean := sum / float64(length)
	var d float64
	var hello []float64
	for i := 0; i < len(numbers); i++ {
		d = mean - numbers[i]
		s := d * d
		hello = append(hello, s)
	}
	sum1 := 0.0
	for i := 0; i < len(hello); i++ {
		sum1 = sum1 + hello[i]
	}
	var variance float64
	variance = sum1 / length
	fmt.Printf("Variance: %d\n", int(math.Round(variance)))
}

func StandardDeviation(numbers []float64) {
	sum := 0.0
	var length float64
	for i := 0; i < len(numbers); i++ {
		length++
	}
	for i := 0; i < len(numbers); i++ {
		sum = sum + numbers[i]
	}
	mean := sum / length
	var d float64
	var hello []float64
	for i := 0; i < len(numbers); i++ {
		d = mean - numbers[i]
		s := d * d
		hello = append(hello, s)
	}
	sum1 := 0.0
	for i := 0; i < len(hello); i++ {
		sum1 = sum1 + hello[i]
	}

	variance := float64(sum1 / length)
	fmt.Printf("Standard Deviation: %d\n", int(math.Round(math.Sqrt(variance))))
}
