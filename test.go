
package main

import(
    "fmt"
	"reflect"
	"io/ioutil"
	"log"
	"strings"
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

type summaryDat struct {
	name   string
	length int
	gc     float64
}


// get the length of a seq
func (sq Seq) len() int {
	return len(sq.sequence)
}

// gc content of a seq
func (sq Seq) percGC() float64 {
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

// represent the summary data structures for printing
func (sd summaryDat) String() string {
	return fmt.Sprintf("%v\t%v\t%.2f\n", sd.name, sd.length, sd.gc)
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

func main() {
	test_1 := Fasta{ []Seq{ Seq{ name:"SAC1",
							sequence:"TGCATGTTGGAAACATGGCCCTGGCATATGTCTATCTTTCTCTCGGTCTCTAGGGCTAAGCTTCTCTTCCTTATATTTTT"},
						Seq{ name:"SAC2",
							sequence:"TGCATTACTTAGAGTGGTAGGTCAGTAGGGACGGTGCCTAACGTGTGAATCTCAAATGACACACAATTCCTTAAACATAGTACACGTC"},
						Seq{ name:"SAC3",
							sequence:"TGCATAAGGCTACCATCTAGGCGAACGTTATATTGGAATGGAATATGCAAGATCGGAAGAGCGGTTCAGCAGGACCGAGACC"},
						Seq{ name:"SAC4",
							sequence:"TGCATTAGGGGTGTGTGTGTGTGTGTGTGTGTGTGTTTGGTATGCAAGATCGGAAGAGCGGTTCAGCAGGAATGCCGAGACCATC"},
						Seq{ name:"SAC5",
							sequence:"TGCATGCTTTTCAGATTTTATTTTGTGTGGATCCCCATTGAGGCCTTTCCTCGCGTTCATCAGAAATGTATAGAGATGGCTTCCA"},
						Seq{ name:"SAC6",
							sequence:"TGCATTGATGCTAAATAGGGCCTCAGCTTGACAATTCATTCAAGTAACAAATGTTTGGATTCCTGATTTTGATTTTTTTTTCTTAACATTT"},
						Seq{ name:"SAC7",
							sequence:"TGCATTCCTACTTGGACACGTTACAGTAGATGTAACAACCCACTTTGTAAGCCTCATACTACCGACATGCGTAATTAGAAAAAGAGAGAGTTAAG"},
						Seq{ name:"SAC8",
							sequence:"TGCATGATTTTTCACCAGTACTTGAAAATGTGAATAAAACCGTAAATACGTTACGCTCCATACATGCAAGATCGGAAGAGCGGTT"},
						Seq{ name:"SAC9",
							sequence:"TGCATGTGTGTATACATATCTGTACTGTCTCGCTTGACAGGCTTACAGTGGCACTGGGAATGATACTTATGCCTGTAATGTGTTT"},
						Seq{ name:"SAC10",
							sequence:"TGCATTTAATCTTATCCGATAGGCGAAATCTAAATAGATAACTTTGGAAAAACTGGGCCCAGAGGCTTACTGAGGATTAAACTAC"},
						}}

	ex_file := "./example_data/example1.fasta"
	
	test_in := Read(ex_file)

	if reflect.DeepEqual(test_1, test_in) != true {
		fmt.Printf("Test of fasta reader produced incorrect data. Received: %v\n", test_in)
	}
	fmt.Println(test_1)
	fmt.Println(test_in)


}
