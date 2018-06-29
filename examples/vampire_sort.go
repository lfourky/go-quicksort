package main

import (
	"github.com/lfourky/go-quicksort"
)

type vampires []vampire

type vampire struct {
	name         string
	prefersWomen bool
	killCount    int
}

// Chooose your sorting key (wisely)
func (v vampires) LessEqual(i, j int) bool {
	return v[i].killCount <= v[j].killCount
}

func (v vampires) Len() int {
	return len(v)
}

func (v vampires) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}

func main() {
	vampz := make(vampires, 0, 2)
	vampz = append(vampz,
		vampire{name: "Dracula", prefersWomen: true, killCount: 100001},
		vampire{name: "Sava Savanović", prefersWomen: false, killCount: 3},
	)
	sort.QuickSort(vampz, false)
}
