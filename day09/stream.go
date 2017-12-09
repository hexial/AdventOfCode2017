package day09

import (
	log "github.com/sirupsen/logrus"
)

func ParseStreamString(stream string, doGarbage bool) int {
	var countGarbage int
	var depth int
	var score int
	var garbage bool
	var cancelNext bool
	for _, char := range stream {
		if cancelNext {
			cancelNext = false
		} else if char == '!' {
			cancelNext = true
		} else if garbage && char == '>' {
			garbage = false
		} else if garbage {
			countGarbage++
		} else if char == '<' {
			garbage = true
		} else if char == '{' {
			depth++
		} else if char == '}' {
			score += depth
			depth--
		} else if char == ',' {
			//
			// Group
		} else {
			log.Panicf("Unhandled: %c", char)
		}
	}
	if doGarbage {
		return countGarbage
	}
	return score
}
