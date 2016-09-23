package main

import (
	"fmt"
	"sort"
)

func main() {

	m := mapmaker2(22)

	fmt.Println(m)

	ms := make([]string, len(m))

	i := 0
	for names := range m {
		ms[i] = names
		i++
	}

	sort.Strings(ms)

	for sortname := range ms {
		fmt.Println(m[ms[sortname]])
		fmt.Printf("%T %v", m[ms[sortname]], m[ms[sortname]])
	}

	fmt.Println(ms)
}

func mapmaker2(i int) map[string]int {

	m := make(map[string]int)

	m["key1"] = i + 1
	m["key4"] = i + 4
	m["key2"] = i + 2
	m["key5"] = i + 5
	m["key3"] = i + 3

	fmt.Println(m)

	return m
}
