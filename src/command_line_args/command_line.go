package main	

import "os"
import "flag"
import "strings"
import "fmt"

func main() {
	arguments := os.Args

	fmt.Println(arguments)
	

	sizeArg := flag.String("size", "100,100", "Size of the grid\nInput must be like '-size=sizeX,sizeY'")
	posArg := flag.String("pos", "50,50", "Starting position of the ant\nInput must be like '-pos=posX,posY'")

	flag.Parse()

	fmt.Println(*sizeArg)
	fmt.Println(*posArg)

	size := strings.Split(*sizeArg, ",")
	pos := strings.Split(*posArg, ",")

	fmt.Println("size : ", size)
	fmt.Println("position : ", pos)

	sizeX := size[0]
	sizeY := size[1]
	posX := pos[0]
	posY := pos[1]

	fmt.Println("SizeX = ", sizeX, " ; SizeY = ", sizeY," ; PosX = ", posX," ; PosY= ", posY)

}