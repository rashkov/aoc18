package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
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

type Node struct {
	id           int
	num_children int
	num_meta     int
	children     []*Node
	meta         []int
}

var sum_meta int

func main() {
	input, err := ioutil.ReadFile("./input.txt")
	check(err)
	str_arr := strings.Split(strings.TrimSpace(string(input)), " ")

	var nums []int
	for _, str := range str_arr {
		num, err := strconv.ParseInt(str, 10, 64)
		check(err)
		nums = append(nums, int(num))
	}
	var id = 65
	read_cursor, node_ptr := parse(&nums, &id, 0)

	fmt.Println("Read", read_cursor+1, "items") // +1 to include zeroth element
	// fmt.Println(node_ptr)
	fmt.Println("Part 1:", sum_meta)
	fmt.Println("Part 2:", calc_node_score(node_ptr))
}

func parse(input_ptr *[]int, current_id *int, start int) (read_cursor int, self_ptr *Node) {
	input := *input_ptr
	read_cursor = start + 1
	num_children := input[start]
	num_meta := input[read_cursor]

	self := Node{*current_id, num_children, num_meta, nil, nil}
	(*current_id)++

	var child_ptr *Node
	for i := 0; i < num_children; i++ {
		read_cursor, child_ptr = parse(input_ptr, current_id, read_cursor+1)
		self.children = append(self.children, child_ptr)
	}

	for i := 0; i < num_meta; i++ {
		read_cursor = read_cursor + 1
		sum_meta += input[read_cursor]
		self.meta = append(self.meta, input[read_cursor])
	}
	// fmt.Println("Read cursor:", read_cursor)
	// fmt.Println("Node:", self)
	return read_cursor, &self
}

func calc_node_score(node *Node) int {
	sum := 0
	num_children := len(node.children)
	if num_children == 0 {
		for _, m := range node.meta {
			sum += m
		}
	} else {
		for _, m := range node.meta {
			if m == 0 {
				continue
			}
			// m is a a reference into node's children array,
			// starting w/ 1 for 1st child, 2 for 2nd, etc.
			if num_children >= m {
				child := node.children[m-1]
				sum += calc_node_score(child)
			}
		}
	}
	return sum
}
