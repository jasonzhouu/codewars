package main

import "fmt"

func median(seconds_slice_sorted []int) int {
	array_len := len(seconds_slice_sorted)
	if (array_len % 2) == 1 {
		median_index := (array_len + 1) / 2
		median := seconds_slice_sorted[median_index]
		return median
	} else if (array_len % 2) == 0 {
		median_index := array_len / 2
		median := (seconds_slice_sorted[median_index] + seconds_slice_sorted[median_index + 1]) / 2
		return median
	} else {
		return 0
	}
}

func main() {
	seconds_slice_sorted := []int{1993, 8917, 8583, 1847}
	fmt.Println(median(seconds_slice_sorted))
	seconds_slice_sorted = append(seconds_slice_sorted, 1746)
	fmt.Println(median(seconds_slice_sorted))
}
