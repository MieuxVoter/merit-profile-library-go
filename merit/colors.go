package merit

import (
	"fmt"
	"image/color"
)

// toHex outputs a SVG/CSS-compatible string like `#ff3399`.
// Be mindful that PRECISION MAY BE LOST because hex format has fewer bits.
// Additionally, color.Color has pre-multiplied alpha, so we have to compensate for that.
func toHex(c color.Color) string {
	nc, _ := color.NRGBAModel.Convert(c).(color.NRGBA)
	return fmt.Sprintf(`#%02x%02x%02x`, nc.R, nc.G, nc.B)
}

// toOpacity outputs a SVG/CSS-compatible opacity like `0.618`.
func toOpacity(c color.Color) string {
	_, _, _, a := c.RGBA()
	return fmt.Sprintf(`%.3f`, float32(a)/0xffff)
}

func toFillAttrs(c color.Color) string {
	return fmt.Sprintf(`fill="%s" fill-opacity="%s"`, toHex(c), toOpacity(c))
}

func toStrokeAttrs(c color.Color) string {
	return fmt.Sprintf(`stroke="%s" stroke-opacity="%s"`, toHex(c), toOpacity(c))
}
