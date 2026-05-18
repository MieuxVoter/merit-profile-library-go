package main

import (
	"fmt"
	"github.com/mieuxvoter/merit-profile-library-go/merit"
	"image/color"
)

func main() {
	proposals := []merit.Proposal{
		{
			Name:  "Pizza 4 Dimensions",
			Tally: []uint64{5, 4, 11},
		},
		{
			Name:  "Lasagnes Assange",
			Tally: []uint64{9, 5, 6},
		},
		{
			Name:  "Jurassique Pâtes",
			Tally: []uint64{14, 0, 6},
		},
	}

	svg, err := merit.RenderLinearProfileSVG(
		proposals,
		merit.WithBgColor(color.Black),
	)
	if err != nil {
		panic(err)
	}

	fmt.Print(svg)
}
