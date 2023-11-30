package main

import "fmt"

func main() {
	/*
		pow := make([]int, 10)
		for i := range pow {
			pow[i] = 1 << uint(i)
		}
		for _, value := range pow {
			fmt.Printf("%d\n", value)
		}
	*/
	ary := [5]string{"taro", "jiro", "sato", "hoo", "foo"}
	for i, s := range ary {
		fmt.Printf("index: %d, name: %s\n", i, s)
	}
}
