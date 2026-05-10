# Merit Profile Generation for Golang

Generate merit profiles (in SVG), for use for example in [Majority Judgment] polls.

[Majority Judgment]: https://mieuxvoter.fr

## Usage

```shell
go get github.com/mieuxvoter/merit-profile-library-go
```

```go
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
            Name:  "Alice the wonderful napping kangaroo 🦘 of the Æther whose name is clipped",
        },
    }
    
    svg, err := merit.RenderSvg(proposals)

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


## Test

```shell
go test ./merit
```
