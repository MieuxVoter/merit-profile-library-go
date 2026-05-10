# Merit Profile Generation for Golang

Generate merit profiles (in SVG), for use for example in [Majority Judgment] polls.

[Majority Judgment]: https://mieuxvoter.fr

## Usage

```shell
go get github.com/mieuxvoter/merit-profile-library-go
```

```go
package main

import "github.com/mieuxvoter/merit-profile-library-go/merit"

proposals := []merit.Proposal{
	// TODO
}

svg, err := merit.RenderSvg(proposals)


```


## Test

```shell
go test ./merit
```
