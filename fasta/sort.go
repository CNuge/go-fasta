// take a fasta struct and sort the entries into alphabetical order by their name
// improvement -> sort the sequence in place

package fasta

import (
	"sort"
)

// sort a Fasta by the seq names, modify it in place
func (fa *Fasta) Sort() {
	// make a dict where the keys are the sequence names
	// and the values are a pointer to the seq structs
	fasta_dict := make(map[string]int)
	name_list := []string{}

	// value is a index in original entries struct,
	// avoids moving the seq into intermediate structure
	for i, s := range fa.entries {
		fasta_dict[s.name] = i
		name_list = append(name_list, s.name)
	}

	// quicksort of the keys (seq.name)
	sort.Strings(name_list)

	// make a new fasta
	out_fasta := Fasta{}
	// append the original seq to the output in the correct order
	for _, k := range name_list {
		original_pos := fasta_dict[k]

		out_fasta.AddItem(fa.entries[original_pos])
	}

	*fa = out_fasta
}
