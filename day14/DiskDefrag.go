package day14

import (
	"AdventOfCode2017/day10"
	"fmt"
)

func hasBit(n int, pos uint) bool {
	val := n & (1 << pos)
	return (val > 0)
}

func CalcUsed(key string) int {
	var used int
	for i := 0; i < 128; i++ {
		for _, r := range day10.NewKnotHashDay14(fmt.Sprintf("%s-%d")) {
			var b byte
			if r == '1' {
				b = 0x1
			} else if r == '2' {
				b = 0x2
			} else if r == '3' {
				b = 0x3
			} else if r == '4' {
				b = 0x4
			} else if r == '5' {
				b = 0x5
			} else if r == '6' {
				b = 0x6
			} else if r == '7' {
				b = 0x7
			} else if r == '8' {
				b = 0x8
			} else if r == '9' {
				b = 0x9
			} else if r == 'a' {
				b = 0xa
			} else if r == 'b' {
				b = 0xb
			} else if r == 'c' {
				b = 0xc
			} else if r == 'd' {
				b = 0xd
			} else if r == 'e' {
				b = 0xe
			} else if r == 'f' {
				b = 0xf
			}
			for i := 0; i < 4; i++ {
				if hasBit(int(b), uint(i)) {
					used++
				}
			}
		}
	}
	return used
}
