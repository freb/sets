package main

import (
	"flag"
	"fmt"
	"strings"

	mapset "github.com/deckarep/golang-set"
)

func main() {
	a := flag.String("a", "", "set A (comma-separated list")
	b := flag.String("b", "", "set B (comma-separated list")
	sep := flag.String("s", ",", "list separator")
	difference := flag.Bool("d", false, "difference (A \\ B)")
	intersection := flag.Bool("i", false, "insersection (A ∩ B)")
	union := flag.Bool("u", false, "union (A ∪ B)")

	flag.Parse()

	setA := mapset.NewSet()
	for _, el := range strings.Split(*a, *sep) {
		setA.Add(el)
	}
	setB := mapset.NewSet()
	for _, el := range strings.Split(*b, *sep) {
		setB.Add(el)
	}

	if !*difference && !*intersection && !*union {
		fmt.Printf("Set A: %v\n", setA)
		fmt.Printf("Set B: %v\n", setB)
	}

	if *difference {
		result := setA.Difference(setB)
		fmt.Printf("\\: %v\n", result)
	}

	if *intersection {
		result := setA.Intersect(setB)
		fmt.Printf("∩: %v\n", result)
	}

	if *union {
		result := setA.Union(setB)
		fmt.Printf("∪: %v\n", result)
	}
}
