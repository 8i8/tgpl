package main

import (
	"errors"
	"fmt"
	"os"
)

type inmap map[string]bool

type graph map[string]inmap

var prereqs graph

func buildMap(values ...string) inmap {
	m := make(inmap)
	for _, v := range values {
		m[v] = true
	}
	return m
}

func init() {
	prereqs = make(graph)
	prereqs["algorithms"] = buildMap("data structures")
	prereqs["calculus"] = buildMap("linear algebra")
	prereqs["compiler"] = buildMap(
		"data structures",
		"formal languages",
		"computer organization",
	)
	prereqs["data structures"] = buildMap("discrete math")
	prereqs["databases"] = buildMap("data structures")
	prereqs["discrete math"] = buildMap("intro to programming")
	prereqs["formal languages"] = buildMap("discrete math")
	prereqs["networks"] = buildMap("operating systems")
	prereqs["operating systems"] = buildMap(
		"data structures",
		"computer organization")
	prereqs["operating systems"] = buildMap(
		"data structures",
		"computer organization")
	prereqs["programming languages"] = buildMap(
		"data structures",
		"computer organization")
	prereqs["linear algebra"] = buildMap("calculus")
}

func main() {
	res, err := toposortmap(prereqs)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	for i, course := range res {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func (g graph) isCycle(c1, c2 string) bool {
	return g[c1][c2] && g[c2][c1]
}

var errCycle = errors.New("cyclic dependency")

func toposortmap(g graph) ([]string, error) {
	fname := "toposortmap"
	var output []string
	seen := make(map[string]bool)
	var visitAll func(key string, items inmap) error

	visitAll = func(key string, items inmap) error {
		fname := "visitAll"
		for item, _ := range items {
			if g.isCycle(key, item) {
				return fmt.Errorf(fname+
					": %q <-> %q: %w", key, item, errCycle)
			}
			if !seen[item] {
				seen[item] = true
				visitAll(key, g[item])
				output = append(output, item)
			}
		}
		return nil
	}

	for key, items := range g {
		err := visitAll(key, items)
		if err != nil {
			return nil, fmt.Errorf(fname+": %w", err)
		}
		if !seen[key] {
			seen[key] = true
			output = append(output, key)
		}
	}

	return output, nil
}
