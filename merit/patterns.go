package merit

// In here we define some SVG background patterns,
// whose purpose is to help with accessibility.

import (
	"fmt"
	"github.com/ajstarks/gensvg"
	"image/color"
	"math"
	"regexp"
	"strings"
)

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
				fmt.Sprintf("hexagons_%d", i),
				float64((i*2)%21), // beyond 20 all strokes are out of bounds
			)
		}
		return defs
	}
}

//var availablePatternDefinitions = map[string]func(canvas *gensvg.SVG, colour color.Color){
//	"nothing":          func(canvas *gensvg.SVG, colour color.Color) {},
//	"hearts":           definePatternOfHearts,
//	"ascending_lines":  definePatternOfAscendingLines,
//	"descending_lines": definePatternOfDescendingLines,
//	"vertical_lines":   definePatternOfVerticalLines,
//	"zigzag":           definePatternOfZigZagLines,
//	"hexagons_1":       curryDefinePatternOfHexagons(1),
//	"hexagons_2":       curryDefinePatternOfHexagons(2),
//	"hexagons_3":       curryDefinePatternOfHexagons(3),
//	"hexagons_4":       curryDefinePatternOfHexagons(4),
//	"hexagons_5":       curryDefinePatternOfHexagons(5),
//	"hexagons_6":       curryDefinePatternOfHexagons(6),
//	"hexagons_7":       curryDefinePatternOfHexagons(7),
//	"hexagons_8":       curryDefinePatternOfHexagons(8),
//	"hexagons_9":       curryDefinePatternOfHexagons(9),
//	"hexagons_10":      curryDefinePatternOfHexagons(10),
//	"hexagons_11":      curryDefinePatternOfHexagons(11),
//	"hexagons_12":      curryDefinePatternOfHexagons(12),
//	"hexagons_13":      curryDefinePatternOfHexagons(13),
//	"hexagons_14":      curryDefinePatternOfHexagons(14),
//	"hexagons_15":      curryDefinePatternOfHexagons(15),
//	"hexagons_16":      curryDefinePatternOfHexagons(16),
//	"hexagons_17":      curryDefinePatternOfHexagons(17),
//	"hexagons_18":      curryDefinePatternOfHexagons(18),
//	"hexagons_19":      curryDefinePatternOfHexagons(19),
//	"hexagons_20":      curryDefinePatternOfHexagons(20),
//}

//func chooseDefaultPatternDefinitionForIndex(
//	index int,
//	totalAmount int,
//) func(canvas *gensvg.SVG, colour color.Color) {
//	index = index % totalAmount
//	//coll := make([]func(canvas *gensvg.SVG, colour color.Color), 0)
//	switch totalAmount {
//	case 0:
//		return availablePatternDefinitions["nothing"]
//	case 1:
//		return availablePatternDefinitions["nothing"]
//	case 2:
//		return availablePatternDefinitions[[]string{
//			"nothing",
//			"ascending_lines",
//		}[index]]
//	case 3:
//		return availablePatternDefinitions[[]string{
//			"nothing",
//			"hexagons_2",
//			"ascending_lines",
//		}[index]]
//	case 4:
//		return availablePatternDefinitions[[]string{
//			"nothing",
//			"descending_lines",
//			"hexagons_2",
//			"ascending_lines",
//		}[index]]
//	case 5:
//		return availablePatternDefinitions[[]string{
//			"nothing",
//			"hexagons_1",
//			"descending_lines",
//			"hexagons_5",
//			"ascending_lines",
//		}[index]]
//	case 6:
//		return availablePatternDefinitions[[]string{
//			"nothing",
//			"hexagons_1",
//			"descending_lines",
//			"hexagons_5",
//			"ascending_lines",
//			"hearts",
//		}[index]]
//	case 7:
//		return availablePatternDefinitions[[]string{
//			"nothing",
//			"hexagons_1",
//			"descending_lines",
//			"hexagons_5",
//			"ascending_lines",
//			"hexagons_11",
//			"hearts",
//		}[index]]
//	default:
//		return availablePatternDefinitions[fmt.Sprintf("hexagons_%d", (index+1)%21)]
//	}
//}

var spacesRegex = regexp.MustCompile("\\s+")
var lettersRegex = regexp.MustCompile("[a-zA-Z]\\s+")

func trimPathWhitespaces(path string) string {
	path = spacesRegex.ReplaceAllString(path, " ")
	path = lettersRegex.ReplaceAllStringFunc(path, func(s string) string { return s[:1] })
	path = strings.TrimSpace(path)
	return path
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

func definePatternOfHearts(canvas *gensvg.SVG, colour color.Color) {
	definePatternWithFilledPath(
		canvas, colour,
		"hearts",
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
}

func definePatternOfAscendingLines(canvas *gensvg.SVG, colour color.Color) {
	definePatternWithFilledPath(
		canvas, colour,
		"ascending_lines",
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
}

func NothingPatternDefinition(canvas *gensvg.SVG, colour color.Color) string {
	id := "nothing"
	canvas.Pattern(
		id,
		0.0, 0.0,
		1.0, 1.0,
		"user",
	)
	canvas.PatternEnd()
	return id
}

func AscendingLinesPatternDefinition(canvas *gensvg.SVG, colour color.Color) string {
	id := "descending_lines"
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
	id := "descending_lines"
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

func definePatternOfVerticalLines(canvas *gensvg.SVG, colour color.Color) {
	definePatternWithStrokedPath(
		canvas, colour,
		"vertical_lines",
		trimPathWhitespaces(`
        M 50,0
        L 50,100
        `),
		`viewBox="0 0 100 100"`,
	)
}

func definePatternOfZigZagLines(canvas *gensvg.SVG, colour color.Color) {
	definePatternWithStrokedPath(
		canvas, colour,
		"zigzag",
		trimPathWhitespaces(`
        M 0,0
        L 62,38
        L 38,62
        L 100,100
        `),
		`viewBox="0 0 100 100"`,
	)
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

func HexagonPatternDefinitionCurried(
	id string,
	radius float64,
) PatternDefinition {
	return func(canvas *gensvg.SVG, colour color.Color) string {
		h := 16.0
		w := h * 0.5 * math.Sqrt(3)
		r := radius
		canvas.Pattern(
			//fmt.Sprintf("hexagons_%d", r),
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
