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

	// Find the one that has no back links, that's our "root"
	// In case of multiple root nodes, choose the alphabetically first one
	var root_nodes []Node
	for _, node := range all_nodes {
		if len(node.back_links) == 0 {
			fmt.Println(node.id, " is a root node.")
			root_nodes = append(root_nodes, node)
		}
	}
	fmt.Println("root nodes: ", root_nodes)
	// Use(root_nodes)
	for _, root_node := range root_nodes {
		path := bfs(all_nodes, root_node)
		if len(path) != 0 {
			fmt.Println("Found a path!")
			fmt.Println(path)
			break
		}
	}
	// for _, node := range all_nodes {
	// 	path := bfs(all_nodes, node)
	// 	if len(path) != 0 {
	// 		fmt.Println("Found a path!")
	// 		fmt.Println(path)
	// 		break
	// 	}
	// }
}

func bfs(all_nodes map[string]Node, start_node Node) []string {
	// Perform BFS using an alphabetic queue
	var (
		bfs_alpha_queue []string
		resolved_nodes  map[string]bool = make(map[string]bool)
		path            []string
	)

	// Add the start_node to our queue
	bfs_alpha_queue = append(bfs_alpha_queue, start_node.id)

	for len(bfs_alpha_queue) > 0 {
		// sort the queue
		sort.Strings(bfs_alpha_queue)
		fmt.Println(bfs_alpha_queue)

		// remove the first element from the queue which has all its back_links visited
		// (On the first pass, that should be our start_node)
		var current_node Node
		for index, node_id := range bfs_alpha_queue {
			node := all_nodes[node_id]
			fmt.Println("starting with ", node_id, " has ", len(node.back_links), " links")
			all_back_links_resolved := true

			for _, back_linked_node := range node.back_links {
				_, ok := resolved_nodes[back_linked_node.id]
				if ok == false{
					all_back_links_resolved = false
					fmt.Println("can't use, its back link not resolved: ", back_linked_node.id)
					break
				} else {
					fmt.Println("back link resolved: ", back_linked_node.id)
				}
			}
			if all_back_links_resolved {
				fmt.Println("all back links resolved, process it: ", node.id)
				// save it for processing
				current_node = node
				// remove it from queue
				bfs_alpha_queue = append(bfs_alpha_queue[:index], bfs_alpha_queue[(index+1):]...)
				break
			} else {
				// keep looking
				continue
			}
		}

		fmt.Println("Current node: ", current_node)
		fmt.Println("BFS queue: ", bfs_alpha_queue)
		if current_node.id == "" {
			fmt.Println("Hit a dead end!")
			fmt.Println("")
			return []string{}
		}

		// Got a node to process
		// every time we remove an element from the queue, add its symbol to the answer string
		path = append(path, current_node.id)
		// and mark it as resolved
		// (it is resolved because its backlinks are resolved, which is why we dequeued it)
		resolved_nodes[current_node.id] = true

		// add its links to the queue
		// NOTE: Don't add a node if it's already in the queue
		for _, linked_node := range current_node.forward_links {
			fmt.Println(linked_node.id, " follows it")
			var already_in_queue = false
			for _, queued := range bfs_alpha_queue {
				if queued == linked_node.id {
					already_in_queue = true
				}
			}
			if !already_in_queue {
				bfs_alpha_queue = append(bfs_alpha_queue, linked_node.id)
			}
		}
		fmt.Println("BFS queue: ", bfs_alpha_queue)

		// repeat
	}

	return path
}
