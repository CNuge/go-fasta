package main

// compile this, for use of the fasta package from a high level

import(
	"fmt"
	"./fasta"
	"flag"
	)




// include arguments for:

	// -m merge (multiple, output file name)
		// take a list of files space delimited (or a .txt with filenames within)
		// merge them and write to -f or default name

		// this is done by:
			// parsing the filnames,
			// read in the first file with Read

			// for subsequent files, read in with Read, then append then iterate through their
			// Fasta struct and use AddItem to add all the seq to the first Fasta struct

			// take the large Fasta struct and Write it to -f


	// -n ncbi (batch or single)
		// take either one string, multiple space delimited string or a text file
		// parse the above into a slice of accession numbers, query NCBI for the accession numbers


	// -f if passed, change the output file names 
		// for instance of both a .fasta and a summary, 
		// take this name and split on a . , take the first bit and append .fasta and .txt to it and use accordingly
	

	// -s summary:
		// if passed then produce a summary file


	// -a alphabetize the sequences in a fasta by name
		// read in the fasta
		// call the sort function on the fasta
		// Write() to the input name

	// -split (single, output file names == fasta names)
		// for the split, have it take a list fasta struct, and split each member of 
		// the fasta into its own fasta struct, then take this and write each to a file
		// using the fasta.Write() function with the name of the sequence + ".fasta"
		// passed in as the second name.

		// this is done by:
			// Read to access all the data in the file
			// for each seq in the Fasta struct, make a new Fasta struct and
			// call Write() pass in the seq.name + ".fasta" as the output name for each 

func main(){

	mergePtr := flag.String("m", nil, "A comma delimited list of fasta filenames to be merged.\n" +
										"The final fasta will contain the sequences in the order of the .fasta inputs.\n" +
										"You an also pass in a .txt filename which contains a list of filnames (all names specified on seprate lines).\n"
										"Use in conjunction with the -f flag to alter the output file name.")

	ncbiPtr := flag.String("n", nil, "A comma delimited list of unique NCBI IDs.\n" +
									"The .fasta files associated with the accession IDs will be downloaded and saved to a .fasta file.\n" +
									"You an also pass in a .txt filename which contains a list of IDs (all specified on seprate lines).\n" +
									"Use in conjunction with the -f flag to alter the output file name.")

	filePtr := flag.String("f", "output.fasta", "A .fasta or .txt filename.\n" + 
												"For use with -m -n and -s flags to speficy an output name.\n" +
												"If both a fasta and summary are needed, just passed a .fasta name\n" + 
												"and it will produce a summary file with the same name and a .txt extension\n")

	alphaPtr := flag.Bool("a", false, "Pass this flag name in conjunction with a -f flag.\n" +
										"Sequences in the -f specified file will be sorted alphabetically by sequence name.\n")


	splitPtr := flag.Bool("split", false, "Pass this flag name in conjunction with a -f flag.\n" +
											"The Sequences in the -f specified file will be split into a set of fasta files, one for each sequence in the file.\n")



	// parse the command line arguments
	flag.Parse() 

	//can then be accessed using:
	*mergePtr
	*ncbiPtr
	//etc

}