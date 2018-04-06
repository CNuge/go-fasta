package fasta

// this contains the code for making a summary report of a fasta file
// a summary struct can be generated, or written to an output file

import (
	"fmt"
	"os"
)

type summaryDat struct {
	name   string
	length int
	gc     float64
}

// represent the summary data structures for printing
func (sd summaryDat) String() string {
	return fmt.Sprintf("%v\t%v\t%.2f\n", sd.name, sd.length, sd.gc)
}

// get the length of a seq
func (sq Seq) len() int {
	return len(sq.Sequence)
}

// gc content of a seq
func (sq Seq) percGC() float64 {
	bp := 0
	gc := 0
	// below we check to make sure the bases are ATGC
	// to skip the N on the GC count
	for _, base := range sq.Sequence {
		if base == 'G' || base == 'C' {
			bp++
			gc++
		} else if base == 'A' || base == 'T' {
			bp++
		}else if base != 'N' && warning == false {
			warning = true
		}
	}

	if warning == true {
		fmt.Printf("Warning! The fasta file entry: %v appears to have non-nucleotide sequence.\n Ignore the Perc_GC content in the summary file!\n", sq.Name )
		return 0.0
	}
	
	return float64(gc) / float64(bp) * 100.0
}

// This method should be used with nucleotide Fasta structures only.
// Calling this method returns a slice of structs with three fields: 
// the name, the length and the percent GC content of the sequences in the Fasta
func (fa Fasta) Summary() []summaryDat {
	output := []summaryDat{}
	// iterate through the entries in the fasta structure
	for _, entry := range fa {
		data := summaryDat{entry.Name, entry.len(), entry.percGC()}
		output = append(output, data)
	}
	return output
}

// This method has the same functionality as the Summary method, but instead of 
// providing the output slice with the summary data in memory, it writes the summary 
// directly to the file specified as a string in the method call.
func (fa Fasta) WriteSummary(file ...string) {
	filename := ""

	if len(file) > 0 {
		filename = file[0]
	} else {
		filename = "summary.tsv"
	}

	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// get the summary data
	sum_data := fa.Summary()
	// write header to file
	header_string := "Name\tLen\tPerc_GC\n"
	f.WriteString(header_string)

	// iterate through the rows of the summary table
	// write the data to file in .tsv fmt
	for _, row := range sum_data {
		f.WriteString(row.String())
	}
}
