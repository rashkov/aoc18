package main

import (
	"fmt"
	"math"
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

const INPUT = 793031
const INPUT2 = 793031

// const INPUT = 2018
// const INPUT2 = 59414
const FUDGE_FACTOR = 30

const NUM_ELVES = 2

type Kitchen struct {
	elves               [NUM_ELVES]int
	recipes             [(INPUT + 10) * 2 * FUDGE_FACTOR]int
	latest_recipe_index int
}

func (kitchen *Kitchen) assign_new_tasks() {
	for elf_id, elfs_current_recipe := range kitchen.elves {
		offset := kitchen.recipes[elfs_current_recipe] + 1
		new_current_recipe := int(math.Mod(float64(elfs_current_recipe+offset), float64(kitchen.latest_recipe_index+1)))
		kitchen.elves[elf_id] = new_current_recipe
	}
}
func (kitchen *Kitchen) create_new_recipes() {
	elem := strconv.FormatInt(int64(kitchen.sum()), 10)
	var new_recipes []int
	for i := 0; i < len(elem); i++ {
		new_recipe, err := strconv.ParseInt(string(elem[i]), 10, 64)
		check(err)
		new_recipes = append(new_recipes, int(new_recipe))
	}
	copy(kitchen.recipes[kitchen.latest_recipe_index+1:], new_recipes)

	new_latest_recipe_index := kitchen.latest_recipe_index + len(new_recipes)
	kitchen.latest_recipe_index = new_latest_recipe_index
}

func (kitchen *Kitchen) initialize() {
	kitchen.recipes[0] = 3
	kitchen.recipes[1] = 7
	kitchen.latest_recipe_index = 1
	for n, _ := range kitchen.recipes {
		if n > kitchen.latest_recipe_index {
			kitchen.recipes[n] = -1
		}
	}
	for index, _ := range kitchen.elves {
		kitchen.elves[index] = index
	}
}

func (kitchen *Kitchen) sum() int {
	sum := 0
	for _, elfs_current_recipe := range kitchen.elves {
		sum += kitchen.recipes[elfs_current_recipe]
	}
	return sum
}

func main() {
	var kitchen Kitchen
	kitchen.initialize()
	for i := 0; i < (INPUT+10)*FUDGE_FACTOR; i++ {
		kitchen.create_new_recipes()
		kitchen.assign_new_tasks()
		// fmt.Println(kitchen.recipes[0 : kitchen.latest_recipe_index+1])

		input_str := strconv.FormatInt(int64(INPUT2), 10)
		input_str_len := len(input_str)
		if kitchen.latest_recipe_index > input_str_len {
			var output_str_arr []string
			for _, el := range kitchen.recipes[i : i+input_str_len] {
				output_str_arr = append(output_str_arr, strconv.FormatInt(int64(el), 10))
			}
			output_str := strings.Join(output_str_arr, "")
			//fmt.Println("Comparing", output_str, input_str)
			if output_str == input_str {
				fmt.Println("Part 2, found string after", i, "recipes")
			}
		}

		if i == INPUT+10-1 {
			var last_10 []string
			for _, k := range kitchen.recipes[INPUT : INPUT+10] {
				last_10 = append(last_10, strconv.FormatInt(int64(k), 10))
			}
			fmt.Println("Part 1:", strings.Join(last_10, ""))
		}
	}

	//fmt.Println("Sum of all:", strings.Join(all, ""))

}
