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


// represent a list of sequences
type Fasta struct {
	entries []seq
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


//Write fasta
//TODO - need to print the sq.sequence with line breaks and spaces in it!
//		- determine the desired number of nucleotides per line and the breakpoints
//		- code in a for loop to pretty write the sequences to file

func (sq seq) PrettyString() string {
	outstring := fmt.Sprint(">%v\n", sq.name)

	for i = 10 ; i <= len(sq.sequence), i = i + 10{
		// if i + 10 greater then length, change the back side of window to len(sq.sequence)

		// if 10/60 has no remainder, append a newline
		
		// else append a space

		// append whole string to outstring

	}
	
	//for loop through the sq.sequence
		// every 10 bp put a space,
		// every 60 basepairs put a newline
		// end with a newline
	// need to make it stop and provide the end of the sequence if not multiple of 10

	// change the stringer to do the every 10 newline
	return 
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

	// loop through all of the seq structs in the fasta,
		// send each one to the PrettyString() and then write the string to the file

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