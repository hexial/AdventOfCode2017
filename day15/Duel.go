package day15

func Duel(a *Generator, b *Generator, pairs int64) int {
	var matches int
	var i int64
	for i = 0; i < pairs; i++ {
		valA := a.Next()
		valB := b.Next()
		//log.Infof("valA=%d valB=%d", valA, valB)
		if valA&0xFFFF == valB&0xFFFF {
			matches++
			//log.Infof("Match")
		} else {
			//log.Infof("Missmatch")
		}
	}
	return matches
}
