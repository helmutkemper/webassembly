package utilsMath

import (
	"errors"
	"fmt"
	"image/color"
	"strconv"
	"strings"
)

func HexToColor(hex string) (converted color.Color, err error) {
	if !strings.HasPrefix(hex, "#") {
		err = errors.New("value must be in format #rgb, eg. #fff or #ffffff")
		return
	}

	if !(len(hex) == 4 || len(hex) == 7) {
		err = errors.New("value must be in format #rgb, eg. #fff or #ffffff")
		return
	}

	hex = strings.TrimPrefix(hex, "#")

	var r, g, b uint64

	if len(hex) == 3 {
		hex = fmt.Sprintf("%c%c%c%c%c%c", hex[0], hex[0], hex[1], hex[1], hex[2], hex[2])
	}

	r, err = strconv.ParseUint(hex[0:2], 16, 8)
	if err != nil {
		return nil, err
	}
	g, err = strconv.ParseUint(hex[2:4], 16, 8)
	if err != nil {
		return nil, err
	}
	b, err = strconv.ParseUint(hex[4:6], 16, 8)
	if err != nil {
		return nil, err
	}

	return color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: 255}, nil
}
