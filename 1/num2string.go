package main

import "fmt"

func num2string(sec int) string {
        var num_array [3]int
        for i := 2; i >= 0; i-- {
                num_array[i] = sec % 60
                sec = sec / 60
        }
	return fmt.Sprintf("%02d|%02d|%02d", num_array[0], num_array[1], num_array[2])
}

func main() {
	fmt.Println(num2string(4198))
}
