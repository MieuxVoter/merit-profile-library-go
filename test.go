package main

import (
	"fmt"
	"github.com/mieuxvoter/merit-profile-library-go/merit"
)

func main() {
	proposals := []merit.Proposal{
		{
			Tally: []uint{3, 7, 2},
			Name:  "Dominique",
		},
		{
			Tally: []uint{2, 4, 6},
			Name:  "Théo 🗳",
		},
		{
			Tally: []uint{5, 0, 7},
			Name:  "Alice the wonderful napping kangaroo 🦘 of the Æther",
		},
	}

	svg, err := merit.RenderSvg(proposals)
	if err != nil {
		panic(err)
	}

	fmt.Print(svg)
}
