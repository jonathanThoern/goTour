package main

import (
		"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	img := make([][]uint8, dy)
	for i := range img{
		img[i] = make([]uint8, dx)
	}
	
	for i := range img{
		for j := range img[i]{
			//img[i][j] = uint8(4*(i+j)/3)
			img[i][j] = uint8(i^(j*i))
		}
	}	
	return img
}

func main() {
	pic.Show(Pic)
}
