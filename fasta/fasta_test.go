package fasta



	single := UID{[]string{"AY646679.1"}}
	out1 := Query(single)
	fmt.Println(out1)
	list_of_ids := UID{[]string{"AY646679.1", "AF298042.1"}}
	out2 := Query(list_of_ids)
	fmt.Println(out2)

	QueryToFile(list_of_ids, "outfile.fasta")