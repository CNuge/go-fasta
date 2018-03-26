package fasta



	single := UID{[]string{"AY646679.1"}}
	out1 := Query(single)
	fmt.Println(out1)
	list_of_ids := UID{[]string{"AY646679.1", "AF298042.1"}}
	out2 := Query(list_of_ids)
	fmt.Println(out2)

	QueryToFile(list_of_ids, "outfile.fasta")


// For sort test

	test_fasta := Fasta{entries : []seq{seq{name: "ZZZZ", sequence: "CAT"},
								seq{name: "BBBB", sequence: "GC"},
								seq{name: "AAAA", sequence: "ATGC"},
								seq{name: "TTTT", sequence: "AATT"}}}

	test_fasta.Sort()

	fmt.Println(test_fasta)