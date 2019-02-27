package main

import "fmt"
import s "strings"
import "strconv"

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

func main() {
	var a string = "01|15|59, 1|47|16, 01|17|20, 1|32|34, 2|17|17"
	fmt.Println(string_to_seconds_slice(a))
}
