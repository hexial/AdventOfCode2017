package day17

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

func init() {
	//log.SetLevel(log.DebugLevel)
}

type BufferItem struct {
	Next *BufferItem
	Prev *BufferItem
	val  int
}

type CircularBuffer struct {
	ItemZero *BufferItem
	Curr     *BufferItem
	steps    int
}

func NewCircularBuffer(steps int, rounds int) *CircularBuffer {
	cb := new(CircularBuffer)
	cb.steps = steps
	cb.ItemZero = new(BufferItem)
	cb.ItemZero.Next = cb.ItemZero
	cb.ItemZero.Prev = cb.ItemZero
	cb.Curr = cb.ItemZero
	for i := 0; i < rounds; i++ {
		cb.Insert(i + 1)
		if i%1000000 == 0 {
			log.Infof("i=%d", i)
		}
	}
	return cb
}

func (cb *CircularBuffer) Insert(val int) {
	//
	// Step forward x times
	for i := 0; i < cb.steps; i++ {
		cb.Curr = cb.Curr.Next
	}
	//
	//
	log.Debugf("Insert : val=%d", cb.Curr.val)
	//
	// Insert value
	left := cb.Curr.Prev
	rigth := cb.Curr
	n := new(BufferItem)
	n.val = val
	left.Next = n
	rigth.Prev = n
	n.Prev = left
	n.Next = rigth
	//
	cb.Curr.Prev = n
	cb.State()
	//log.Debugf("Insert (done): pos=%d : val=%d : %v", cb.pos, cb.buffer[cb.pos], cb.buffer)
}

func (cb *CircularBuffer) NextVal() int {
	//log.Debugf("len of buffer: %d", len(cb.buffer))
	return cb.Curr.val
}

func (cb *CircularBuffer) ValPos1() int {
	return cb.ItemZero.Next.val
}

func (cb *CircularBuffer) State() {
	str := ""
	done := false
	curr := cb.ItemZero
	for !done {
		str += fmt.Sprintf("%d,", curr.val)
		if curr.Next == cb.ItemZero {
			done = true
		} else {
			curr = curr.Next
		}
	}
	log.Debugf("state=%s", str)
}
