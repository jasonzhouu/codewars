package main

import "fmt"
import "strconv"
import s "strings"

func num2string(time_in_seconds int)string {
	num_array := make([]int, 3)
	for i := 2; i >= 0; i-- {
		num_array[i] = time_in_seconds % 60
		time_in_seconds = time_in_seconds / 60
	}
	string_array := make([]string, 3)
	for i, v := range num_array {
		if v < 10 {
			string_array[i] = "0" + strconv.Itoa(v)
			continue
		}
		string_array[i] = strconv.Itoa(v)
	}
	var time_in_string string
	for i := 2; i >= 0; i-- {
		time_in_string = s.Join(string_array, "|")
	}
	return time_in_string
}

func main() {
	fmt.Println(num2string(3919))
}
