package auth

import (
	"bytes"
	"encoding/base64"
	"errors"
	"image"
	"image/color"
	"image/png"
	"strconv"

	"github.com/Yeuoly/coshub/internal/utils/math"
	"github.com/Yeuoly/coshub/internal/utils/typ"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"

	_ "embed"
)

var font *truetype.Font
var colors [10]color.RGBA

//go:embed Consolas.ttf
var fontBytes []byte

func init() {
	font, _ = freetype.ParseFont(fontBytes)
	colors[0] = color.RGBA{255, 0, 0, 255}
	colors[1] = color.RGBA{25, 25, 112, 255}
	colors[2] = color.RGBA{220, 20, 60, 255}
	colors[3] = color.RGBA{25, 25, 112, 255}
	colors[4] = color.RGBA{255, 165, 0, 255}
	colors[5] = color.RGBA{128, 0, 0, 255}
	colors[6] = color.RGBA{178, 34, 34, 255}
	colors[7] = color.RGBA{255, 140, 0, 255}
	colors[8] = color.RGBA{0, 100, 0, 255}
	colors[9] = color.RGBA{100, 149, 237, 255}
}

func generateB64ImageFromText(exp string, char_size int) (string, error) {
	padding := 10
	word_size := char_size + char_size/3
	w := padding*2 + word_size*len(exp) + 30
	h := padding*2 + word_size + 30

	img := image.NewNRGBA(image.Rect(0, 0, w, h))

	// fill background
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			img.Set(x, y, color.RGBA{255, 255, 255, 255})
		}
	}

	// draw lines
	ctx := freetype.NewContext()
	ctx.SetFont(font)
	ctx.SetDPI(72)
	ctx.SetFontSize(float64(char_size))
	ctx.SetClip(img.Bounds())
	ctx.SetDst(img)
	ctx.SetSrc(&image.Uniform{color.RGBA{255, 0, 0, 255}})

	lines := math.Random(2, 4)
	for i := 0; i < lines; i++ {

		color := colors[math.Random(0, 9)]

		mode := math.Random(0, 1)
		switch mode {
		case 0:
			// in random height
			h1, h2 := math.Random(0, h-1), math.Random(0, h-1)
			drawLine(0, h1, w-1, h2, func(x, y int) {
				img.Set(x, y, color)
			})
		case 1:
			// in random width
			w1, w2 := math.Random(0, w-1), math.Random(0, w-1)
			drawLine(w1, 0, w2, h-1, func(x, y int) {
				img.Set(x, y, color)
			})
		}
	}

	//一个一个画字符，间距不能一样，要骚一点，先随机出一个高度区间
	rect_left := math.Random(padding, w-word_size*len(exp)-padding)
	rect_top := math.Random(padding, h-word_size-padding)

	for i := range exp {
		pos := freetype.Pt(rect_left+i*word_size+math.Random(0, 10), rect_top+math.Random(0, 10)+word_size)
		ctx.SetSrc(&image.Uniform{colors[math.Random(0, 9)]})
		_, err := ctx.DrawString(string(exp[i]), pos)
		if err != nil {
			return "", err
		}
	}

	buff := bytes.NewBuffer(nil)
	err := png.Encode(buff, img)
	if err != nil {
		return "", err
	}

	//对比了一下向下取整再做if的效率和使用float向上取整的效率，好像用if效率更高，任何数学库的底层也都还是if
	buf_size := buff.Len()
	if buf_size%3 != 0 {
		padding = 4
	} else {
		padding = 0
	}
	b64_buf_size := (buf_size/3.0)*4 + padding
	b64_buf := make([]byte, b64_buf_size)
	base64.StdEncoding.Encode(b64_buf, buff.Bytes())

	return "data:image/png;base64," + typ.BytesToString(b64_buf), nil
}

func drawLine(x1 int, y1 int, x2 int, y2 int, callback func(x int, y int)) {
	dx := math.Abs(x1 - x2)
	dy := math.Abs(y1 - y2)
	sx, sy := 1, 1
	if x1 >= x2 {
		sx = -1
	}
	if y1 >= y2 {
		sy = -1
	}

	angle := dx - dy

	for {
		callback(x1, y1)
		if x1 == x2 && y1 == y2 {
			return
		}
		a2 := angle * 2
		if a2 > -dy {
			angle -= dy
			x1 += sx
		}
		if a2 < dx {
			angle += dx
			y1 += sy
		}
	}
}

// generateB64ImageVercode generates a vercode image and returns the base64 string and the result
func generateB64ImageVercode() (string, string, error) {
	// exp presents the expression of the vercode
	var exp string = ""
	//生成两个数用于计算
	x, y, ans := math.Random(1, 30), math.Random(1, 30), ""
	switch math.Random(0, 3) {
	// mul
	case 0:
		ary := []string{strconv.Itoa(x), strconv.Itoa(y), strconv.Itoa(x * y)}
		index := math.Random(0, 2)
		ans = ary[index]
		ary[index] = "?"
		exp = ary[0] + "*" + ary[1] + "=" + ary[2]
	// div
	case 1:
		ary := []string{strconv.Itoa(x * y), strconv.Itoa(y), strconv.Itoa(x)}
		index := math.Random(0, 2)
		ans = ary[index]
		ary[index] = "?"
		exp = ary[0] + "/" + ary[1] + "=" + ary[2]
	// add
	case 2:
		ary := []string{strconv.Itoa(x), strconv.Itoa(y), strconv.Itoa(x + y)}
		index := math.Random(0, 2)
		ans = ary[index]
		ary[index] = "?"
		exp = ary[0] + "+" + ary[1] + "=" + ary[2]
	// sub
	case 3:
		ary := []string{strconv.Itoa(x + y), strconv.Itoa(y), strconv.Itoa(x)}
		index := math.Random(0, 2)
		ans = ary[index]
		ary[index] = "?"
		exp = ary[0] + "-" + ary[1] + "=" + ary[2]
	}

	// generate image
	b64, err := generateB64ImageFromText(exp, 30)
	if err != nil {
		return "", "", err
	}

	return b64, ans, nil
}

// SendVercodeEmail sends verification code to user's email and return the base64 encoded image and the encrypted vercode
func GetVercodeImage(method string) (string, string, error) {
	vercode_image, result, err := generateB64ImageVercode()
	if err != nil {
		return "", "", errors.New("failed to generate vercode image")
	}

	vercode := NewImageVercode(result, 1, method)
	token, err := vercode.ToToken()
	if err != nil {
		return "", "", errors.New("failed to generate vercode token")
	}

	return vercode_image, token, nil
}

// CheckImageVercode checks if the email verification code is correct
func CheckImageVercode(token string, result string, method string) error {
	vercode, err := ImageVercodeFromToken(token)
	if err != nil {
		return err
	}

	if !vercode.Check(method, result) {
		return errors.New("invalid vercode")
	}

	return nil
}
