package main

import (
	"fmt"
	"github.com/mieuxvoter/merit-profile-library-go/merit"
	"image/color"
)

func main() {
	proposals := []merit.Proposal{
		{
			Name:  "Une promenade sur la face cachée de la Lune",
			Tally: []uint64{8, 0, 10, 0, 9, 6, 7},
		},
		{
			Name:  "Une randonnée dans les monts de la Chartreuse",
			Tally: []uint64{5, 4, 10, 5, 5, 3, 8},
		},
		{
			Name:  "Un crossfit dans les égouts de Paris",
			Tally: []uint64{19, 5, 8, 3, 0, 3, 2},
		},
		{
			Name:  "Une baignade sur une plage nudiste avec Patrick Bruel, Morandini & Depardieu",
			Tally: []uint64{38, 0, 0, 0, 0, 0, 2},
		},
	}

	svg, err := merit.RenderLinearProfileSVG(
		proposals,
		merit.WithWidth(800),
		merit.WithHeight(400),
		merit.WithPadding(32),
		merit.WithVerticalSpacing(32),
		merit.WithHorizontalSpacing(8),
		merit.WithBgColor(color.NRGBA{R: 30, G: 20, B: 5, A: 128}),
		merit.WithMedianLineColor(color.NRGBA{R: 50, G: 50, B: 255, A: 255}),
		merit.WithMedianLineOutlineColor(color.NRGBA{R: 255, G: 255, B: 0, A: 120}),
		merit.WithTextColor(color.NRGBA{R: 220, G: 200, B: 200, A: 200}),
		merit.WithTextOutlineColor(color.NRGBA{R: 20, G: 20, B: 20, A: 200}),
		merit.WithGradesPalette([]color.Color{
			color.NRGBA{R: 0, G: 0, B: 0, A: 255},
			color.NRGBA{R: 36, G: 36, B: 36, A: 255},
			color.NRGBA{R: 73, G: 73, B: 73, A: 255},
			color.NRGBA{R: 109, G: 109, B: 109, A: 255},
			color.NRGBA{R: 146, G: 146, B: 146, A: 255},
			color.NRGBA{R: 219, G: 219, B: 219, A: 255},
			color.NRGBA{R: 255, G: 255, B: 255, A: 255},
		}),
		//merit.WithPatterns(…),
	)
	if err != nil {
		panic(err)
	}

	fmt.Print(svg)
}
