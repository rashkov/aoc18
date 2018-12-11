package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
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
	id            string
	forward_links []*Node
	back_links    []*Node
	locked_by     int
}

var all_nodes map[string]*Node = make(map[string]*Node)

const NUM_WORKERS = 5

var workers [NUM_WORKERS]int

func main() {
	f, err := os.Open("./input.txt")
	check(err)
	defer f.Close()
	var input []string
	scanner := bufio.NewScanner(f)
	defer check(scanner.Err())
	for scanner.Scan() {
		input = append(input, strings.TrimSpace(scanner.Text()))
	}
	for _, line := range input {
		parsed := regexp.MustCompile(` (.) .*? (.) `).FindStringSubmatch(line)
		preceding_node_id := parsed[1]
		following_node_id := parsed[2]

		var (
			preceding_node         *Node
			following_node         *Node
			ok                     bool
			already_linked_forward bool
			already_linked_back    bool
		)
		if preceding_node, ok = all_nodes[preceding_node_id]; ok != true {
			preceding_node = &Node{preceding_node_id, nil, nil, -1}
			all_nodes[preceding_node_id] = preceding_node
		}
		if following_node, ok = all_nodes[following_node_id]; ok != true {
			following_node = &Node{following_node_id, nil, nil, -1}
			all_nodes[following_node_id] = following_node
		}
		for _, link := range preceding_node.forward_links {
			if link.id == following_node_id {
				already_linked_forward = true
			}
		}
		if !already_linked_forward {
			preceding_node.forward_links = append(preceding_node.forward_links, following_node)
		}
		for _, back_link := range following_node.back_links {
			if back_link.id == preceding_node_id {
				already_linked_back = true
			}
		}
		if !already_linked_back {
			following_node.back_links = append(following_node.back_links, preceding_node)
		}
		all_nodes[preceding_node_id] = preceding_node
		all_nodes[following_node_id] = following_node
	}

	path, time := bfs(all_nodes)
	fmt.Println("Part 1: ", strings.Join(path, ""))
	fmt.Println("Part 2: ", time)
}

func bfs(all_nodes map[string]*Node) ([]string, int) {
	// Perform BFS using an alphabetic queue
	var (
		bfs_alpha_queue []string
		path            []string
		time            int
	)

	// Add the all nodes to our queue & sort it
	for _, node := range all_nodes {
		bfs_alpha_queue = append(bfs_alpha_queue, node.id)
	}
	sort.Strings(bfs_alpha_queue)

	for time = 0; len(bfs_alpha_queue) > 0; time++ {
		decrement_workers()
		complete_steps(&bfs_alpha_queue, &path, all_nodes)

		// remove the first element from the queue which has all its back_links visited
		// (On the first pass, that should be our start_node)
		var current_node *Node
		var available_nodes []*Node
		for _, node_id := range bfs_alpha_queue {
			node := all_nodes[node_id]
			all_back_links_resolved := true

			for _, back_linked_node := range node.back_links {
				ok := index_of_node(path, back_linked_node.id)
				if ok == -1 {
					all_back_links_resolved = false
					break
				}
			}
			if all_back_links_resolved && node.locked_by == -1 {
				// save it for processing
				if current_node == nil {
					current_node = node
				}
				available_nodes = append(available_nodes, node)
			}
		}

		distribute_steps(available_nodes)
	}

	return path, time-1
}

func distribute_steps(available_steps []*Node) {
	take := num_available_workers()
	if take > len(available_steps) {
		take = len(available_steps)
	}
	for _, step := range available_steps[0:take] {
		for worker_id, w := range workers {
			if w == 0 {
				workers[worker_id] += calculate_step_cost(step.id)
				step.locked_by = worker_id
				//fmt.Println("Giving step", step.id, "to worker", worker_id, "at a cost of", workers[worker_id], "locking it to", step.locked_by)
				break
			}
		}
	}
}

func complete_steps(queue *[]string, path *[]string, all_nodes map[string]*Node) {
	// run through the queue, find any locked ones, see if their worker is at zero, then process it
	for _, step_id := range *queue {
		step := all_nodes[step_id]
		if step.locked_by != -1 {
			worker_id := step.locked_by
			if workers[worker_id] == 0 {
				process_step(queue, step.id, path)
			}
		}
	}
}

func index_of_node(node_list []string, node_id string) int {
	var position = -1
	for index, queued := range node_list {
		if queued == node_id {
			position = index
		}
	}
	return position
}

/* process_step()
 * removes a step from the queue and adds it to our solution
 * The passed node must be eligible. Its backlinks must all be resolved
 */
func process_step(queue *[]string, node_id string, path *[]string) {
	// Add it to path
	*path = append(*path, node_id)
	// Remove it from queue
	index := index_of_node(*queue, node_id)
	*queue = append((*queue)[:index], (*queue)[(index+1):]...)
}

func calculate_step_cost(step_id string) int {
	return int(step_id[0]) - 64 + 60
}

func decrement_workers() {
	for i, _ := range workers {
		if workers[i] > 0 {
			workers[i]--
		}
	}
	//fmt.Println(workers)
}

func num_available_workers() int {
	count := 0
	for _, w := range workers {
		if w == 0 {
			count += 1
		}
	}
	return count
}
