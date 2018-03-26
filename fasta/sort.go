// take a fasta struct and sort the entries into alphabetical order by their name
// improvement -> sort the sequence in place

package fasta

import(
	"fmt"
	"sort"
	)


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