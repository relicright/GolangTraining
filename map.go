package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println(mapmaker("key", 1))

	m2 := mapmaker("key2", 23)
	fmt.Println(m2["key2"])

	states := make(map[string]string)
	fmt.Println(states)

	states["WA"] = "Washington"
	states["TX"] = "Texas"
	states["CA"] = "California"

	fmt.Println(states)

	california := states["CA"]
	fmt.Println(california)

	delete(states, "CA")
	fmt.Println(states)

	states["NY"] = "New York"

	for k, v := range states {
		fmt.Printf("%v: %v\n", k, v)
	}

	keys := make([]string, len(states))

	i := 0
	for k := range states {
		keys[i] = k
		i++
	}

	sort.Strings(keys)
	fmt.Println("\nSorted")

	for i := range keys {
		fmt.Println(states[keys[i]])
	}
}

func mapmaker(n string, i int) map[string]int {

	m := make(map[string]int)
	m[n] = i
	return m
}
