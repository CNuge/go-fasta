package fasta

import (
	"fmt"
	"reflect"
	"testing"
)

/*
The unit tests are broken down into test functions which correspond to the
package file names.
*/

func TestIO(t *testing.T){

}


func TestQuery(t *testing.T){

	AF298042 := Fasta{entries: []Seq{ Seq{	name: "AF298042.1 Salvelinus alpinus alpinus haplotype ARCTIC_16 mitochondrial D-loop, partial sequence",
											sequence: "CCACTAATATGTACAATAATGAATATTGTATCTCAACAAATTAGTGTCATAATACATCTATGTATAATATTGCATATTATGTATTTACCCATATATACAATACCTGTATGATGAGTAGTACATCATATGTATTATCAACATAAGTGAATTTAAGCCCTCATATATCAGCATAAACCCAAGATTTACATAAGCTAAACACGTGATAATAACCAACTAGGTTGTTTTAACCTAGATAATTGCTACATTAACAAAACTCCAACTAACACGGGCTCCGTCTTTACCCACCAACTTTCAGCATCAGTCCTACTTAATGTAGTAAGAACCGACCAACGATTTATCAGTAGGCATACTCTTATTGATGGTGAGGGGCAAATATCGTATTAGGTAACATCTCGTGAACTATTCCTGGCATTTGGTTCCTAAGTCGAGGGCTATCCTTAAGAAACCAGCCCCTGAAAGCCGAATGTTAAGCATCTGGTTAATGGTGTCAATCTTATTGTTCGTTACCCACAAAGCCGGGCGTTCTCTTATATGCATAGGGTTCTCCTTT"},
										}}

	single_id := []string{"AF298042.1"}
	out1 := Query(single_id)
	
	if reflect.DeepEqual(out1, AF298042) != true {
		t.Errorf("Single Query of Accession Number from NCBI did not match expected output. Received:\n %v", out1)
	}

	AF298042_AY646679 := Fasta{entries : []Seq{ Seq{	name: "AF298042.1 Salvelinus alpinus alpinus haplotype ARCTIC_16 mitochondrial D-loop, partial sequence",
														sequence: "CCACTAATATGTACAATAATGAATATTGTATCTCAACAAATTAGTGTCATAATACATCTATGTATAATATTGCATATTATGTATTTACCCATATATACAATACCTGTATGATGAGTAGTACATCATATGTATTATCAACATAAGTGAATTTAAGCCCTCATATATCAGCATAAACCCAAGATTTACATAAGCTAAACACGTGATAATAACCAACTAGGTTGTTTTAACCTAGATAATTGCTACATTAACAAAACTCCAACTAACACGGGCTCCGTCTTTACCCACCAACTTTCAGCATCAGTCCTACTTAATGTAGTAAGAACCGACCAACGATTTATCAGTAGGCATACTCTTATTGATGGTGAGGGGCAAATATCGTATTAGGTAACATCTCGTGAACTATTCCTGGCATTTGGTTCCTAAGTCGAGGGCTATCCTTAAGAAACCAGCCCCTGAAAGCCGAATGTTAAGCATCTGGTTAATGGTGTCAATCTTATTGTTCGTTACCCACAAAGCCGGGCGTTCTCTTATATGCATAGGGTTCTCCTTT"},
												Seq{	name: "AY646679.1 Spironucleus barkhanus from wild Arctic charr small subunit ribosomal RNA gene, partial sequence",
														sequence: "AAGATTAAAGCCCTGCATGCCTATGTGTAGACAGTTATATTCATTATTGTGGAGCAAAAACGGCGAACAGCTCATTTATCAGTGGTAAGTGCATACAATGTATTTCGTTGGATAGTAACGGAAAATCTGTTAGTAATACATGAACTGTTTTTAGCATTATGTTAAAAATAATAGTAAGTGCGATTGTATATCTGCCACTGCAGCATCATCTTACGTTGGTGGGATATTTGCCTACCAAGGATTCGACGCTTACGGGGAATTAGGGTTTGACTCCGGAGAATGAGCATGAGAAACAGCTCATACATCTAAGGAAGGCAGCAGGCGCGGAAATTGCCCAATGTATCTTTTATACGAGGCAGTGACAAGAAATGGTAGGCACTTTTGTGCACTATCGAGGGTTAGTGGTATCTTTGCTAACCGTGACTCGTGGGCAAGCTCGGTGCCAGCAGCCGCGGTAATTCCGACACAGGGAGTTTTCCATTTGGTTGCTGCAGTTAAAAAGTTCGTAGTTTACTGACTCTTTCACTATAAGCAAAGCCGAATGCTCCAAGTTTTTTAGCAGTATTTATAGTATGAAATTATAGCGCGGCATTGAACGTAGTTTGGGGTACTCGATAGGGACAGGTGAAATAGGATGATCTATCGAAGACCCACGGTAGCGGAGGCTCCCAACGAAGTCCAAGTGTCACGATCAAGAACTAAAGTCAGGGGATAGACGACGATTAGACACCGTTTTATTCCTGACCCTAAACGATGTCGCCTAGCTGATGGGATTTTTTTCATTTGCCAAGAGAAATCGTAAGGTTTCAGACTCTGGGGGAAGTATGATCGCAAGGTTGAAACTTGAAGGGATTGACGGAGAGGTACCACCAGACGTGGAGTCTGCGGCTCAATTTGACTCAACACGCAAACATTACTAGGCCCAGAAGCTTTGAGGATTGACAGATGAGTGATCTTTCATGATTAAGTTGTTGGTGGTGCATGGCCGTTCTTAGTCCGTGATTTAAATTGTCTGCTTTATTGCGATAACGAACGAGACCTCTATCAGATTTATTATCTGAGACTGCTAGTGATGAACTAGAGGAAGGCAGAGGCAAAAACAGGTCTGTGATGCCCTTAGAAGCCCTAGGCCGCACGCGTACTACAATGGCAGGTTCATCGTGTTGCTTCCCTGAAAATGGTGGCAGTTCATTAAAACTTGTCGTGGTTAGGACTGAAGGTTGAAATTATCCTTCACGAATGAGGAATGTCTAGTAAGTGTAGGTTATGAATCTACGCTGATTACGTCCCTACCCCTTGTACACACCGCCCGTCGCTCCTACTGATTGGGAAGATCTGGTGAGTTATTCGGACCCATAGGTAAGCAATTATCTGTGGTAACAATTGCGAGCCAACTCTTCTAGAGGAAGG"},
												}}

	list_of_ids := []string{"AY646679.1","AF298042.1" }
	out2 := Query(list_of_ids)

	if reflect.DeepEqual(out2, AF298042_AY646679) != true {
		t.Errorf("Multiple Query of Accession Number from NCBI did not match expected output. Received:\n %v", out2)
	}
}

