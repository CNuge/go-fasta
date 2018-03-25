package main

// compile this, for use of the fasta package from a high level

import(
	"fmt"
	"./fasta"
	"flag"
	"log"
	)

//CAM - try to find some of the functions that you can refactor into goroutines
 // split function is prime for this!
 // so is merge


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
	
func mergeWorkFlow( merge_data string, file_data string, summary bool) {

}


// -n ncbi (batch or single)
	// take either one string, multiple space delimited string or a text file
	// parse the above into a slice of accession numbers, and make UID struct

	// if summary true, call Query() then run Write()  in parallel to the Write()
	
	// if summary == false
	// query NCBI for the accession numbers and write via the QueryToFile func
	// it is direct and faster


func ncbiWorkflow( ncbi_data string,  file_data string, summary bool) {

}

// -a alphabetize the sequences in a fasta by name
	// read in the fasta
	// call the sort function on the fasta
	// Write() to the input name (parallel with the summary if needed)

func aplhaWorkflow(file_data string, summary bool) {

}


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


func splitWorkflow(file_data string, summary bool) {
	// call the summary on the whole thing, parallel to the split
}

// -f if passed, change the output file names 
	// for instance of both a .fasta and a summary, 
	// take this name and split on a . , take the first bit and append .fasta and .txt to it and use accordingly


// -s summary:
	// if passed then produce a summary file

// call this chunk in other workflows, run in parallel to other tasks where posisble
if summary bool != false{
	fasta.WriteSummary()	
}





// refactor needed - the bools for the most part work on files. have them accept a file name to work


func main(){

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
												"sequence name\t sequence length\t percent gc content\n")

	filePtr := flag.String("f", "output.fasta", "File name. A .fasta or .txt filename.\n" + 
												"For use with -m -n -a -split and -summary flags to specify an output name.\n" +
												"If both a fasta and summary are needed, just passed a .fasta name\n" + 
												"and it will produce a summary file with the same name and a .txt extension\n")

	// parse the command line arguments
	flag.Parse() 

	//can then be accessed using:
	*mergePtr
	*ncbiPtr
	//etc

	// need to then determine the workflow.

	// count the passed pointers from exclusive list, if more then one, log an error
	ex_count := 0
	if *mergePtr != "__none__" {
		ex_count++
	} 
	if *ncbiPtr != "__none__"{
		ex_count++
	} 
	if *alphaPtr != false{
		ex_count++
	}
	if *splitPtr != false {
		ex_count++
	} 

	// if multiple or 0 exclusive flags passed, raise an error	
	if ex_count > 1{
		err := fmt.Errorf("You may only pass in one of the followin three flags at a time: -m -n -a -split.\n" + 
							"They cannot function in conjunction with one another."+ 
							"Use the help flag: -h for argument use and formatting help.")
		log.Fatal(err)
	}
	if ex_count == 0{
		err := fmt.Errorf("You must use one of the following flags: -m -n -a -split.\n" + 
							"Use the help flag: -h for argument use and formatting help.")
		log.Fatal(err)
	}

	// call the proper workflow based on the flags passed
	if *mergePtr != "__none__" {
		mergeWorkFlow(*mergePtr, *filePtr, *summaryPtr)
	}
	
	if *ncbiPtr != "__none__"{
		ncbiWorkflow(*ncbiPtr, *filePtr, *summaryPtr)
	} 
	
	if *alphaPtr != false {
		 aplhaWorkflow(*filePtr, *summaryPtr)
	}
	if *splitPtr != false {
		splitWorkflow(*filePtr, *summaryPtr)
	} 

}