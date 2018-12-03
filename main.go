package main

import (
	"fmt"
	trianglen "./triangle-app"
)

func main() {
	fmt.Println(trianglen.CheckTriangleType(5.0,6,1.5))
	fmt.Println(trianglen.CheckTriangleType(5,5,1))

}
