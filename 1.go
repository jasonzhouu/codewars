package main

import "fmt"
import s "strings"
import "strconv"
import "sort"

func string2num(strg string) int{
	string_array := s.Split(strg, "|")
	var num_array [3]int
	for i, v := range string_array {
		num_array[i], _ = strconv.Atoi(v)
	}
	var time_in_seconds int
	for i, v := range num_array {
		switch i {
		case 0:
			time_in_seconds += v * 60 * 60
		case 1:
			time_in_seconds += v * 60
		case 2:
			time_in_seconds += v
		}
	}
	return time_in_seconds
}

func num2string(time_in_seconds int) string {
        var num_array [3]int
        for i := 2; i >= 0; i-- {
                num_array[i] = time_in_seconds % 60
                time_in_seconds = time_in_seconds / 60
        }
        var string_array [3]string
        for i, v := range num_array {
                if v < 10 {
                        string_array[i] = "0" + strconv.Itoa(v)
                        continue
                }
                string_array[i] = strconv.Itoa(v)
        }
        var time_in_string string
        string_array2 := []string{string_array[0], string_array[1], string_array[2],}
        for i := 2; i >= 0; i-- {
                time_in_string = s.Join(string_array2, "|")
        }
        return time_in_string
}

func main() {
	var a string = "01|15|59, 1|47|16, 01|17|20, 1|32|34, 2|17|17"
	var b = s.Split(a, ", ")
	var d [5][3]int
	for i1, v := range b {
		c := s.Split(v, "|")
		for i2, v2 := range c {
			d[i1][i2], _ = strconv.Atoi(v2)
		}
	}
	fmt.Println("input in int array:\t", d)
	var e [5]int
	for i1, v1 := range d {
		for i2, v2 := range v1 {
			switch i2 {
				case 0:
					e[i1] += v2 * 60 * 60
				case 1:
					e[i1] += v2 * 60
				case 2:
					e[i1] += v2
			}
		}
	}
	fmt.Println("input in seconds:\t", e)
	f := []int{e[0], e[1], e[2], e[3], e[4],}
	sort.Ints(f)
	fmt.Println("input in seconds sorted:\t", f)
	min := f[0]
	max := f[4]
	result_range := (max - min)
	fmt.Println("range in seconds:\t", result_range)
	fmt.Println("range in string:\t", num2string(result_range))
}
