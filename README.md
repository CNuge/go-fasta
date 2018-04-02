# go-fasta
## Command line based fasta housekeeping (merge, split, sort and summarize), also queries NCBI to obtain fasta files with sequence corresponding to given accession numbers.

# UNDER CONSTRUCTION - library works but I'm still working on documentation!
Still TO DO:
- documentation for the package in the proper idiomatic go fmt
- test all of the cmd line functionality, make sure parallel parts work as expected
- documentation for the repository. update the readme file
- clean up and delete the unnecessary helper code
- format all the files with go fmt
- make the travis.yml file and have the build pass
- compiled version for download? /makefile?
- summary file throw warning if there are non ATGCN bp in the sequences


######################################################################

## What is in the repository?

## 2. The fasta package 
* Documentation is shown below, can also be called using `godoc` follwed by the path to the package (i.e. `godoc ./fasta`)

The fasta package is designed to provide a suite of functions for reading, writing and manipulating Fasta sequence files. This library can be imported into other go projects to allow for simplified use and creation of Fasta files. The Query functions also provide the ability to retrieve new sequence data in fasta format from the National Center for Biotechnology Information (NCBI) databases by providing a slice of unique sequence ids.
Importing this library provides the following specalized data structures and methods:


#### type `Seq`
* a struct with two fields, Name and Sequence to represent the two parts of a fasta file entry.

#### type `Fasta` 
* a slice of Seq types []Seq
* This structure represents a fasta file for the library, it is a set of Seq structures.
* The library's input/output functionality allows for efficient reading and writing of files in standard fasta format

#### func Read(filename string)
* This function takes one argument, a string specifying the name of a fasta file.
* The function returns an object of the type Fasta

#### func (fa \*Fasta) AddItem(item Seq)
* This method can be called on an existing Fasta type. It takes one argument, specifying a new Seq structure to be added to the Fasta (this is useful for merging multiple Fasta files or adding newly obtained data to an existing Fasta)
* The Fasta will be modified in place (this is a pointer reciever method).

#### func (fa Fasta) Write(filename)
* This method can be called on an existing Fasta type. It takes one argument, a filename (path optional) to which the Fasta will be written.
* The output is in standard Fasta format, with a header line prefaced by a '>' character and a sequence section with 60 characters of sequence per line.

#### func (fa \*Fasta) Sort()
* This method can be called on a Fasta type instance to sort the underlying sequences alphabetically, by the Name fields of the constituent Seqs.
* The Fasta will be modified in place (this is a pointer reciever method).

#### func Query(accession)
* This function takes a slice of strings as an argument, where each of the strings is an NCBI accession number. 
* It will query NCBI for these accession numbers, and return a Fasta type instance containing the a Seq struct corresponding to each of the accession numbers. 

#### func QueryToFile(accession []string, output string)
* This function can be used in lieu of the Query function in instances where the data are not required in memory, they can then be written directly to a file (this is more efficient as the data does not need to be processed into the Fasta structure and the string can be written straight to the file).
* The function takes two argumens. The first argument is a slice of strings where each of the strings is an NCBI accession number. The second argument is a string containing the desired output file name to which the sequences obtained in the NCBI query will be written.

#### func (fa Fasta) Summary()
* This method should be used with nucleotide Fasta structures only. 
* Calling this method will produce a slice of structs with three fields, corresponding to the name, length and percent GC content of the sequences in the Fasta


#### func (fa Fasta) WriteSummary(filename)
* This method has the same functionality as the Summary method, but instead of providing the output slice with the summary data in memory, it writes the summary directly to the file specified as a string in the method call.





######################################################################


goal:
	have an executable file that can be compiled and used without a need to touch
	the internals, and also a package that can be imported elsewhere and provide 
	several simple commands for interacting with fasta files.



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
