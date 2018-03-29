package fasta

import (
	"fmt"
	"reflect"
	"testing"
)


	single := UID{[]string{"AY646679.1"}}
	out1 := Query(single)
	fmt.Println(out1)
	list_of_ids := UID{[]string{"AY646679.1", "AF298042.1"}}
	out2 := Query(list_of_ids)
	fmt.Println(out2)

	QueryToFile(list_of_ids, "outfile.fasta")


// For sort test

	test_fasta := Fasta{entries : []Seq{Seq{name: "ZZZZ", sequence: "CAT"},
								Seq{name: "BBBB", sequence: "GC"},
								Seq{name: "AAAA", sequence: "ATGC"},
								Seq{name: "TTTT", sequence: "AATT"}}}

	test_fasta.Sort()

	fmt.Println(test_fasta)


// a test syntax example I have from elsewhere
func TestQueue(t *testing.T) {
	q := Queue{}
	compare_q := Queue{ord: []int{7, 8, 9}}
	compare_q2 := Queue{ord: []int{8, 9}}

	q.Add(7)
	q.Add(8)
	q.Add(9)

	if reflect.DeepEqual(q.ord, compare_q.ord) != true {
		t.Errorf("Adding to Queue incorrect: %v, want: %v.", q, compare_q)
	}
}