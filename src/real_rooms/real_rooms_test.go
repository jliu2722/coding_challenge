package main

import (
	"fmt"
	"testing"
)

func TestRealRooms(t *testing.T) {
	var sum [4]int
	sum[0] = RealRooms("aaaaa-bbb-z-y-x-123[abxyz]")
	sum[1] = RealRooms("not-a-real-room-404[oarel]")
	sum[2] = RealRooms("a-b-c-d-e-f-g-h-987[abcde]")
	sum[3] = RealRooms("totally-real-room-200[decoy]")
	total := 0
	for _, v := range sum{
		total += v
	}
	if total != 1514 {
		fmt.Println("Tests failed!")
	}
}