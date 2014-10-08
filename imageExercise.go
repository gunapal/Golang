package main

import (
    "code.google.com/p/go-tour/pic"
    "image"
    "image/color"
)

type Image struct{
    width int
    height int
}

func (i Image) ColorModel() color.Model {
    return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
    return image.Rect(0,0,i.width,i.height)
}

func (i Image) At(x, y int) color.Color {
    return color.RGBA{25,45,255,255}
}

func main() {
    m := Image{50,100}
    pic.ShowImage(m)
}