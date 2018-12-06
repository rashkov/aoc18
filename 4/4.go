package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
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

type Guard struct {
	guard_id     int64
	total_asleep int
	minutes      [60]int
}

func main() {
	f, err := os.Open("./input.txt")
	check(err)
	defer f.Close()

	var input []string
	guards := make(map[int64]Guard)
	var current_guard int64
	var awoke_minute, asleep_minute int

	scanner := bufio.NewScanner(f)
	defer check(scanner.Err())
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	sort.Strings(input)
	for _, txt := range input {
		if match, _ := regexp.MatchString("Guard", txt); match {
			parsed := regexp.MustCompile(`#(\d+)`).FindStringSubmatch(txt)
			guard_id, err := strconv.ParseInt(parsed[1], 10, 64)
			check(err)
			//fmt.Println("Guard!", guard_id)
			if guards[guard_id].guard_id == 0 {
				guards[guard_id] = Guard{guard_id, 0, [60]int{}}
			}
			current_guard = guard_id
		} else if match, _ := regexp.MatchString("asleep", txt); match {
			parsed := regexp.MustCompile(`\d+:(\d+)`).FindStringSubmatch(txt)
			asleep_minute64, err := strconv.ParseInt(parsed[1], 10, 8)
			check(err)
			asleep_minute = int(asleep_minute64)
			//fmt.Println("asleep", asleep_minute)
		} else if match, _ := regexp.MatchString("wakes", txt); match {
			parsed := regexp.MustCompile(`\d+:(\d+)`).FindStringSubmatch(txt)
			awoke_minute64, err := strconv.ParseInt(parsed[1], 10, 8)
			check(err)
			awoke_minute = int(awoke_minute64)
			//fmt.Println("awake", awoke_minute)
			guard := guards[current_guard]
			guard.total_asleep = guard.total_asleep + awoke_minute - asleep_minute
			minutes := guard.minutes
			for i := asleep_minute; i < awoke_minute; i++ {
				minutes[i] += 1
			}
			guard.minutes = minutes
			guards[current_guard] = guard
		}
	}

	var sleepiest_guard = Guard{}
	for _, guard := range guards {
		if guard.total_asleep > sleepiest_guard.total_asleep {
			sleepiest_guard = guard
		}
	}
	fmt.Println("Sleepiest guard id:", sleepiest_guard.guard_id)
	fmt.Println("Sleepiest guard sleep total:", sleepiest_guard.total_asleep)
	var sleepiest_minute, sleepiest_total int
	for minute, sleep_time := range(sleepiest_guard.minutes){
		if sleep_time > sleepiest_total {
			sleepiest_total = sleep_time
			sleepiest_minute = minute
		}
	}
	fmt.Printf("He slept %d minutes during minute %d!\n", sleepiest_total, sleepiest_minute)

	fmt.Println("Part 1 (id times minute)", sleepiest_guard.guard_id * int64(sleepiest_minute))

	var num_times_asleep_during_minute, most_frequent_minute, most_frequent_guard_id int
	for _, guard := range guards {
		minutes := guard.minutes
		for minute, times_asleep := range(minutes) {
			if times_asleep > num_times_asleep_during_minute {
				num_times_asleep_during_minute = times_asleep
				most_frequent_minute = minute
				most_frequent_guard_id = int(guard.guard_id)
			}
		}
	}
	fmt.Printf("Guard %d slept %d times during minute %d\n",
		most_frequent_guard_id,
		num_times_asleep_during_minute,
		most_frequent_minute)
	fmt.Println("Part 2 (guard_id * most_frequent_minute):", most_frequent_guard_id*most_frequent_minute)
}

//var txt []string
// data_struct := make(map[string]bool)
// txt = strings.Split(scanner.Text(), "\n")
// actionRE := regexp.MustCompile(`#(\d+)`)
// actionRE.FindStringSubmat(in:(\d+)ut)// .ParseInt(parsed_coords[1], 10, 64)
