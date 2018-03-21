//problem note - the WriteSummary currently builds the whole string in memory
// and then writes it to the file
// this is inefficient, especially once files get large
// find a way to make the strings for each line and write them to the file on the fly

package fasta

import (
	"fmt"
	"io"
	"os"
)

type summaryDat struct {
	name   string
	length int
	gc     int
}

// represent the summary data structures for printing
func (sd summaryDat) String() string {
	return fmt.Sprintf("%v\t%v\t%v\n", sd.name, sd.length, sd.gc)
}

// get the length of a seq
func (sq seq) len() int {
	return len(sq.sequence)
}

// gc content of a seq
func (sq seq) percGC() float64 {
	bp := len(sq.sequence)
	gc := 0
	for _, base := range sq.sequence {
		if base == 'G' || base == 'C' {
			gc++
		}
	}
	return float64(gc) / float64(bp) * 100.0
}

// get an output slice containing the summaryDat for each of the sequences
func Summary(fa *Fasta) []summaryDat {
	output = []summaryDat{}
	// iterate through the entries in the fasta structure
	for _, entry := range fa.entries {
		data := summaryDat{name: entry.name, len(entry), percGC(entry)}
		output = append(output, data)
	}
	return output
}

// a wrapper function to write the output summary to a file
func WriteSummary(fa *Fasta, filename string) {
	if filename == nil {
		filename = "summary.tsv"
	}

	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// get the summary data
	sum_data = Summary(fa)
	// write header to file
	header_string = "Name\tLen\tPerc_GC\n"
	f.WriteString(header_string)

	// iterate through the rows of the summary table
	// write the data to file in .tsv fmt
	for _, row := range sum_data {
		f.WriteString(row.String())
	}
}

/* on the fly alternative:

f, err := os.Create(filename)
if err != nil {
	panic(err)
}
defer f.Close()


sum_data = Summary(fa)
header_string = "Name\tLen\tPerc_GC\n"
_, err := f.WriteString(header_string)
if err != nil {
	panic(err)
}

for _ , row := range(sum_data){
	row_string = fmt.Sprintf("%v", row)

	_, err := f.WriteString(row_string)
	if err != nil {
		panic(err)
	}

}
*/
