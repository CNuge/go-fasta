# go-fasta
## command line based fasta housekeeping
## query NCBI to obtain the .fasta for a given accession number
## merge, split, clean and utilize fasta files in other go programs.


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
- fasta.Read()
- fasta.Write()
- fasta.Query()

- fasta.Merge()
- fasta.Split()
- fasta.Sort() - take multiple sequences in a fasta.File struct and sort them
- fasta.Summary() - get a summary table with GC count and nucleotide count for all sequences in a file

- special fasta.Seq and fasta.File structures.
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

// call this chunk in other workflows, run in parallel to other tasks where posisble
if summary bool != false{
	fasta.WriteSummary()	
}