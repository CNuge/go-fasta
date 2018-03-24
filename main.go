package main

// compile this, for use of the fasta package from a high level

import(
	"fmt"
	"./fasta"
	)

func main(){

	// parse the command line arguments
	flag.Parse() 

	// include arguments for:
		// -s summary:
			// if passed then produce a summary file

		// -n ncbi (batch or single)
			// take either one string, multiple space delimited string or a text file
			// parse the above into a slice of accession numbers, query NCBI for the accession numbers
		

		// -m merge (multiple, output file name)
			// take a list of files space delimited (or a .txt with filenames within)
			// merge them and write to -f or default name

			// this is done by:
				// parsing the filnames,
				// read in the first file with Read

				// for subsequent files, read in with Read, then append then iterate through their
				// Fasta struct and use AddItem to add all the seq to the first Fasta struct

				// take the large Fasta struct and Write it to -f


		// -d split (single, output file names == fasta names)
			// for the split, have it take a list fasta struct, and split each member of 
			// the fasta into its own fasta struct, then take this and write each to a file
			// using the fasta.Write() function with the name of the sequence + ".fasta"
			// passed in as the second name.

			// this is done by:
				// Read to access all the data in the file
				// for each seq in the Fasta struct, make a new Fasta struct and
				// call Write() pass in the seq.name + ".fasta" as the output name for each 

		// -f if passed, change the output file names 
			// for instance of both a .fasta and a summary, 
			// take this name and split on a . , take the first bit and append .fasta and .txt to it and use accordingly
		

}