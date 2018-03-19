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





//TODO - need to print the sq.sequence with line breaks and spaces in it!
//		- determine the desired number of nucleotides per line and the breakpoints
//		- code in a for loop to pretty write the sequences to file