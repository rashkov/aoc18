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
	fmt.Println(bfs(all_nodes))

}

func bfs(all_nodes map[string]Node) string{
	// Find the one that has no forward links, that's our "sink"
	// Perform BFS using an alphabetic queue
	var (
		bfs_alpha_queue []string
		sink_node       Node
		visited_links   map[string]bool = make(map[string]bool)
		path            string
	)
	for _, node := range all_nodes {
		if len(node.forward_links) == 0 {
			fmt.Println(node.id, " is the sink node.")
			sink_node = node
		}
	}

	Use(visited_links, path)

	// add the links to the queue
	for _, back_node := range sink_node.back_links {
		fmt.Println(back_node.id, " precedes it")
		bfs_alpha_queue = append(bfs_alpha_queue, back_node.id)
	}

	// sort the queue
	sort.Strings(bfs_alpha_queue)

	// remove the first element from the queue which has all its links visited
	for index, node_id := range bfs_alpha_queue {
		node := all_nodes[node_id];
		all_links_visited := true
		for _, linked_node := range node.back_links {
			visited, ok := visited_links[node_id + linked_node.id]
			if ok == false {
				all_links_visited = false
				visited_links[node_id + linked_node.id] = true
			}
			Use(linked_node, index, visited, all_links_visited)
		}
		if all_links_visited {
			// remove it
			fmt.Println("breaking")
			// break
		}else{
			// move onto the next one
			fmt.Println("continuing")
			// continue
		}
	}

	// repeat

	// every time you remove an element from the queue, add its symbol to our answer string

	// then reverse the string
	return path
}
