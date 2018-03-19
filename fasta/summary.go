



// get the length of a seq
func (sq seq) len() int {
	return len(sq.sequence)
}


// gc content of a seq
func (sq seq) PercGC() float64 {
	bp := len sq.sequence
	gc := 0
	for _ , base := range sq.sequence{
		if base == 'G' || base == 'C' {
			gc++
		}
	}
	return float64(gc)/float64(bp) * 100.0
}


// have the default output be just to io, but give the option to pass in a file
// name and have the program write to a file