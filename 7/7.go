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
}

var all_nodes map[string]Node = make(map[string]Node)

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
			preceding_node         Node
			following_node         Node
			ok                     bool
			already_linked_forward bool
			already_linked_back    bool
		)
		if preceding_node, ok = all_nodes[preceding_node_id]; ok != true {
			preceding_node = Node{preceding_node_id, nil, nil}
			all_nodes[preceding_node_id] = preceding_node
		}
		if following_node, ok = all_nodes[following_node_id]; ok != true {
			following_node = Node{following_node_id, nil, nil}
			all_nodes[following_node_id] = following_node
		}
		for _, link := range preceding_node.forward_links {
			if link.id == following_node_id {
				already_linked_forward = true
			}
		}
		if !already_linked_forward {
			preceding_node.forward_links = append(preceding_node.forward_links, &following_node)
			// following_node.back_links = append(following_node.back_links, &preceding_node)
		}
		for _, back_link := range following_node.back_links {
			if back_link.id == preceding_node_id {
				already_linked_back = true
			}
		}
		if !already_linked_back {
			// preceding_node.forward_links = append(preceding_node.forward_links, &following_node)
			following_node.back_links = append(following_node.back_links, &preceding_node)
		}
		all_nodes[preceding_node_id] = preceding_node
		all_nodes[following_node_id] = following_node
	}

	fmt.Println("Part 1: ", strings.Join(bfs(all_nodes), ""))
}

func bfs(all_nodes map[string]Node) []string {
	// Perform BFS using an alphabetic queue
	var (
		bfs_alpha_queue []string
		path            []string
	)

	// Add the all nodes to our queue
	for _, node := range all_nodes {
		bfs_alpha_queue = append(bfs_alpha_queue, node.id)
	}

	for len(bfs_alpha_queue) > 0 {
		// sort the queue
		sort.Strings(bfs_alpha_queue)

		// remove the first element from the queue which has all its back_links visited
		// (On the first pass, that should be our start_node)
		var current_node Node
		var available_nodes []Node
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
			if all_back_links_resolved {
				// save it for processing
				if current_node.id == ""{
					current_node = node
				}
				available_nodes = append(available_nodes, node)
			}
		}

		process_step(&bfs_alpha_queue, current_node.id, &path)

	}

	return path
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
func process_step(queue *[]string, node_id string, path *[]string){
	// Add it to path
	*path = append(*path, node_id)
	// Remove it from queue
	index := index_of_node(*queue, node_id)
	*queue = append((*queue)[:index], (*queue)[(index+1):]...)
}
