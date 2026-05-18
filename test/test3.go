package main

import (
	"fmt"
	"github.com/mieuxvoter/merit-profile-library-go/merit"
	"image/color"
)

func main() {
	proposals := []merit.Proposal{
		{
			Name:  "[OKFN] Open Knowledge Foundation",
			Tally: []uint64{0, 0, 0, 0, 4},
		},
	}

	svg, err := merit.RenderLinearProfileSVG(
		proposals,
		merit.WithBgColor(color.Black),
		merit.WithWidth(1024),
		merit.WithGradeHeight(96),
		merit.WithProposalFontSize("3em"),
		merit.WithTallyFontSize("3em"),
	)
	if err != nil {
		panic(err)
	}

	fmt.Print(svg)
}
