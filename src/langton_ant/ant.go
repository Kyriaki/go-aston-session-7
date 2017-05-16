package main
 
import (
    "fmt"
    "image"
    "image/draw"
    "image/color"
    "image/png"
    "os"
)
 
const (
    up = iota
    rt
    dn
    lt
)
 
func main() {
    var posX, posY, sizeX, sizeY int
    if (len(os.Args) != 4) {
    // Get command line here
    //  return comment for sample code only
    }
//   sizeX=10 decomment for sample code only
//   sizeY=10 decomment for sample code only
    bounds := image.Rect(0, 0, sizeX, sizeY)
    im := image.NewGray(bounds)
    draw.Draw(im, bounds, image.NewUniform(color.Gray{255}), image.ZP, draw.Src)
    pos := image.Point{posX, posY}
    im.SetGray(pos.X, pos.Y, color.Gray{255})
    // implement algorithm there.
    f, err := os.Create("ant.png")
    if err != nil {
        fmt.Println(err)
        return
    }
    if err = png.Encode(f, im); err != nil {
        fmt.Println(err)
    }
    if err = f.Close(); err != nil {
        fmt.Println(err)
    }
}