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
	// polymer = []byte("dabAcCaCBAcCcaDA")
	polymer_str := strings.TrimSpace(string(polymer))
	removed_lengths := make(map[byte]int)

	_, orig_reduced_length := reduce_until_done(polymer_str)
	fmt.Println("Part 1", orig_reduced_length)

	for _, letter := range get_all_types(polymer_str) {
		str := remove_letter(polymer_str, letter)
		_, reduced_length := reduce_until_done(str)
		removed_lengths[letter] = reduced_length
	}

	var smallest_letter byte
	var smallest_length int = orig_reduced_length
	for letter, length_without_letter := range removed_lengths {
		if length_without_letter < smallest_length {
			smallest_letter = letter
			smallest_length = length_without_letter
		}
	}
	fmt.Println("Part 2: Smallest after removing\n", smallest_length, string(smallest_letter))
}

func remove_letter(str string, remove_letter byte) string {
	var str_without_type string
	for _, letter := range str {
		if byte(letter) != remove_letter && string(letter) != strings.ToUpper(string(remove_letter)) {
			str_without_type += string(letter)
		}
	}
	return str_without_type
}

func reduce_until_done(polymer_str string) (string, int) {
	res1 := reduce(polymer_str)
	len1 := len(res1)
	res2 := reduce(res1)
	len2 := len(res2)
	for len2 < len1 {
		res1 = res2
		len1 = len2
		res2 = reduce(res1)
		len2 = len(res2)
	}
	return res2, len2
}

func get_all_types(str string) []byte {
	var types []byte
	keys := make(map[byte]bool)
	for _, s := range str {
		keys[strings.ToLower(string(s))[0]] = true
	}
	for key, _ := range keys {
		types = append(types, key)
	}
	return types
}

func should_destroy(a byte, b byte) bool {
	return same_type(a, b) && reverse_polarity(a, b)
}

func same_type(a byte, b byte) bool {
	return strings.ToUpper(string(a)) == strings.ToUpper(string(b))
}

func reverse_polarity(a byte, b byte) bool {
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
			if should_destroy(left_last, right_first) {
				return left[0:len(left)-1] + right[1:]
			} else {
				return left + right
			}
		}
	} else if len(str) == 2 {
		if should_destroy(str[0], str[1]) {
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

func test() {
	fmt.Println(same_type('a', 'A'))
	fmt.Println(same_type('b', 'A'))
	fmt.Println(reverse_polarity('a', 'A'))
	fmt.Println(reverse_polarity('A', 'A'))
	fmt.Println(remove_letter("aaabcccc", 'b'))
}
