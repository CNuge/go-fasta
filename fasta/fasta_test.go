package fasta



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