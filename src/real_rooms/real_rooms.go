package main

import (
	"sort"
	"strconv"
	"strings"
)

type Room struct {
	name []string
	id string
	checkSum string
}

func RealRooms(roomStr string) int {
	room := Room{}
	room.checkSum = strings.Split(roomStr, "[")[1]
	room.checkSum = strings.TrimRight(room.checkSum , "]")

	nameIdStr := strings.Split(roomStr, "[")[0]
	nameId := strings.Split(nameIdStr, "-")
	room.id = nameId[len(nameId)-1]
	room.name = nameId[:len(nameId)-1]

	// need a map for the room names
	m := make(map[int32]int)
	for _, s := range room.name {
		for _, c := range s {
			if _, ok := m[c]; ok{
				m[c] += 1
			} else{
				m[c] = 1
			}
		}
	}

	// need an integer array to store keys to compare
	keys := make([]int32, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool { return m[keys[i]] > m[keys[j]] })

	// there could be multiple checksums that are all correct so we need to verify the count sequence
	v1 := make([]int, len(m))
	for i, k := range keys {
		v1[i] = m[k]
	}
	v2 := make([]int, len(m))
	for i, k := range room.checkSum {
		v2[i] = m[k]
	}
	for i:=0; i< len(room.checkSum); i++{
		if v1[i] != v2[i] {
			return 0
		}
	}
	id, _ := strconv.Atoi(room.id)
	return id
}

func main() {

}