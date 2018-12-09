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
	f, err := os.Open("./test_input.txt")
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
			preceding_node Node
			following_node Node
			ok             bool
			already_linked bool
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
				already_linked = true
			}
		}
		if !already_linked {
			preceding_node.forward_links = append(preceding_node.forward_links, &following_node)
			following_node.back_links = append(following_node.back_links, &preceding_node)
			all_nodes[preceding_node_id] = preceding_node
			all_nodes[following_node_id] = following_node
		}
	}

	// Find the one that has no forward links, that's our "sink"
	var sink_node Node
	for _, node := range all_nodes {
		if len(node.forward_links) == 0 {
			fmt.Println(node.id, " is the sink node.")
			sink_node = node
		}
	}

	// Run BFS backwards from the sink to the roots
	fmt.Println(bfs(all_nodes, sink_node))

}

func bfs(all_nodes map[string]Node, start_node Node) []string {
	// Perform BFS using an alphabetic queue
	var (
		bfs_alpha_queue []string
		visited_links   map[string]bool = make(map[string]bool)
		path            []string
	)

	// Add the start_node to our queue
	bfs_alpha_queue = append(bfs_alpha_queue, start_node.id)

	for len(bfs_alpha_queue) > 0{
		// sort the queue
		sort.Strings(bfs_alpha_queue)
		fmt.Println(bfs_alpha_queue)

		// remove the first element from the queue which has all its links visited
		// (On the first pass, that should be our start_node)
		var current_node Node
		for index, node_id := range bfs_alpha_queue {
			node := all_nodes[node_id]
			fmt.Println("starting with ", node_id, " has ", len(node.forward_links), " links")
			all_links_visited := true
			for _, linked_node := range node.forward_links {
				_, ok := visited_links[linked_node.id + node.id]
				if ok == false {
					all_links_visited = false
					visited_links[node_id+linked_node.id] = true
					fmt.Println("naw", linked_node.id)
				}else{
					fmt.Println("visited from ", linked_node.id)
				}
			}
			if all_links_visited {
				fmt.Println("all links visited")
				// save it for processing
				current_node = node
				// remove it from queue
				bfs_alpha_queue = append(bfs_alpha_queue[:index], bfs_alpha_queue[(index+1):]...)
				break
			} else {
				// keep looking
				fmt.Println("has unresolved links, moving on")
				continue
			}
		}

		fmt.Println("Current node: ", current_node)
		fmt.Println("BFS queue: ", bfs_alpha_queue)

		// Got a node to process
		// every time we remove an element from the queue, add its symbol to the answer string
		path = append(path, current_node.id)

		// add its links to the queue, marking the links as visited
		// NOTE: Don't add a node if it's already in the queue
		for _, linked_node := range current_node.back_links {
			fmt.Println(linked_node.id, " precedes it")
			visited_links[current_node.id + linked_node.id] = true
			var already_in_queue = false
			for _, queued := range bfs_alpha_queue {
				if queued == linked_node.id{
					already_in_queue = true
				}
			}
			if !already_in_queue{
				bfs_alpha_queue = append(bfs_alpha_queue, linked_node.id)
			}
		}
		fmt.Println("BFS queue: ", bfs_alpha_queue)

		// repeat
	}

	// then reverse the string
	return path
}
