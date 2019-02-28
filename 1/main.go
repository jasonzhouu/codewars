// https://www.codewars.com/kata/55b3425df71c1201a800009c/train/go

package main

import (
	"fmt"
	s "strings"
	"strconv"
	"sort"
)

func string_to_duration(strg string) []int {
	a := s.Split(strg, ", ")
        duration := make([]int, len(a))
        for i1, v1 := range a {
                c := s.Split(v1, "|")

		h, _ := strconv.Atoi(c[0])
		min, _ := strconv.Atoi(c[1])
		sec, _ := strconv.Atoi(c[2])

		duration[i1] = h*60*60 + min*60 + sec
        }
	return duration
}

func num2string(sec int) string {
        var num_array [3]int
        for i := 2; i >= 0; i-- {
                num_array[i] = sec % 60
                sec = sec / 60
        }
        return fmt.Sprintf("%02d|%02d|%02d", num_array[0], num_array[1], num_array[2])
}

func get_average(seconds_slice []int) int {
	sum := 0
	for _, v := range seconds_slice {
		sum += v
	}
	array_len := len(seconds_slice)
	average := sum / array_len
	return average
}

func get_median(seconds_slice_sorted []int) int {
	array_len := len(seconds_slice_sorted)
	if (array_len % 2) == 1 {
		median_index := (array_len + 1) / 2
		median := seconds_slice_sorted[median_index-1]
		return median
	} else if (array_len % 2) == 0 {
		median_index := array_len / 2
		median := (seconds_slice_sorted[median_index] + seconds_slice_sorted[median_index - 1]) / 2
		return median
	} else {
		return 0
	}
}

func get_range(seconds_slice_sort []int) int {
	min := seconds_slice_sort[0]
	max := seconds_slice_sort[len(seconds_slice_sort) - 1]
	return (max - min)
}

func Stati(strg string) string {
	if strg == "" {
		return ""
	}

	seconds_slice := string_to_duration(strg)
	sort.Ints(seconds_slice)

	return fmt.Sprintf(
		"Range: %s Average: %s Median: %s",
		num2string(get_range(seconds_slice)),
		num2string(get_average(seconds_slice)),
		num2string(get_median(seconds_slice)),
	)
}

func main(){
	var a string = "01|15|59, 1|47|16, 01|17|20, 1|32|34, 2|17|17"
//	var a string = "02|15|59, 2|47|16, 02|17|20, 2|32|34, 2|17|17, 2|22|00, 2|31|41"
//	var a string = "00|00|00, 00|00|00, 00|00|00, 00|00|00, 00|00|00, 00|00|00, 00|00|00"
//	var a string = ""
	fmt.Println(Stati(a))
}
