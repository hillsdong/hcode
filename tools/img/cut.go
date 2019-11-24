package img

import (
	"fmt"
	"image"
	"image/png"
	"math"
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

// Img 图片处理
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
	xs, xe, ys, ye := x, x+w, y, y+h
	if xe > i.width {
		xe = i.width
	}
	if ye > i.height {
		ye = i.height
	}
	//新建图片
	ci := image.NewRGBA(image.Rect(0, 0, w, h))
	//裁剪
	for xi := xs; xi < xe; xi++ {
		for yi := ys; yi < ye; yi++ {
			ci.Set(xi-xs, yi-ys, i.data.At(xi, yi))
		}
	}
	i.width, i.height, i.data = w, h, ci
	return i
}

// Circle 裁剪，输入圆心xy坐标，半径r
func (i *Img) Circle(x, y, r int) *Img {
	//x,y可为负数，按照反方向计算
	if x < 0 {
		x = i.width + x
	}
	if y < 0 {
		y = i.height + y
	}
	//检验
	if x < 0 || y < 0 || r < 0 {
		return i
	}
	//防止超出原图
	xs, xe := x-r, x+r
	if xs < 0 {
		xs = 0
	}
	if xe > i.width {
		xe = i.width
	}
	//新建图片
	ci := image.NewRGBA(image.Rect(0, 0, 2*r, 2*r))
	//裁剪
	for xi := xs; xi < xe; xi++ {
		//根据x位置计算y点范围
		yF := float64(r*r - (xi-x)*(xi-x))
		yI := int(math.Sqrt(yF))
		ys, ye := y-yI, y+yI
		if ys < 0 {
			ys = 0
		}
		if ye > i.height {
			ye = i.height
		}

		for yi := ys; yi < ye; yi++ {
			ci.Set(xi-(x-r), yi-(y-r), i.data.At(xi, yi))
		}
	}
	i.width, i.height, i.data = 2*r, 2*r, ci
	return i
}

func (i *Img) String() string {
	return fmt.Sprintf("w: %d, h: %d", i.width, i.height)
}
