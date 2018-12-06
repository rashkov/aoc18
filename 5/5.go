package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Use(vals ...interface{}) {
	for _, val := range vals {
		_ = val
	}
}

func main() {
	polymer, err := ioutil.ReadFile("./input.txt")
	check(err)
	polymer_str := strings.TrimSpace(string(polymer))

	// var a string = "dabAcCaCBAcCcaDA"
	var a string = polymer_str
	// fmt.Println(len(a), len(strings.TrimSpace(a)))

	res1 := reduce(a)
	len1 := len(res1)
	res2 := reduce(res1)
	len2 := len(res2)
	for len2 < len1 {
		res1 = res2
		len1 = len2
		res2 = reduce(res1)
		len2 = len(res2)
	}
	fmt.Println("Part 1", len(res2))
}

func should_destroy(a byte, b byte) bool{
	return same_type(a,b) && reverse_polarity(a,b)
}

func same_type(a byte, b byte) bool{
	return strings.ToUpper(string(a)) == strings.ToUpper(string(b))
}

func reverse_polarity(a byte, b byte) bool{
	return same_type(a, b) && a != b
}

func reduce(str string) string {
	if len(str) > 2 {
		left := reduce(str[:len(str)/2])
		right := reduce(str[len(str)/2:])
		if len(left) == 0 || len(right) == 0 {
			return left + right
		} else {
			left_last := left[len(left)-1]
			right_first := right[0]
			if should_destroy(left_last, right_first){
				return left[0:len(left)-1] + right[1:]
			}else{
				return left + right
			}
		}
	} else if len(str) == 2 {
		if should_destroy(str[0], str[1]){
			return ""
		} else {
			return str
		}
	} else if len(str) == 1 {
		return str
	} else if len(str) == 0 {
		return ""
	}
	return str
}

func test(){
	fmt.Println(same_type('a', 'A'))
	fmt.Println(same_type('b', 'A'))
	fmt.Println(reverse_polarity('a', 'A'))
	fmt.Println(reverse_polarity('A', 'A'))
}
