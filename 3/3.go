package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"regexp"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

const LENGTH = 10

func main() {
	f, err := os.Open("./test_input.txt")
	check(err)
	defer f.Close()

	type SquareInch struct {
		num_claims int
		claimant_ids []string
	}

	var matx [LENGTH][LENGTH] SquareInch
	for i := 0; i < LENGTH; i++ {
		for j := 0; i < LENGTH; i++ {
			matx[i][j] = SquareInch{ 0, nil}
		}
	}

	valid_claims := make(map[string]bool)

	scanner := bufio.NewScanner(f)
	defer check(scanner.Err())
	for scanner.Scan() {
		txt := strings.Split(scanner.Text(), " ")
		claimant_id := txt[0][1:]
		coords := txt[2]
		extents := txt[3]

		valid_claims[claimant_id] = true

		coordRE := regexp.MustCompile(`(\d+),(\d+)`)
		parsed_coords := coordRE.FindStringSubmatch(coords)
		x_coord, err := strconv.ParseInt(parsed_coords[1], 10, 64)
		check(err)
		y_coord, err := strconv.ParseInt(parsed_coords[2], 10, 64)
		check(err)

		extentsRE := regexp.MustCompile(`(\d+)x(\d+)`)
		parsed_extents := extentsRE.FindStringSubmatch(extents)
		x_extent, err := strconv.ParseInt(parsed_extents[1], 10, 64)
		check(err)
		y_extent, err := strconv.ParseInt(parsed_extents[2], 10, 64)
		check(err)

		fmt.Printf("#%s @ %d,%d: %dx%d\n", claimant_id, x_coord, y_coord, x_extent, y_extent)
		// fmt.Println("x: ", x_coord, "y: ", y_coord)
		// fmt.Println("width: ", x_extent, "height: ", y_extent)
		for i := int64(0); i < x_extent; i++ {
			for j := int64(0); j < y_extent; j++ {
				x := x_coord + i
				y := y_coord + j
				matx[x][y].num_claims += 1
				// add the new claim
				matx[x][y].claimant_ids = append(matx[x][y].claimant_ids, claimant_id)
				// Invalidate claims
				if len(matx[x][y].claimant_ids) > 1 {
					// Invalidate claim
					for _, claim_id := range matx[x][y].claimant_ids{
						valid_claims[claim_id] = false
					}
				}
			}
		}
	}

	count := 0
	for i := 0; i < LENGTH; i++ {
		for j := 0; j < LENGTH; j++ {
			if matx[i][j].num_claims > 1{
				count++
			}
		}
	}

	// fmt.Println("Part 1: ", count)

	// for claimant_id, valid := range(valid_claims){
	// 	if valid{
	// 		fmt.Println("Part 2:", claimant_id)
	// 	}
	// }

	for j := 0; j < LENGTH; j++ {
		for i := 0; i < LENGTH; i++ {
			cell := matx[i][j]
			if len(cell.claimant_ids) == 1 {
				fmt.Print(cell.claimant_ids[0] + " ")
			}else if len(cell.claimant_ids) == 0{
				fmt.Print(". ")
			}else{
				fmt.Print("X ")
			}
			//fmt.Println(matx[i][j])
		}
		fmt.Println()
	}
}
