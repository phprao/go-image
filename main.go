package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/color/palette"
	"image/draw"
	"image/gif"
	"image/jpeg"
	_ "image/jpeg"
	_ "image/png"
	"math/rand"
	"os"
	//draw2 "golang.org/x/image/draw"

	"github.com/fogleman/gg"
	"github.com/ryandao/go-mosaic"
	draw2 "golang.org/x/image/draw"
)

func main() {
	//line1()
	line2()
}

func mask() {
	draw2.Draw()
	draw2.DrawMask()
}

// ??????
func line2() {
	defer func() {
		if err := recover(); err != nil {
			println(err)
		}
	}()
	s := image.NewUniform(color.White)

	d := image.NewRGBA(image.Rect(0, 0, 500, 500))

	draw.Draw(d, d.Rect, s, image.Point{}, draw.Src)

	var y int
	var x0, y0, x1, y1, k, dd float64
	x1 = 400
	y1 = 250
	k = (y1 - y0) / (x1 - x0) // 0.375

	for x := 0; x < int(x1); x++ {
		d.Set(x, y, color.RGBA{R: 255, A: 255})
		dd = dd + k
		if dd > 0.5 {
			y = y + 1
			dd = dd - 1
		}
	}

	f, _ := os.Create("./line2.jpg")
	defer f.Close()

	jpeg.Encode(f, d, nil)
}

func line1() {
	s := image.NewUniform(color.White)

	d := image.NewRGBA(image.Rect(0, 0, 500, 500))

	draw.Draw(d, d.Rect, s, image.Point{}, draw.Src)

	var x0, y0, x1, y1, k float64
	x1 = 400
	y1 = 250
	k = (y1 - y0) / (x1 - x0)
	b := y1 - k*x1
	for i := 0; i < int(x1); i++ {
		yt := k*float64(i) + b
		if yt-float64(int(yt)) > 0.5 {
			d.Set(i, int(yt)+1, color.RGBA{R: 255, A: 255})
		} else {
			d.Set(i, int(yt), color.RGBA{R: 255, A: 255})
		}
	}

	f, _ := os.Create("./line1.jpg")
	defer f.Close()

	jpeg.Encode(f, d, nil)
}

// mosaic4
// 随机的颜色作为马赛克的颜色
func image22() {
	scale := 7

	m1, _ := os.Open("./image1.jpg")
	defer m1.Close()
	img1, _, _ := image.Decode(m1)
	img11 := image.NewRGBA(img1.Bounds())

	for i := 0; i < img1.Bounds().Max.X; i++ {
		for j := 0; j < img1.Bounds().Max.Y; j++ {
			if i%scale == 0 && j%scale == 0 {
				d := rand.Intn(scale)
				po := img1.At(d+i, d+j)
				for m := 0; m < scale; m++ {
					for n := 0; n < scale; n++ {
						img11.Set(i+m, j+n, po)
					}
				}
			}

		}
	}

	f1, _ := os.Create("./image1-22.jpg")
	defer f1.Close()
	jpeg.Encode(f1, img11, nil)
}

// mosaic3
// 取第一个点的颜色作为马赛克的颜色
func image21() {
	scale := 7

	m1, _ := os.Open("./image1.jpg")
	defer m1.Close()
	img1, _, _ := image.Decode(m1)
	img11 := image.NewRGBA(img1.Bounds())

	for i := 0; i < img1.Bounds().Max.X; i++ {
		for j := 0; j < img1.Bounds().Max.Y; j++ {
			if i%scale == 0 && j%scale == 0 {
				po := img1.At(i, j)
				for m := 0; m < scale; m++ {
					for n := 0; n < scale; n++ {
						img11.Set(i+m, j+n, po)
					}
				}
			}

		}
	}

	f1, _ := os.Create("./image1-21.jpg")
	defer f1.Close()
	jpeg.Encode(f1, img11, nil)
}

// mosaic2
// 取中心点的颜色作为马赛克的颜色
func image20() {
	scale := 7

	m1, _ := os.Open("./image1.jpg")
	defer m1.Close()
	img1, _, _ := image.Decode(m1)
	img11 := image.NewRGBA(img1.Bounds())

	x := img1.Bounds().Max.X / scale
	y := img1.Bounds().Max.Y / scale

	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			po := img1.At(i*scale+2, j*scale+2)
			for m := 0; m < scale; m++ {
				for n := 0; n < scale; n++ {
					img11.Set(i*scale+m, j*scale+n, po)
				}
			}
		}
	}

	f1, _ := os.Create("./image1-20.jpg")
	defer f1.Close()
	jpeg.Encode(f1, img11, nil)
}

