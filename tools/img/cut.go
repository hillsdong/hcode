package img

import (
	"fmt"
	"image"
	"image/png"
	"os"
)

// New 新建
func New(path string) *Img {
	i := &Img{}
	if err := i.load(path); err != nil {
		panic("load Img fail!")
	}
	size := i.data.Bounds().Size()
	i.width = size.X
	i.height = size.Y
	return i
}

type Img struct {
	width  int
	height int
	data   image.Image
}

func (i *Img) load(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	if i.data, _, err = image.Decode(file); err != nil {
		return err
	}
	return nil
}

// Save 保存为PNG
func (i *Img) Save(path string) error {
	Imgfile, err := os.Create(path)
	if err != nil {
		return err
	}
	defer Imgfile.Close()

	if err = png.Encode(Imgfile, i.data); err != nil {
		return err
	}
	return nil
}

// Cut 裁剪图片
func (i *Img) Cut(x, y, w, h int) *Img {
	//x,y可为负数，按照反方向计算
	if x < 0 {
		x = i.width + x
	}
	if y < 0 {
		y = i.height + y
	}
	//检验
	if x < 0 || y < 0 || w < 0 || h < 0 {
		return i
	}
	//防止超出原图
	xe, ye := x+w, y+h
	if xe > i.width {
		xe = i.width
	}
	if ye > i.height {
		ye = i.height
	}
	//新建图片
	ci := image.NewRGBA(image.Rect(0, 0, w, h))
	//裁剪
	for xi := x; xi < xe; xi++ {
		for yi := y; yi < ye; yi++ {
			ci.Set(xi-x, yi-y, i.data.At(xi, yi))
		}
	}
	i.width, i.height, i.data = w, h, ci
	return i
}

func (i *Img) String() string {
	return fmt.Sprintf("w: %d, h: %d", i.width, i.height)
}
