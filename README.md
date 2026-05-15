# Merit Profile Generation for Golang

Generate merit profiles (in SVG), for use for example in [Majority Judgment] polls.

[Majority Judgment]: https://mieuxvoter.fr

> [!INFO]
> This library focuses on rendering the merit profiles, not ranking the proposals.
> If you want to rank the proposals as well, there is [a library](https://github.com/MieuxVoter/majority-judgment-library-go) for that. 

## Usage

```shell
go get github.com/mieuxvoter/merit-profile-library-go
```

```golang
package main

import (
	"fmt"
	"github.com/mieuxvoter/merit-profile-library-go/merit"
)

func main() {
	proposals := []merit.Proposal{
		{
			Name:  "Alice the wonderful napping kangaroo 🦘 of the Æther",
			Tally: []uint64{4, 0, 3, 7}, // 4 grades, 14 judgments in total
		},
		{
			Name:  "Dominique",
			Tally: []uint64{5, 6, 1, 2}, // same
		},
		{
			Name:  "Théo 🗳",
			Tally: []uint64{3, 3, 2, 6}, // same
		},
	}

	svg, err := merit.RenderLinearProfileSVG(proposals)

	if err != nil {
		panic(err)
	}

	fmt.Print(svg)
}
```

> [!WARNING]
> Make sure your tallies are:
> - **Consistent**: Their shape must be the same.
> - **Balanced**: Their total must be the same.

![Merit profiles of the above code example](./test.svg)


## Options

There are options you can pass to `RenderLinearProfileSVG()` to customize the output.

> We use the _functional options pattern_, because it rocks.

Here's an example:

```golang
svg, err := merit.RenderLinearProfileSVG(
	proposals,
	merit.WithWidth(1024),
	merit.WithHeight(2048),
	merit.WithPadding(32),
	merit.WithVerticalSpacing(32),
	merit.WithBgColor(color.NRGBA{R: 0, G: 0, B: 0, A:255}),
	merit.WithMedianLineColor(color.NRGBA{R: 0, G: 0, B: 255, A:255}),
	merit.WithTextColor(color.NRGBA{R: 255, G: 0, B: 255, A:255}),
	merit.WithOutlineColor(color.NRGBA{R: 0, G: 255, B: 255, A:200}),
	//merit.WithGradesPalette(…),
	//merit.WithPatterns(…),
)
```

## Development Goodies

> Unit-testing on SVG generation is clunky at best, and not really worth it.

Therefore, we used a custom flavor of `svgplay` for convenience.

    go run svgplay.go

Visit http://localhost:1999/test.go

