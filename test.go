
package main

import(
    "fmt"
    "net/http"
    "strings"
    "log"
    "io/ioutil"	
    "reflect"
)


// represent a single sequence
type Seq struct {
	name     string
	sequence string
}


// represent a list of sequences as a Fasta
type Fasta struct {
	entries []Seq
}


// add a seq struct instance to the fasta struct
func (fa *Fasta) AddItem(item Seq) []Seq {
	fa.entries = append(fa.entries, item)
	return fa.entries
}




// take a raw entry string from a fasta file and build a seq structure
func ParseFasta(fasta_entry string) Seq {
	entry := strings.Split(fasta_entry, "\n")
	// first position is the name,
	// join everything but the first line into a single string
	return Seq{name: entry[0],
		sequence: strings.Join(entry[1:], "")}
}

func main() {

}
