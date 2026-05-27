package main

import (
	"fmt"
	"github.com/mieuxvoter/merit-profile-library-go/merit"
	"image/color"
)

func main() {
	proposals := []merit.Proposal{
		{
			Name:  "🙈",
			Tally: []uint64{0, 1, 0, 1, 2},
		},
		{
			Name:  "🙉",
			Tally: []uint64{1, 0, 2, 1, 0},
		},
		{
			Name:  "🙊",
			Tally: []uint64{1, 1, 0, 0, 2},
		},
	}

	svg, err := merit.RenderLinearProfileSVG(
		proposals,
		merit.WithBgColor(color.Black),
		merit.WithBestGradeOnLeft(true),
		merit.WithMedianLineStrategy(merit.DeadCenter),
		merit.WithGradesOutlines(
			[][]int{
				{3},
				{2},
				{1},
			},
		),
		merit.WithGradesOutlinesWidth(3.0),
		merit.WithGradesOutlinesColor(color.White),
	)
	if err != nil {
		panic(err)
	}

	fmt.Print(svg)
}
