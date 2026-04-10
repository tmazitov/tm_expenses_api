package category

import (
	"fmt"
	"regexp"
	"strconv"
)

type Color uint32

var hexColorRegexp = regexp.MustCompile(`^#([0-9A-Fa-f]{3}|[0-9A-Fa-f]{6})$`)

func NewColor(hex string) (Color, error) {
	if !hexColorRegexp.MatchString(hex) {
		return 0, ErrColorInvalidFormat
	}

	raw := hex[1:]

	// #FFF → FFFFFF
	if len(raw) == 3 {
		raw = string([]byte{raw[0], raw[0], raw[1], raw[1], raw[2], raw[2]})
	}

	val, err := strconv.ParseUint(raw, 16, 32)
	if err != nil {
		return 0, ErrColorInvalidFormat
	}

	return Color(val), nil
}

func RestoreColor(hex uint32) (Color, error) {

	if hex > 0xFFFFFF {
		return 0, ErrColorInvalidFormat
	}

	return Color(hex), nil
}

func (c Color) Hex() string {
	return fmt.Sprintf("#%06X", uint32(c))
}

func (c Color) Uint32() uint32 {
	return uint32(c)
}
