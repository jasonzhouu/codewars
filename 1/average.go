package main

import "fmt"

func average(seconds_array []int) int {
	var sum int
	for _, v := range seconds_array {
		sum += v
	}
	array_len := len(seconds_array)
	average := sum / array_len
	return average
}

func main(){
	seconds_array := []int{4559, 4640, 5554, 6436, 8237}
	fmt.Println(seconds_array)
	fmt.Println(average(seconds_array))
}
