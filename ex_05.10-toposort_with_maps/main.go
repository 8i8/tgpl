package main

import "fmt"

type inmap map[string]bool

var prereqs map[string]inmap

func buildMap(values ...string) inmap {
	m := make(inmap)
	for _, v := range values {
		m[v] = false
	}
	return m
}

func init() {
	prereqs = make(map[string]inmap)
	prereqs["algorithms"] = buildMap([]string{"data structures"}...)
	prereqs["calculus"] = buildMap([]string{"linear algebra"}...)
	prereqs["compiler"] = buildMap([]string{
		"data structures",
		"formal languages",
		"computer organization",
	}...)
	prereqs["data structures"] = buildMap([]string{"discrete math"}...)
	prereqs["databases"] = buildMap([]string{"data structures"}...)
	prereqs["discrete math"] = buildMap([]string{"intro to programming"}...)
	prereqs["formal languages"] = buildMap([]string{"discrete math"}...)
	prereqs["networks"] = buildMap([]string{"operating systems"}...)
	prereqs["operating systems"] = buildMap([]string{
		"data structures",
		"computer organization",
	}...)
	prereqs["operating systems"] = buildMap([]string{
		"data structures",
		"computer organization",
	}...)
	prereqs["programming languages"] = buildMap([]string{
		"data structures",
		"computer organization",
	}...)
}

func main() {
	for i, course := range toposortmap(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func toposortmap(m map[string]inmap) []string {
	var output []string
	seen := make(map[string]bool)
	var visitAll func(items inmap)

	visitAll = func(items inmap) {
		for key, _ := range items {
			if !seen[key] {
				seen[key] = true
				visitAll(m[key])
				output = append(output, key)
			}
		}
	}

	for key, items := range m {
		visitAll(items)
		if !seen[key] {
			seen[key] = true
			output = append(output, key)
		}
	}
	return output
}