// mosaic1
// 取区域平均颜色作为马赛克的颜色，有问题
func image17() {
	scale := 7
	mosa := uint64(scale * scale)

	m1, _ := os.Open("./image1.jpg")
	defer m1.Close()
	img1, _, _ := image.Decode(m1)
	img11 := image.NewRGBA(img1.Bounds())

	var sumr, sumg, sumb uint64

	for i := 0; i < img1.Bounds().Max.X; i++ {
		for j := 0; j < img1.Bounds().Max.Y; j++ {
			if i%scale == 0 && j%scale == 0 {
				sumr = 0
				sumg = 0
				sumb = 0
				for m := 0; m < scale; m++ {
					for n := 0; n < scale; n++ {
						r, g, b, _ := img1.At(i+m, j+n).RGBA()
						//fmt.Printf("x=%d,y=%d,(%d,%d,%d)\n", i*scale+m, j*scale+n, r, g, b)
						sumr += uint64(r)
						sumg += uint64(g)
						sumb += uint64(b)
					}
				}

				R := sumr / mosa
				G := sumg / mosa
				B := sumb / mosa

				//fmt.Printf("(%d,%d,%d)\n", sumr, sumg, sumb)
				//fmt.Printf("(%d,%d,%d)\n", sumr/mosa, sumg/mosa, sumb/mosa)
				//fmt.Printf("(%d,%d,%d)\n", R, G, B)

				for m := 0; m < scale; m++ {
					for n := 0; n < scale; n++ {
						img11.Set(i+m, j+n, color.RGBA{R: uint8(R), G: uint8(G), B: uint8(B), A: 255})
					}
				}
			}
		}
	}

	f1, _ := os.Create("./image1-17.jpg")
	defer f1.Close()
	jpeg.Encode(f1, img11, nil)
}

func image19() {
	scale := 4

	m1, _ := os.Open("./image1.jpg")
	defer m1.Close()
	img1, _, _ := image.Decode(m1)
	img11 := image.NewRGBA(img1.Bounds())

	mosaic.Mosaic(img11, []image.Image{img1}, scale)

	f1, _ := os.Create("./image1-1.jpg")
	defer f1.Close()
	jpeg.Encode(f1, img11, nil)
}

func image18() {
	dc := gg.NewContext(200, 200)
	dc.DrawCircle(100, 100, 50)
	dc.SetRGB255(255, 0, 0)
	dc.Fill()

	img := image.NewRGBA(image.Rect(0, 0, 200, 200))
	draw.Draw(img, img.Bounds(), dc.Image(), dc.Image().Bounds().Min, draw.Src)

	f1, _ := os.Create("./image18.jpg")
	defer f1.Close()
	jpeg.Encode(f1, img, nil)
}

func image182() {
	dc := gg.NewContext(200, 200)
	dc.DrawCircle(100, 100, 50)
	dc.SetRGB255(255, 0, 0)
	dc.Fill()
	dc.SavePNG("./image182.png")
}

func image16() {
	f, err := os.Open("./333.jpg")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	src, _, _ := image.Decode(f)

	dr := image.Rect(0, 0, 100, 100)
	dst := image.NewRGBA(dr)
	draw2.BiLinear.Scale(dst, dr, src, src.Bounds(), draw2.Src, nil)

	f1, _ := os.Create("./3333.jpg")
	defer f1.Close()

	jpeg.Encode(f1, dst, nil)
}

