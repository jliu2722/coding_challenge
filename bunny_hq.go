package main

import (
	"math"
	"strconv"
	"strings"
)

type Point struct {
	dir int
	x, y int
}

const (
	north int = iota
	east
	south
	west
)
func BunnyHQ(input string) (result float64) {
	p := Point{north,0, 0}

	inputs := strings.Split(input, ",")
	for _,v := range inputs {
		if string(v[0]) == "R" {
			p.dir = (p.dir + 1) % 4
		} else {
			p.dir = (p.dir - 1) % 4
		}
		count, _ := strconv.Atoi(v[1:])
		switch p.dir {
		case north:
			for i := 0; i < count; i++ {
				p.y += 1
			}
		case east:
			for i := 0; i < count; i++ {
				p.x += 1
			}
		case south:
			for i := 0; i < count; i++ {
				p.y -= 1
			}
		case west:
			for i := 0; i < count; i++ {
				p.x -= 1
			}
		}
	}
	return math.Abs(float64(p.x)) + math.Abs(float64(p.y))
}

func main(){
}
