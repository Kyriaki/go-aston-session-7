package main
 
import (
	"fmt"
	"image"
	"image/draw"
	"image/color"
	"image/png"
	"os"
    "flag"
    "strings"
    "strconv"
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
        /***** Retrieve arguments from command line *****/
        // We flag the arguments as strings, then parse the output and split the strings to get the values
        sizeArg := flag.String("size", "300,300", "Size of the grid\nInput must be like '-size=sizeX,sizeY'")
        posArg := flag.String("pos", "150,150", "Starting position of the ant\nInput must be like '-pos=posX,posY'")

        flag.Parse()
        size := strings.Split(*sizeArg, ",")
        pos := strings.Split(*posArg, ",")

        //If the values are correctly given, we convert them to integers (error if it doesn't work) and give them to the corresponding variables
        if len(size) == 2 && len(pos) == 2 {
            sX, err1 := strconv.Atoi(size[0])
            sY, err2 := strconv.Atoi(size[1])
            pX, err3 := strconv.Atoi(pos[0])
            pY, err4 := strconv.Atoi(pos[1])
            if err1 != nil || err2 != nil || err3 != nil || err4 != nil  { 
                fmt.Println("Error : invalid argument type\nArguments must be integers")
            }
            sizeX = sX
            sizeY = sY
            posX = pX
            posY = pY       
        } else {
            fmt.Println("Error : invalid arguments\nAs a reminder : -size = Size of the grid => Input must be like '-size=sizeX,sizeY'\n-pos = Starting position of the ant => Input must be like '-pos=posX,posY'")
        }
	}

    //Setting up variables for the algorithm

	cWhite:=color.Gray{255}
	cBlack:=color.Gray{0}	
	steps= 100000
    // Steps wanted : 10 - 100 - 1000 - 10000 - 100000

	bounds := image.Rect(0, 0, sizeX, sizeY)
	im := image.NewGray(bounds)
	draw.Draw(im, bounds, image.NewUniform(cWhite), image.ZP, draw.Src)
	pos := image.Point{posX, posY}
	im.SetGray(pos.X, pos.Y, cWhite)
    direction := up

	/***** Algorithm implementation *****/
    // Loop on the number of steps 
	for i := 0; i < steps; i++{  
        switch im.At(pos.X, pos.Y).(color.Gray).Y{
        //switch to determine direction depending on the color of the case the ant is on
		case cWhite.Y:
			im.SetGray(pos.X, pos.Y, cBlack)	
			direction++
		case cBlack.Y:
			im.SetGray(pos.X, pos.Y, cWhite)
			direction--

		}

        //switch to determine the ant's movement depending on the direction:
        //With the constants defined above : 1 = right; 2 = down; 3 = left; 4 = up; 5 = right and so on
        //So, when direction is even, we move up/down; when it's odd, we move left/right
		switch direction%2{
		case 0:
            //Even case : if direction is a multiple of 4, that means the ant goes up; else it goes down
			if direction%4 == 0 {
				pos.Y -= 1
			} else {
				pos.Y += 1
			}
		case 1:
            //Odd case : the ant goes left for every number matching 4n+3; else it goes right
			if (direction%4)%3 == 0{
				pos.X -= 1
			} else {
				pos.X += 1
			}

		}
	}

    /***** Graphic generation *****/
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