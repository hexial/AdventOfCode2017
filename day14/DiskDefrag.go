package day14

import (
	"AdventOfCode2017/day10"
	"encoding/hex"
	"fmt"

	log "github.com/sirupsen/logrus"
)

func getBit(b byte, bit uint) int {
	return int(b >> bit & 1)
}

func calcByte(b byte) (int, string) {
	sum := 0
	for i := 0; i < 8; i++ {
		sum += getBit(b, uint(i))
	}
	s := ""
	for i := 7; i >= 0; i-- {
		if getBit(b, uint(i)) == 1 {
			s += "#"
		} else {
			s += "."
		}
	}
	return sum, s
}

func DoRoundUsed(s string) int {
	var used int
	str := day10.NewKnotHashPartTwo(s, 256).DenseHash
	val, err := hex.DecodeString(str)
	if err != nil {
		panic(err)
	}
	for _, b := range val {
		u, _ := calcByte(b)
		used += u
		//log.Infof("%s", s)
	}
	log.Infof("%s : %s : %d", s, str, used)
	return used
}

func CalcUsed(key string) int {
	var used int
	var num int
	num = 128
	for i := 0; i < num; i++ {
		used += DoRoundUsed(fmt.Sprintf("%s-%d", key, i))
	}
	return used
}

func DoRoundRegion(s string, i int, regions *[128][128]int, maxRegion *int) {
	str := day10.NewKnotHashPartTwo(s, 256).DenseHash
	val, err := hex.DecodeString(str)
	if err != nil {
		panic(err)
	}
	s1 := ""
	for _, b := range val {
		_, s2 := calcByte(b)
		s1 += s2
	}
	var r int
	runes := []rune(s1)
	var inRegion bool
	for pos := 0; pos < 128; pos++ {
		//log.Infof("pos=%d, r=%d, inRegion=%t", pos, r, inRegion)
		if !inRegion && runes[pos] == '#' {
			//
			// Going into new region
			r = 0
			inRegion = true
			//log.Infof("In region")
			if pos == 127 {
				if inRegion && i > 0 && regions[i-1][pos] > 0 {
					//
					// Above has region
					r = regions[i-1][pos]
				} else {
					*maxRegion++
					r = *maxRegion
				}
				regions[i][pos] = r
			}
		} else if inRegion && (runes[pos] == '.' || pos == 127) {
			//
			// Going out of region
			if r == 0 {
				*maxRegion++
				r = *maxRegion
			}
			t := pos - 1
			if runes[pos] == '#' {
				t = pos
			}
			for t >= 0 && runes[t] == '#' {
				regions[i][t] = r
				t--
			}
			inRegion = false
			//log.Infof("Out region")
		}
		if inRegion && i > 0 && regions[i-1][pos] > 0 {
			//
			// Above has region
			r = regions[i-1][pos]
			//log.Info("Above region: %d", r)
		}
	}
	//log.Infof("%s", s)
	//log.Infof("%s", s1)
	//log.Infof("%v", regions[i])
}

func CalcRegions(key string) int {
	var num int
	var maxRegion int
	var regions [128][128]int
	num = 128
	for i := 0; i < num; i++ {
		DoRoundRegion(fmt.Sprintf("%s-%d", key, i), i, &regions, &maxRegion)
	}
	return maxRegion
}
