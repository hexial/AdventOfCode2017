package day17

import (
	log "github.com/sirupsen/logrus"
)

func init() {
	//log.SetLevel(log.DebugLevel)
}

type BufferItem struct {
	Next *BufferItem
	val  int
}

type CircularBuffer struct {
	ItemZero *BufferItem
	Curr     *BufferItem
	Prev     *BufferItem
	steps    int
}

func NewCircularBuffer(steps int, rounds int) *CircularBuffer {
	cb := new(CircularBuffer)
	cb.steps = steps
	cb.ItemZero = new(BufferItem)
	cb.ItemZero.Next = cb.ItemZero
	cb.Curr = cb.ItemZero
	cb.Prev = cb.Curr
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
		cb.Prev = cb.Curr
		cb.Curr = cb.Curr.Next
	}
	//
	//
	log.Debugf("Insert : val=%d", cb.Curr.val)
	//
	// Insert value
	n := new(BufferItem)
	n.val = val
	//
	//
	n.Next = cb.Curr
	cb.Prev.Next = n
	//
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
	/*
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
	*/
}
