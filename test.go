
package main

import(
	"fmt"
	"sort"
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


// sort a Fasta by the seq names, modify it in place
func (fa *Fasta) Sort() {
	// make a dict where the keys are the sequence names
	// and the values are a pointer to the seq structs
	fasta_dict := make(map[string]seq)
	name_list := []string{}

	// value is a pointer, to avoid moving the seq around twice
	for _, s := range fa.entries {
		fasta_dict[s.name] = s

		name_list = append(name_list, s.name)
	}

	// quicksort of the keys (seq.name)
	sort.Strings(name_list)

	// make a new fasta
	out_fasta := Fasta{}
	// append the original seq to the output in the correct order
	for _, i := range name_list {
		to_add := fasta_dict[i]

		out_fasta.AddItem(to_add)
	}
	fmt.Println(out_fasta)

	*fa =  out_fasta
}


func main() {

	test_fasta := Fasta{entries : []seq{seq{name: "ZZZZ", sequence: "CAT"},
									seq{name: "BBBB", sequence: "GC"},
									seq{name: "AAAA", sequence: "ATGC"},
									seq{name: "TTTT", sequence: "AATT"}}}

	test_fasta.Sort()

	fmt.Println(test_fasta)

}
