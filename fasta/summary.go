

type summaryDat struct{
	name string
	length int
	gc int
}

// represent the summary data structures for printing
func (sd summaryDat) String() string{
	return fmt.Sprintf("%v\t%v\t%v\n", sd.name, sd.length, sd.gc)
}

// get the length of a seq
func (sq seq) len() int {
	return len(sq.sequence)
}

// gc content of a seq
func (sq seq) percGC() float64 {
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

// get an output slice containing the summaryDat for each of the sequences
func Summary(fa *Fasta) []summaryDat {
	output = []summaryDat{}
	// iterate through the entries in the fasta structure
	for _, entry := range fa.entries{
		data := summaryDat{name: entry.name, len(entry), percGC(entry)}
		output = append(output, data)
		}
	return output
}
