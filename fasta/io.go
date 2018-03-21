package fasta

import(
	"fmt"
	"flag"
	"strings"
	"log" // for logging errors
	"io/ioutil" //input/output utilities https://golang.org/pkg/io/ioutil/ 
)

// represent a single sequence
type seq struct {
	name string
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

func (fa *Fasta) String() string {
	outstring := ""
	for _ , s := range fa.entries {
		outstring = append(outstring, s.String())
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
	return seq{ name : entry[0],
				sequence : strings.Join(entry[1:], "")}
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
	data := strings.Split(string(file) , ">")
	// the first position is empty because of the leading >
	// so we iterate from 1:end and get the sequence
	// here we parse the fasta and add it to the slice of seq
	for _ , entry := range data[1:] {
		fileseqs.AddItem(ParseFasta(entry))
	}
	return fileseqs
}


// print a sequence in fasta fmt with newline characters
// after every 60 nucleotides
func (sq seq) fileString() string {
	outstring := fmt.Sprintf(">%v\n", sq.name)
	for i = 0 ; i <= len(sq.sequence), i = i + 60 {
		// check if we have reached the end of the sequence
		if i+60 > len(sq.sequence) {
			back := len(sq.sequence) 
		} else {
			back := i+60
		}

		line = fmt.Sprintf("%v\n", sq.sequence[i : back])		
		outstring = append(outstring, line)

	}
	return outstring
}


func Write(fa *Fasta, filename string ) {

	// accepts any filename, if none given
	if filename == nil {
		filename = "output.fasta"
	}

	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	for _ , s := range fa.entries {
		// get string with newlines and write to file
		f.WriteString(s.fileString())
	} 
}

/*

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