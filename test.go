
package main

import(
    "fmt"
	"reflect"
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


func main() {

	test_in := Fasta{entries : []Seq{	Seq{name: "test_seq1", 
											sequence: "ATGCATGCATGC"},
										Seq{name: "test_seq2", 
											sequence: "ATATATATATATATATATATAAAAAGC"},
										Seq{name: "test_seq3", 
											sequence: "GCGCGCGCATGCGCGCGC"},
										Seq{name: "test_seq4", 
											sequence: "GGGCGGGCGGGCCC"},
								}}

	// Correct #s for test summary sequence
	//test_seq1, len = 12, gc = .5
	//test_seq2, len = 27, gc = .93
	//test_seq3, len = 18, gc = .89
	//test_seq4, len = 14, gc = 1.00
	exp_output := []summaryDat{ summaryDat{ name : "test_seq1",
											length : 12,
											gc : 50.00} ,
								summaryDat{ name : "test_seq2",
											length : 27,
											gc : 7.41} ,
								summaryDat{ name : "test_seq3",
											length : 18,
											gc : 88.89} ,
								summaryDat{ name : "test_seq4",
											length : 14,
											gc : 100.00}}
	
	test_output := test_in.Summary()

	test_output_string := ""
	for _ , i := range test_in.Summary(){
		test_output_string = fmt.Sprintf("%v%v", test_output_string , i.String())
	}

	exp_output_string := ""
	for _ , i := range test_in.Summary(){
		exp_output_string = fmt.Sprintf("%v%v", exp_output_string , i.String())
	}

	fmt.Println(exp_output)

	fmt.Println(test_output)

	if reflect.DeepEqual(test_output_string , exp_output_string) != true {
		fmt.Printf("Summary of Fasta incorrect: %v\n want: %v.", test_output, exp_output)
	}

}
