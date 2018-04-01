package fasta

// contains code for seq and fasta structures
// along with associated input/output methods

import (
	"fmt"
	"io/ioutil" // input/output utilities https://golang.org/pkg/io/ioutil/
	"log"       // for logging errors
	"os"
	"strings"
)

// represent a single sequence
type Seq struct {
	Name     string
	Sequence string
}

// the function to return the sequence in fasta format when printed
func (sq Seq) String() string {
	return fmt.Sprintf(">%v\n%v\n", sq.Name, sq.Sequence)
}

// print a sequence in fasta fmt with newline characters
// after every 60 nucleotides
func (sq Seq) fileString() string {
	outstring := fmt.Sprintf(">%v\n", sq.Name)
	for i := 0; i <= len(sq.Sequence); i = i + 60 {
		// check if we have reached the end of the sequence
		var back int
		if i+60 > len(sq.Sequence) {
			back = len(sq.Sequence)
		} else {
			back = i + 60
		}

		line := fmt.Sprintf("%v\n", sq.Sequence[i:back])
		outstring = fmt.Sprintf("%v%v", outstring, line)

	}
	return outstring
}

// represent a list of sequences as a Fasta
type Fasta []Seq

func (fa Fasta) String() string {
	outstring := ""
	for _, s := range fa {
		outstring = fmt.Sprintf("%v%v", outstring, s.String())
	}
	return outstring
}

// add a seq struct instance to the fasta struct
func (fa *Fasta) AddItem(item Seq) {
	*fa = append(*fa, item)
}

// take a raw entry string from a fasta file and build a seq structure
func ParseFasta(fasta_entry string) Seq {
	entry := strings.Split(fasta_entry, "\n")
	// first position is the name,
	// join everything but the first line into a single string
	return Seq{Name: entry[0],
		Sequence: strings.Join(entry[1:], "")}
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

func (fa Fasta) Write(file ...string) {
	filename := ""
	// accepts filename, if none given makes it default
	if len(file) > 0 {
		filename = file[0]
	} else {
		filename = "output.fasta"
	}

	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	for _, s := range fa {
		// get string with newlines and write to file
		f.WriteString(s.fileString())
	}
}
