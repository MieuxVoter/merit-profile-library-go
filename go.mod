module github.com/mieuxvoter/merit-profile-library-go

// gensvg requires 1.16 (others are 1.12, afaik)
go 1.16

require (
	// Lots of things are wrong in this lib, but I'm not about to go shave that yak.
	github.com/ajstarks/gensvg v0.0.0-20210923152200-4042c242e95e
	// gensvg above is the same as svgo below, but with float64 instead of int.
	//require github.com/ajstarks/svgo v0.0.0-20211024235047-1546f124cd8b

	// We use only the color palette generation from this (we do not rank ourselves — yet)
	github.com/mieuxvoter/majority-judgment-library-go v0.3.3
)

require github.com/lucasb-eyer/go-colorful v1.4.0 // indirect
