package main

import(
	"fmt"
	"./fasta"
	)

func main(){

	// parse the command line arguments
	flag.Parse() 

	// include arguments for:
		// -f reformat (to pretty print)
		// -s summary
		// -q query (batch or single)
		// -m merge (multiple, output file name)
		// -s split (sinlge, output file names == fasta names)

		// option to specify output filename if want it to be different from
		// the default



	// for the split, have it take a list fasta struct, and split each member of 
	// the fasta into its own fasta struct, then take this and write each to a file
	// using the fasta.Write() function with the name of the sequence + ".fasta"
	// passed in as the second name.

}