func image15() {
	p1 := image.NewPaletted(image.Rect(0, 0, 100, 100), palette.Plan9)
	p2 := image.NewPaletted(image.Rect(0, 0, 100, 100), palette.Plan9)
	p3 := image.NewPaletted(image.Rect(0, 0, 100, 100), palette.Plan9)

	m1, _ := os.Open("./1111.jpg")
	defer m1.Close()
	img1, _, _ := image.Decode(m1)

	m2, _ := os.Open("./2222.jpg")
	defer m2.Close()
	img2, _, _ := image.Decode(m2)

	m3, _ := os.Open("./3333.jpg")
	defer m3.Close()
	img3, _, _ := image.Decode(m3)

	draw.Draw(p1, p1.Bounds(), img1, img1.Bounds().Min, draw.Src)
	draw.Draw(p2, p2.Bounds(), img2, img2.Bounds().Min, draw.Src)
	draw.Draw(p3, p3.Bounds(), img3, img3.Bounds().Min, draw.Src)

	g := &gif.GIF{
		Image:     []*image.Paletted{p1, p2, p3},
		Delay:     []int{100, 100, 100},
		LoopCount: 0,
	}
	f1, _ := os.Create("./image15.gif")
	defer f1.Close()

	gif.EncodeAll(f1, g)
}

func image14() {
	p1 := image.NewPaletted(image.Rect(0, 0, 100, 100), palette.Plan9)
	p2 := image.NewPaletted(image.Rect(0, 0, 100, 100), palette.Plan9)

	for y := 0; y < 100; y++ {
		p1.Set(50, y, color.RGBA{R: 0, G: 0, B: 255, A: 255})
	}

	for y := 0; y < 100; y++ {
		p2.Set(y, 50, color.RGBA{R: 255, G: 0, B: 0, A: 255})
	}

	g := &gif.GIF{
		Image:     []*image.Paletted{p1, p2},
		Delay:     []int{10, 100},
		LoopCount: 0,
	}
	f1, _ := os.Create("./image14.gif")
	defer f1.Close()

	gif.EncodeAll(f1, g)
}

func image13() {
	f, err := os.Open("./image1.jpg")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	src, _, _ := image.Decode(f)

	scale := 2
	dstWidth := src.Bounds().Max.X * scale
	dstHeight := src.Bounds().Max.Y * scale

	dr := image.Rect(0, 0, dstWidth, dstHeight)
	dst := image.NewRGBA(dr)
	draw2.BiLinear.Scale(dst, dr, src, src.Bounds(), draw2.Src, nil)

	f1, _ := os.Create("./image13.jpg")
	defer f1.Close()

	b := bufio.NewWriter(f1)
	jpeg.Encode(b, dst, nil)
	b.Flush()
}

func image12() {
	f, err := os.Open("./image1.jpg")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scale := 2

	src, _, _ := image.Decode(f)
	width := src.Bounds().Max.X
	height := src.Bounds().Max.Y

	dstWidth := width
	dstHeight := height * scale

	dst := image.NewRGBA(image.Rect(0, 0, dstWidth, dstHeight))

	for i := 0; i < dstWidth; i++ {
		for j := 0; j < dstHeight; j++ {
			dst.Set(i, j, src.At(i, j/scale))
		}
	}

	f1, _ := os.Create("./image12.jpg")
	defer f1.Close()

	b := bufio.NewWriter(f1)
	jpeg.Encode(b, dst, nil)
	b.Flush()
}

func image11() {
	f, err := os.Open("./image1.jpg")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scale := 2 // 等比例缩小一半

	src, _, _ := image.Decode(f)
	width := src.Bounds().Max.X
	height := src.Bounds().Max.Y

	dstWidth := width / scale
	dstHeight := height / scale

	dst := image.NewRGBA(image.Rect(0, 0, dstWidth, dstHeight))

	for i := 0; i < dstWidth; i++ {
		for j := 0; j < dstHeight; j++ {
			dst.Set(i, j, src.At(i*scale, j*scale))
		}
	}

	f1, _ := os.Create("./image11.jpg")
	defer f1.Close()

	b := bufio.NewWriter(f1)
	jpeg.Encode(b, dst, nil)
	b.Flush()
}

func image10() {
	f, err := os.Open("./image1.jpg")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	src, _, _ := image.Decode(f)

	// jpeg 使用的 YCbCr 颜色
	dst := src.(*image.YCbCr).SubImage(image.Rect(0, 0, 500, 500))

	f1, _ := os.Create("./image10.jpg")
	defer f1.Close()

	b := bufio.NewWriter(f1)
	jpeg.Encode(b, dst, nil)
	b.Flush()
}

