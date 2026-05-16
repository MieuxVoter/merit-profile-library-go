package merit

import (
	"fmt"
	"github.com/ajstarks/gensvg"
	"image/color"
	"math"
)

// In here we define some SVG background patterns, whose purpose is to help with accessibility.
// We use prefixed, ridiculously long ids in order to avoid collisions with eventual HTML ids.

// A PatternDefinition should call methods on the [*gensvg.SVG] canvas that are allowed in between
// [*gensvg.SVG.Pattern] and [*gensvg.SVG.PatternEnd], as well as these methods themselves.
// It may instead skip that and directly write to [*gensvg.SVG.Writer], so long as it is a pattern definition.
// It must return the id attribute it gave to the pattern, without the leading #.
// If you make your own, keep in mind that color.Color outputs components with pre-multiplied alpha.
type PatternDefinition func(canvas *gensvg.SVG, colour color.Color) string

// CreateDefaultPatterns is a factory for slices of PatternDefinition.
// It can support up to 20 different patterns (and then it'll loop).
func CreateDefaultPatterns(amount int) []PatternDefinition {
	switch amount {
	case 0:
		return []PatternDefinition{}
	case 1:
		return []PatternDefinition{
			NothingPatternDefinition,
		}
	default:
		defs := make([]PatternDefinition, amount)
		defs[0] = NothingPatternDefinition
		for i := 1; i < amount; i++ {
			defs[i] = HexagonPatternDefinitionCurried(
				fmt.Sprintf("merit_pattern_hexagons_%d", i),
				float64((i*2)%21), // beyond 20 all strokes are out of bounds
			)
		}
		return defs
	}
}

func NothingPatternDefinition(canvas *gensvg.SVG, _ color.Color) string {
	id := "merit_pattern_nothing"
	canvas.Pattern(
		id,
		0.0, 0.0,
		1000.0, 1000.0,
		"user",
	)
	canvas.PatternEnd()
	return id
}

func VerticalLinesPatternDefinition(canvas *gensvg.SVG, colour color.Color) string {
	id := "merit_pattern_vertical_lines"
	definePatternWithStrokedPath(
		canvas, colour, id,
		trimPathWhitespaces(`
        M 50,0
        L 50,100
        `),
		`viewBox="0 0 100 100"`,
	)
	return id
}

func AscendingLinesPatternDefinition(canvas *gensvg.SVG, colour color.Color) string {
	id := "merit_pattern_descending_lines"
	definePatternWithStrokedPath(
		canvas, colour, id,
		trimPathWhitespaces(`
        M 0,100
        L 100,0
        `),
		`viewBox="0 0 100 100"`,
	)
	return id
}

func DescendingLinesPatternDefinition(canvas *gensvg.SVG, colour color.Color) string {
	id := "merit_pattern_descending_lines"
	definePatternWithStrokedPath(
		canvas, colour, id,
		trimPathWhitespaces(`
        M 0,0
        L 100,100
        `),
		`viewBox="0 0 100 100"`,
	)
	return id
}

// PerfectAscendingLinesPatternDefinition is perfect but quite expensive, so we do not use it — but you might.
func PerfectAscendingLinesPatternDefinition(canvas *gensvg.SVG, colour color.Color) string {
	id := "merit_pattern_perfect_ascending_lines"
	definePatternWithFilledPath(
		canvas, colour, id,
		trimPathWhitespaces(`
        M 0,100
        L 0,95
        L 95,0
        L 100,0
        L 100,5
        L 5,100
        L 0,100
        Z
        M 100,100
        L 95,100
        L 100,95
        L 100,100
        Z
        M 0,0
        L 0,5
        L 5,0
        L 0,0
        Z
        `),
		`viewBox="0 0 100 100"`,
	)
	return id
}

func ZigZagPatternDefinition(canvas *gensvg.SVG, colour color.Color) string {
	id := "merit_pattern_zigzag"
	definePatternWithStrokedPath(
		canvas, colour, id,
		trimPathWhitespaces(`
        M 0,0
        L 62,38
        L 38,62
        L 100,100
        `),
		`viewBox="0 0 100 100"`,
	)
	return id
}

func HeartsPatternDefinition(canvas *gensvg.SVG, colour color.Color) string {
	id := "merit_pattern_hearts"
	definePatternWithFilledPath(
		canvas, colour, id,
		trimPathWhitespaces(`
        M 10,30
        A 20,20 0,0,1 50,30
        A 20,20 0,0,1 90,30
        Q 90,60 50,90
        Q 10,60 10,30
        Z
        `),
		`viewBox="0 0 100 100"`,
	)
	return id
}

func HexagonPatternDefinitionCurried(
	id string,
	radius float64,
) PatternDefinition {
	return func(canvas *gensvg.SVG, colour color.Color) string {
		h := 16.0
		w := h * 0.5 * math.Sqrt(3) // yields interesting emergent patterns
		r := radius
		canvas.Pattern(
			id,
			0.0, 0.0,
			w, h,
			"user",
		)
		canvas.Path(
			trimPathWhitespaces(hexPath(w*0.5, h*0.5, r)+
				hexPath(w, 0.0, r)+
				hexPath(0.0, 0.0, r)+
				hexPath(w, h, r)+
				hexPath(0.0, h, r),
			),
			toStrokeAttrs(colour),
			`stroke-width="0.8"`,
			`fill="none"`,
		)
		canvas.PatternEnd()

		return id
	}
}

func hexPath(x, y, r float64) string {
	hr := r * 0.5
	hw := hr * math.Sqrt(3)
	return fmt.Sprintf(" M %.2f %.2f", x, y-r) +
		fmt.Sprintf(" L %.2f %.2f", x-hw, y-hr) +
		fmt.Sprintf(" L %.2f %.2f", x-hw, y+hr) +
		fmt.Sprintf(" L %.2f %.2f", x, y+r) +
		fmt.Sprintf(" L %.2f %.2f", x+hw, y+hr) +
		fmt.Sprintf(" L %.2f %.2f", x+hw, y-hr) +
		" Z"
}

func definePatternWithFilledPath(
	canvas *gensvg.SVG,
	colour color.Color,
	id string,
	path string,
	attrs ...string,
) {
	canvas.Pattern(
		id,
		0.0, 0.0,
		16.0, 16.0,
		"user",
		attrs...,
	)
	canvas.Path(
		path,
		toFillAttrs(colour),
	)
	canvas.PatternEnd()
}

func definePatternWithStrokedPath(
	canvas *gensvg.SVG,
	colour color.Color,
	id string,
	path string,
	attrs ...string,
) {
	canvas.Pattern(
		id,
		0.0, 0.0,
		16.0, 16.0,
		"user",
		attrs...,
	)
	canvas.Path(
		path,
		toStrokeAttrs(colour),
		`stroke-width="5"`,
		`fill="none"`,
	)
	canvas.PatternEnd()
}
