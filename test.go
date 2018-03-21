package main


import (
	"fmt"
	"io/ioutil" //input/output utilities https://golang.org/pkg/io/ioutil/
	"log"       // for logging errors
	"os"
	"strings"
)

// represent a single sequence
type seq struct {
	name     string
	sequence string
}

// the function to return the sequence in fasta format when printed
func (sq seq) String() string {
	return fmt.Sprintf(">%v\n%v\n", sq.name, sq.sequence)
}

// represent a list of sequences as a Fasta
type Fasta struct {
	entries []seq
}

func (fa Fasta) String() string {
	outstring := ""
	for _, s := range fa.entries {
		outstring = fmt.Sprintf("%v%v", outstring, s.String())
	}
	return outstring
}



// add a seq struct instance to the fasta struct
func (fa *Fasta) AddItem(item seq) []seq {
	fa.entries = append(fa.entries, item)
	return fa.entries
}

// take a raw entry string from a fasta file and build a seq structure
func ParseFasta(fasta_entry string) seq {
	entry := strings.Split(fasta_entry, "\n")
	// first position is the name,
	// join everything but the first line into a single string
	return seq{name: entry[0],
		sequence: strings.Join(entry[1:], "")}
}

func Read(filename string) Fasta {
	fileseqs := Fasta{} // start an empty Fasta instance
	//Opening a file
	file, err := ioutil.ReadFile(filename)
	// check if that caused an error
	if err != nil {
		log.Fatal(err)
	}
	// split the input file on the new seq characters
	data := strings.Split(string(file), ">")
	// the first position is empty because of the leading >
	// so we iterate from 1:end and get the sequence
	// here we parse the fasta and add it to the slice of seq
	for _, entry := range data[1:] {
		fileseqs.AddItem(ParseFasta(entry))
	}
	return fileseqs
}




type summaryDat struct {
	name   string
	length int
	gc     float64
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
	bp := 0
	gc := 0
	// below we check to make sure the bases are ATGC 
	// to skip the N on the GC count
	for _, base := range sq.sequence {
		if base == 'G' || base == 'C' {
			bp++
			gc++
		} else if base == 'A' || base == 'T'{
			bp++
		}

	}
	return float64(gc) / float64(bp) * 100.0
}

// get an output slice containing the summaryDat for each of the sequences
func (fa Fasta) Summary() []summaryDat {
	output := []summaryDat{}
	// iterate through the entries in the fasta structure
	for _, entry := range fa.entries {
		data := summaryDat{ entry.name, entry.len(), entry.percGC()}
		output = append(output, data)
	}
	return output
}

// a wrapper function to write the output summary to a file
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


func main() {
	test := Read("./example_data/example1.fasta")
	test.WriteSummary("putotu.txt")
}