func image9() {
	f, err := os.Open("./image1.jpg")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	src, _, _ := image.Decode(f)

	dst := image.NewGray(src.Bounds())

	draw.Draw(dst, src.Bounds(), src, src.Bounds().Min, draw.Src)

	f1, _ := os.Create("./image9.jpg")
	defer f1.Close()

	b := bufio.NewWriter(f1)
	jpeg.Encode(b, dst, nil)
	b.Flush()
}

func image8() {
	// 蓝色 500*500
	img1 := image.NewRGBA(image.Rect(0, 0, 500, 500))
	for i := 0; i < img1.Bounds().Max.X; i++ {
		for j := 0; j < img1.Bounds().Max.Y; j++ {
			img1.Set(i, j, color.RGBA{B: 255, A: 255})
		}
	}

	// 红色 200*200
	img2 := image.NewRGBA(image.Rect(0, 0, 200, 200))
	for i := 0; i < img2.Bounds().Max.X; i++ {
		for j := 0; j < img2.Bounds().Max.Y; j++ {
			img2.Set(i, j, color.RGBA{R: 255, A: 255})
		}
	}

	draw.Draw(img1, img2.Bounds(), img2, img2.Bounds().Min, draw.Src)

	f, _ := os.Create("./image8.jpg")
	defer f.Close()

	b := bufio.NewWriter(f)
	jpeg.Encode(b, img1, nil)
	b.Flush()
}

func image7() {
	dx := 500
	dy := 500
	r := image.Rect(0, 0, dx, dy)
	img := image.NewRGBA(r)
	imgBack := image.NewUniform(image.Black)
	draw.Draw(img, r, imgBack, image.Point{}, draw.Src)

	f, _ := os.Create("./image7.jpg")
	defer f.Close()

	b := bufio.NewWriter(f)
	jpeg.Encode(b, img, nil)
	b.Flush()
}

func image6() {
	dx := 500
	dy := 500
	r := image.Rect(0, 0, dx, dy)
	img := image.NewRGBA(r)
	for i := 0; i < dx; i++ {
		for j := 0; j < dy; j++ {
			img.Set(i, j, color.Black)
		}
	}

	f, _ := os.Create("./image6.jpg")
	defer f.Close()

	b := bufio.NewWriter(f)
	jpeg.Encode(b, img, nil)
	b.Flush()
}

func image5() {
	dx := 500
	dy := 500
	img := image.NewCMYK(image.Rect(0, 0, dx, dy))
	for i := 0; i < dx; i++ {
		for j := 0; j < dy; j++ {
			img.Set(i, j, color.CMYK{C: uint8(i % 256), M: uint8(i % 256), Y: uint8(i % 256), K: uint8(i % 256)})
		}
	}

	f, _ := os.Create("./image5.jpg")
	defer f.Close()

	b := bufio.NewWriter(f)
	jpeg.Encode(b, img, nil)
	b.Flush()
}

func image4() {
	dx := 500
	dy := 500
	img := image.NewGray(image.Rect(0, 0, dx, dy))
	for i := 0; i < dx; i++ {
		for j := 0; j < dy; j++ {
			img.Set(i, j, color.Gray{Y: uint8(i % 256)})
		}
	}

	f, _ := os.Create("./image4.jpg")
	defer f.Close()

	b := bufio.NewWriter(f)
	jpeg.Encode(b, img, nil)
	b.Flush()
}

func image3() {
	dx := 500
	dy := 500
	img := image.NewAlpha(image.Rect(0, 0, dx, dy))
	for i := 0; i < dx; i++ {
		for j := 0; j < dy; j++ {
			img.Set(i, j, color.Alpha{A: uint8(i % 256)})
		}
	}

	f, _ := os.Create("./image3.jpg")
	defer f.Close()

	b := bufio.NewWriter(f)
	jpeg.Encode(b, img, nil)
	b.Flush()
}

func image2() {
	img := image.NewRGBA(image.Rect(0, 0, 100, 100))
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			img.Set(i, j, color.RGBA{R: 255, G: 0, B: 0, A: 255})
		}
	}

	f, _ := os.Create("./image2.jpg")
	defer f.Close()

	b := bufio.NewWriter(f)
	jpeg.Encode(b, img, nil)
	b.Flush()
}

func imageMosaic(source string, length int) {
	f, err := os.Open("./image1.jpg")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	img, fmtName, err := image.Decode(f)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Name: %v, Bounds: %v, Color: %+v", fmtName, img.Bounds(), img.ColorModel())
}
