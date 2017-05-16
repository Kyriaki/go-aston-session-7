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
    var posX, posY, sizeX, sizeY, steps int
    if (len(os.Args) != 4) {
	// Get command line here
	// return comment for sample code only
    }

    cWhite:=color.Gray{255}
    cBlack:=color.Gray{0}    
    steps=2147483647 // 10 - 100 - 1000 - 10000 - 100000

    posX=100
    posY=100
    sizeX=200
    sizeY=200

    bounds := image.Rect(0, 0, sizeX, sizeY)
    im := image.NewGray(bounds)
    draw.Draw(im, bounds, image.NewUniform(cWhite), image.ZP, draw.Src)
    pos := image.Point{posX, posY}
    im.SetGray(pos.X, pos.Y, cWhite)

	// implement algorithm there.
    direction := up
    for i := 0; i < steps; i++{   
        //fmt.Println("step ", i)
        switch im.At(pos.X, pos.Y).(color.Gray).Y{
        case cWhite.Y:
            im.SetGray(pos.X, pos.Y, cBlack)
            direction++
        case cBlack.Y:
            im.SetGray(pos.X, pos.Y, cWhite)
            direction--

        }

        //fmt.Println(direction)

        switch direction%2{
        case 0:
            if direction%4 == 0 {
                pos.Y -= 1
            } else {
                pos.Y += 1
            }
        case 1:
            if (direction%4)%3 == 0{
                pos.X -= 1
            } else {
                pos.X += 1
            }

        }

        //fmt.Println(pos)
    }


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