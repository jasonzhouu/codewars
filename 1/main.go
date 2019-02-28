// https://www.codewars.com/kata/55b3425df71c1201a800009c/train/go

package main

import "fmt"
import s "strings"
import "strconv"
import "sort"

func string2string_array(strg string) []string {
	a := s.Split(strg, ", ")
	return a
}

func string_to_duration(strg string) []int {
	a := string2string_array(strg)
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

func average(seconds_slice []int) int {
	var sum int
	for _, v := range seconds_slice {
		sum += v
	}
	array_len := len(seconds_slice)
	average := sum / array_len
	return average
}

func median(seconds_slice_sorted []int) int {
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

func Stati(strg string) string {
	if strg == "" {
		return ""
	}
	seconds_slice := string_to_duration(strg)
	slice_len := len(seconds_slice)
	var result string

	seconds_slice_sort := make([]int, slice_len)
	copy(seconds_slice_sort, seconds_slice)
	sort.Ints(seconds_slice_sort)
	min := seconds_slice_sort[0]
	max := seconds_slice_sort[slice_len - 1]
	result_range := (max - min)
	result += "Range: "
	result += num2string(result_range)

	average := average(seconds_slice)
	result += " Average: "
	result += num2string(average)

	median := median(seconds_slice_sort)
	result += " Median: "
	result += num2string(median)

	return result
}

func main(){
	var a string = "01|15|59, 1|47|16, 01|17|20, 1|32|34, 2|17|17"
//	var a string = "02|15|59, 2|47|16, 02|17|20, 2|32|34, 2|17|17, 2|22|00, 2|31|41"
//	var a string = "00|00|00, 00|00|00, 00|00|00, 00|00|00, 00|00|00, 00|00|00, 00|00|00"
//	var a string = ""
	fmt.Println(Stati(a))
}
