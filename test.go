package main

import (
	"fmt"
	"github.com/mieuxvoter/merit-profile-library-go/merit"
)

func main() {
	proposals := []merit.Proposal{
		{
			Name:  "Alice the wonderful napping kangaroo 🦘 of the Æther",
			Tally: []uint64{4, 0, 3, 7},
		},
		{
			Name:  "Dominique",
			Tally: []uint64{5, 6, 1, 2},
		},
		{
			Name:  "Théo 🗳",
			Tally: []uint64{3, 3, 2, 6},
		},
	}

	svg, err := merit.RenderLinearProfileSVG(proposals)
	if err != nil {
		panic(err)
	}

	fmt.Print(svg)
}