func TestSort(t *testing.T){

	unsorted := Fasta{entries : []Seq{	Seq{name: "sdsdsd", sequence: "CAT"},
										Seq{name: "chr1", sequence: "GC"},
										Seq{name: "1chr", sequence: "ATGC"},
										Seq{name: "chr2", sequence: "AATT"}}}

	sorted := Fasta{entries : []Seq {	Seq{name: "1chr", sequence: "ATGC"},
										Seq{name: "chr1", sequence: "GC"},
										Seq{name: "chr2", sequence: "AATT"},
										Seq{name: "sdsdsd", sequence: "CAT"}}}
	
	unsorted.Sort()			
	if reflect.DeepEqual(unsorted, sorted) != true {
		t.Errorf("Sorting of Fasta structure incorrect: %v\n want: %v.", unsorted, sorted)
	}
}

// to this one, add an error that gets thrown when an AA seq is passed in
// 
func TestSummary(t *testing.T){
	//test_seq1, len = 12, gc = .5
	//test_seq2, len = 27, gc = .93
	//test_seq3, len = 18, gc = .89
	//test_seq4, len = 14, gc = 1.00
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
	exp_output := []summaryDat{ summaryDat{ name : ,
											length : ,
											gc : , } ,
								summaryDat{ name : ,
											length : ,
											gc : , } ,
								summaryDat{ name : ,
											length : ,
											gc : , } ,
								summaryDat{ name : ,
											length : ,
											gc : , } }
	
	test_output := test_in.Summary()

	if reflect.DeepEqual(test_output , exp_output) != true {
		t.Errorf("Summary of Fasta incorrect: %v\n want: %v.", test_output, exp_output)
	}

}



// REMOVE BELOW WHEN DONE


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