module github.com/mieuxvoter/merit-profile-library-go

// This can perhaps be decreased ; feel free to, if you can test it.
go 1.24.0

require (
	github.com/ajstarks/gensvg v0.0.0-20210923152200-4042c242e95e
	github.com/mieuxvoter/majority-judgment-library-go v0.3.3
)

require github.com/lucasb-eyer/go-colorful v1.2.0 // indirect

// gensvg above is the same as svgo, but with float64 instead of int.
//require github.com/ajstarks/svgo v0.0.0-20211024235047-1546f124cd8b
