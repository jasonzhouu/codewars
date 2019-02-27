// https://www.codewars.com/kata/55b3425df71c1201a800009c/train/go

package main

import "fmt"
import s "strings"
import "strconv"
import "sort"

func string2string_array(strg string) []string {
	// "01|15|59, 1|47|16, 01|17|20, 1|32|34, 2|17|17"
        // [01|15|59 1|47|16 01|17|20 1|32|34 2|17|17]
        // 将第一行的格式转成第二行的格式
	a := s.Split(strg, ", ")
//	fmt.Println(a)
	return a
}

func string_to_2_level_int_slice(strg string) [][3]int {
	// [01|15|59 1|47|16 01|17|20 1|32|34 2|17|17]
	// [[1 15 59] [1 47 16] [1 17 20] [1 32 34] [2 17 17]]
	// 将第一行的格式转成第二行的格式

	a := string2string_array(strg)
        // b 双层slice用于存储时：分：秒 int 格式的列表
        b := make([][3]int, len(a))
        for i1, v1 := range a {
                c := s.Split(v1, "|")
                for i2, v2 := range c {
                        b[i1][i2], _ = strconv.Atoi(v2)
                }
        }
//	fmt.Println(b)
	return b
}

func string_to_seconds_slice(strg string) []int{

	b := string_to_2_level_int_slice(strg)

	// d 单层 slice 用于存储换算得到的秒钟数
	d := make([]int, len(b))
        for i1, v1 := range b {
                for i2, v2 := range v1 {
                        switch i2 {
                                case 0:
                                        d[i1] += v2 * 60 * 60
                                case 1:
                                        d[i1] += v2 * 60
                                case 2:
                                        d[i1] += v2
                        }
                }
        }
//	fmt.Println("input in seconds:\t", d)
	return d
}


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
	var a string = "01|15|59, 1|47|16, 01|17|20, 1|32|34, 2|17|17"
	seconds_slice := string_to_seconds_slice(a)
	seconds_slice_sort := make([]int, 5)
	copy(seconds_slice_sort, seconds_slice)
	sort.Ints(seconds_slice_sort)
	fmt.Println("input in seconds sorted:\t", seconds_slice_sort)
	min := seconds_slice_sort[0]
	max := seconds_slice_sort[4]
	result_range := (max - min)
	fmt.Println("range in seconds:\t", result_range)
	fmt.Println("range in string:\t", num2string(result_range))
}
