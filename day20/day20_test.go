package day20

import (
	"AdventOfCode2017/util"
	"fmt"
	"math"
	"strings"
	"testing"
)

type Vector3 struct {
	X int
	Y int
	Z int
}

func NewVector3(s string) Vector3 {
	a := strings.IndexRune(s, '<') + 1
	b := strings.IndexRune(s, '>')
	s2 := strings.Split(s[a:b], ",")
	var v Vector3
	v.X = util.Atoi(s2[0])
	v.Y = util.Atoi(s2[1])
	v.Z = util.Atoi(s2[2])
	util.LogDebugf("s=%s > %d,%d,%d", s, v.X, v.Y, v.Z)
	return v
}

func (v *Vector3) Debug() string {
	return fmt.Sprintf("%d,%d,%d", v.X, v.Y, v.Z)
}

func (v *Vector3) Add(o Vector3) {
	v.X += o.X
	v.Y += o.Y
	v.Z += o.Z
}

type Particle struct {
	Position     Vector3
	Velocity     Vector3
	Acceleration Vector3
}

func (p *Particle) Increase() {
	p.Velocity.Add(p.Acceleration)
	p.Position.Add(p.Velocity)
	util.LogDebugf("p=%s v=%s", p.Position.Debug(), p.Velocity.Debug())
}

func (v *Vector3) Equal(p Vector3) bool {
	return v.X == p.X && v.Y == p.Y && v.Z == p.Z
}

func (p *Particle) SamePosition(o *Particle) bool {
	return p.Position.Equal(o.Position)
}

func (p *Particle) Distance() int {
	return int(math.Abs(float64(p.Position.X))) + int(math.Abs(float64(p.Position.Y))) + int(math.Abs(float64(p.Position.Z)))
}

func NewParticle(s1 string) *Particle {
	p := new(Particle)
	s2 := strings.Split(s1, " ")
	p.Position = NewVector3(s2[0])
	p.Velocity = NewVector3(s2[1])
	p.Acceleration = NewVector3(s2[2])
	return p
}

func Part1(filename string) int {
	particles := make([]*Particle, 0)
	//
	// Parse
	for _, s := range util.FileAsLineArray(filename) {
		particles = append(particles, NewParticle(s))
	}
	//
	//
	for i := 0; i < 10000; i++ {
		util.LogDebugf("**********************************'")
		for _, p := range particles {
			p.Increase()
		}
	}
	//
	//
	var min int
	var index int
	for i, p := range particles {
		if min == 0 {
			min = p.Distance()
			index = 0
		} else {
			if p.Distance() < min {
				min = p.Distance()
				index = i
			}
		}
	}
	return index
}

func count(list []*Particle, particle *Particle) int {
	var c int
	for _, p := range list {
		if particle.SamePosition(p) {
			c++
		}
	}
	return c
}

func remove(l1 []*Particle, p *Particle) []*Particle {
	l2 := make([]*Particle, 0)
	for i := 0; i < len(l1); i++ {
		if !p.SamePosition(l1[i]) {
			l2 = append(l2, l1[i])
		}
	}
	return l2
}

func collider(l1 []*Particle) (int, []*Particle) {
	i := 0
	removed := 0
	for {
		util.LogDebugf("i=%d len=%d removed=%d", i, len(l1), removed)
		if count(l1, l1[i]) > 1 {
			removed += count(l1, l1[i])
			l1 = remove(l1, l1[i])
			i = 0
		} else {
			i++
			if i == len(l1) {
				return removed, l1
			}
		}
	}
}

func Part2(filename string) int {
	particles := make([]*Particle, 0)
	//
	// Parse
	for _, s := range util.FileAsLineArray(filename) {
		particles = append(particles, NewParticle(s))
	}
	for i := 0; i < 1000; i++ {
		util.LogDebugf("**********************************'")
		//
		// Increase all
		for _, p := range particles {
			p.Increase()
		}
		var removed int
		removed, particles = collider(particles)
		if removed > 0 {
			i = 0
		}
	}
	return len(particles)
}

func Test1(t *testing.T) {
	util.Debug = false
	util.AssertEqual(t, 0, Part1("sample.1"))
	util.AssertEqual(t, 172, Part1("input"))
}

func Test2(t *testing.T) {
	util.Debug = false
	util.AssertEqual(t, 1, Part2("sample.2"))
	util.AssertEqual(t, 502, Part2("input"))
}
