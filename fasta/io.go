// The fasta package is designed to provide a suite of functions for reading, writing and manipulating Fasta sequence files.
// This library can be imported into other go projects to allow for simplified use and creation of Fasta files.
// The Query functions also provide the ability to retrieve new sequence data in fasta format from the
// National Center for Biotechnology Information (NCBI) databases by providing a slice of unique sequence ids.
// Importing this library provides the following specalized data structures and methods:
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

// A struct with two fields, Name and Sequence to represent the two parts of a fasta file entry.
type Seq struct {
	Name     string
	Sequence string
}

// return the sequence in fasta format when called
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

// A slice of Seq types []Seq
// This structure represents a fasta file for the library, it is a set of Seq structures.
// The library's input/output functionality allows for efficient reading and writing of files in standard fasta format
type Fasta []Seq

// Return the Fasta file in fasta format
func (fa Fasta) String() string {
	outstring := ""
	for _, s := range fa {
		outstring = fmt.Sprintf("%v%v", outstring, s.String())
	}
	return outstring
}

// take a raw entry string from a fasta file and build a seq structure
func parseSeq(fasta_entry string) Seq {
	entry := strings.Split(fasta_entry, "\n")
	// first position is the name,
	// join everything but the first line into a single string
	return Seq{Name: entry[0],
		Sequence: strings.Join(entry[1:], "")}
}

// add a seq struct instance to the fasta struct
func (fa *Fasta) AddItem(item Seq) {
	*fa = append(*fa, item)
}

// This function takes one argument, a string specifying the name of a fasta file.
// It returns an Fasta object with the sequence information from the file.
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
		fileseqs.AddItem(parseSeq(entry))
	}
	return fileseqs
}

// This method can be called on an existing Fasta type. It takes one argument, a filename (path optional) to which the Fasta will be written.
// The output is in standard Fasta format, with a header line prefaced by a '>' character and a sequence section with 60 characters of sequence per line.
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
