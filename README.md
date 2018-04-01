# go-fasta
## command line based fasta housekeeping
## query NCBI to obtain the .fasta for a given accession number
## merge, split, clean and utilize fasta files in other go programs.


TODO:
- unit tests for io file
- documentation for the package in the proper idiomatic go fmt
- test all of the cmd line functionality, make sure parallel parts work as expected
- documentation for the repository. update the readme file
- clean up and delete the unnecessary helper code
- format all the files with go fmt
- make the travis.yml file and have the build pass
- compiled version for download? /makefile?
- summary file throw warning if there are non ATGCN bp in the sequences



goal:
	have an executable file that can be compiled and used without a need to touch
	the internals, and also a package that can be imported elsewhere and provide 
	several simple commands for interacting with fasta files.


first:
- go through and assemble the read and write functions for the fasta package



to do:

from cmd line:
- download fasta from NCBI when provided with an accession number, numbers or a list
	of multiple in a .txt file
- split a fasta into multiple files, named by their header line
	have option of .fa or .fasta extension (.fasta as default)
- concatenate multiple fasta files into a single fasta
- pretty print the sequences when writing to files, i.e. spaces every x number 
	of nucleotides
- summary, get an output .csv file with the names of all the fasta sequences in a file
	the length and GC content of each.

for the fasta package:
- fasta.ParseFasta()
- fasta.Read()
- fasta.Write()
- fasta.Query()
- fasta.QueryToFile()
- fasta.Merge()
- fasta.Split()
- fasta.Sort() - take multiple sequences in a fasta.File struct and sort them
- fasta.Summary() - get a summary table with GC count and nucleotide count for all sequences in a file
- fasta.WriteSummary()

- special fasta.Seq and fasta.Fasta structures.
- sequences are





## Notes on main


// include arguments for:


// -m merge (multiple, output file name)
	// take a list of files space delimited (or a .txt with filenames within)
	// merge them and write to -f or default name

	// this is done by:
		// parsing the filnames,

		// read in the files with Read run via parallel goroutine
		// merge them all to a single Fasta struct by sending the goroutines to the same struct

		// take the large Fasta struct and Write it to -f
		// if summary true, run it in parallel to the Write()
	

// if a test file was passed with a flag, open a file and load in the 
// newline separated data and split to a slice of strings.



// -n ncbi (batch or single)
	// take either one string, multiple space delimited string or a text file
	// parse the above into a slice of accession numbers, and make UID struct

	// if summary true, call Query() then run Write()  in parallel to the Write()
	
	// if summary == false
	// query NCBI for the accession numbers and write via the QueryToFile func
	// it is direct and faster




// -a alphabetize the sequences in a fasta by name
	// read in the fasta
	// call the sort function on the fasta
	// Write() to the input name (parallel with the summary if needed)


// -split (single, output file names == fasta names)
	// for the split, have it take a list fasta struct 
	// goroutine that takes each seq in the fasta into its own fasta struct, 
	// and take this and write each to a file
	// using the fasta.Write() function with the name of the sequence + ".fasta"
	// passed in as the second name.

	// this is done by:
		// Read to access all the data in the file
		// goroutine that for each seq in the Fasta struct, make a new Fasta struct and
		// call Write() pass in the seq.name + ".fasta" as the output name for each 
		// summary parallel to goroutine above



// -f if passed, change the output file names 
	// for instance of both a .fasta and a summary, 
	// take this name and split on a . , take the first bit and append .fasta and .txt to it and use accordingly


// -s summary:
	// if passed then produce a summary file

	mergePtr := flag.String("m", "__none__", "Merge Fastas. A comma delimited list of fasta filenames to be merged.\n" +
										"The final fasta will contain the sequences in the order of the .fasta inputs.\n" +
										"You an also pass in a .txt filename which contains a list of filnames (all names specified on seprate lines).\n"
										"Use in conjunction with the -f flag to alter the output file name.")

	ncbiPtr := flag.String("n", "__none__", "Query NCBI. A comma delimited list of unique NCBI IDs.\n" +
										"The .fasta files associated with the accession IDs will be downloaded and saved to a .fasta file.\n" +
										"You an also pass in a .txt filename which contains a list of IDs (all specified on seprate lines).\n" +
										"Use in conjunction with the -f flag to alter the output file name.")

	alphaPtr := flag.Bool("a", false, "Alphabetize Fasta. Pass this flag name in conjunction with a -f flag.\n" +
										"Sequences in the -f specified file will be sorted alphabetically by sequence name.\n")

	splitPtr := flag.Bool("split", false, "Split Fasta. Pass this flag name in conjunction with a -f flag.\n" +
											"The Sequences in the -f specified file will be split into a set of fasta files, one for each sequence in the file.\n")
	
	summaryPtr := flag.Bool("summary", false, "Make a summary file of output. Pass this flag and a summary file will be constructed which\n" +
												"gives the following information for each sequence in the fasta produced:\n" +
												"sequence name\t sequence length\t percent gc content\n" + 
												"IMPORTANT NOTE: summary is designed for use with nucleotide based fasta files" +
												"if you call it on a protein sequence fasta file the gc content column will be nonsense!")

	filePtr := flag.String("f", "output.fasta", "File name. A .fasta or .txt filename.\n" + 
												"For use with -m -n -a -split and -summary flags to specify an output name.\n" +
												"If both a fasta and summary are needed, just passed a .fasta name\n" + 
												"and it will produce a summary file with the same name and a .txt extension\n")
