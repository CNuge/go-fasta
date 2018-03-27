package main

// compile this, for use of the fasta package from a high level

import(
	"fmt"
	"./fasta"
	"flag"
	"log"
	"strings"
	"io/ioutil"
	"sync"
	)


func parseFlagData(flagDat string) []string {
	flag_input := strings.Split(flagDat, ",") //split on commas

}


func parseFlagFileData(filename string) []string {
	//Opening a file
	file, err := ioutil.ReadFile(filename)
	// check if that caused an error
	if err != nil {
		log.Fatal(err)
	}
	data := strings.Split(file, "\n")
	// remove leading and trailing strings if they exist
	if data[len(data)-1] == ""{
		data = data[:len(data)-1]
	}
	if data[0] == ""{
		data = data[1:]
	}
	return data
}


func parseFastaFlags(flagDat string) []string {
	// split the flags
	flag_input := parseFlagData(flagDat)

	// iterate through the flags, handling .txt and .fastas separately
	fasta_files := []string{}
	for _ , i := range flag_input {
		if i[len(i)-4:] == ".txt" {
			add_fastas := parseFlagFileData(i)
			fasta_files = append(fasta_files, add_fastas...)
		} else if i[len(i)-6:] == ".fasta" || i[len(i)-3:] == ".fa" {
			fasta_files = append(fasta_files, i)
		} else {
			err := fmt.Errorf("You have passed an filname to -m that does not have a .txt, .fasta or .fa extension.")
			log.Fatal(err)	
		}
	} 
	return fasta_files
}

func parseNCBIFlagData( flagDat string) []string {
	// split the flags
	flag_input := parseFlagData(flagDat)

	// iterate through the flags, handling .txt and .fastas separately
	ncbi_ids := []string{}
	for _ , i := range flag_input {
		// check if it is a text file, if so pass to the flag file reader
		if i[len(i)-4:] == ".txt" {
			add_ids := parseFlagFileData(i)
			ncbi_ids = append(ncbi_ids, add_ids...)
		} else {
			ncbi_ids = append(ncbi_ids, i)
		}
	} 
	return ncbi_ids	
}

func getSummaryName(name string) string {
	// note here the len(name) =< 6 is to avoid a bug of filenames shorter than '.fasta'
	if name[len()name-3:] == ".fa" {
		return name[:len()name-3] + ".txt"
	} else if len(name) =< 6 {
		return name
	} else if name[len(name)-6:] == ".fasta" {
		return name[:len(name)-6] + '.txt'
	} else {
		return name
	}
}

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
// if a test file was passed with a flag, open a file and load in the 
// newline separated data and split to a slice of strings.


// this is the file read function used as a goroutine, it runs the 
// fasta.Read() function and passes the data to the channel provided
// it also requires a sync.WaitGroup{} to be initialized and referenced 
// so that it can report completion of the read file
func readToCh(filename string, ch chan fasta.Fasta, wg *sync.WaitGroup){

	// read the data in from the file
	fasta_dat := fasta.Read(filename)

	// push the final object down the channel
    ch <- fasta_dat
    // let the wait group know we finished
    wg.Done()
}

func mergeWorkFlow( merge_data string, file_data string, summary bool) {
	fasta_list := parseFastaFlaga(merge_data)
	
	// an output slice with len 0 and capacity equal to length of the fasta list
	// output_fasta.entries is empty to start, but we can append to it
	output_fasta := fasta.Fasta{}

	// a buffered channel the length of the number of input fastas
	ch := make(chan fasta.Fasta, len(fasta_list))
    
    // waitgroup - basically block until N tasks say they are done
    wg := sync.WaitGroup{}

    for i=0 ; i<len(fasta_list); i++ {
    	// say there is one more routine to wait for
    	wg.Add(1)
    	// spawn a goroutine to run the Read function
    	// this recieves the filename, the output channel and a pointer to the waitgroup as inputs
    	go readToCh(fasta_list[i], ch, &wg)
    }
    // wait for everyone to finish
    wg.Wait()
    // close the channel so that we can 
    close(ch)

    // iterate through the channel data
    for file_dat := range ch{
    	//append each of the entries to the output entries list
    	output_fasta.entries = append(output_fasta.entries, file_dat.entries...)
    }

    // go call to summary
    if summary == true {
    	summary_name := getSummaryName(file_data)
    	go output_fasta.WriteSummary(summary_name)
    }
	
	// go call to write concurrent to summary
	go output_fasta.Write(file_data)
}


// -n ncbi (batch or single)
	// take either one string, multiple space delimited string or a text file
	// parse the above into a slice of accession numbers, and make UID struct

	// if summary true, call Query() then run Write()  in parallel to the Write()
	
	// if summary == false
	// query NCBI for the accession numbers and write via the QueryToFile func
	// it is direct and faster



func ncbiWorkflow( ncbi_data string,  file_data string, summary bool) {
	accessions := parseNCBIFlagData(ncbi_data)

}

// -a alphabetize the sequences in a fasta by name
	// read in the fasta
	// call the sort function on the fasta
	// Write() to the input name (parallel with the summary if needed)

func aplhaWorkflow(file_data string, summary bool) {
	fasta_file := fasta.Read(file_data)
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
	fasta_file := fasta.Read(file_data)

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
		err := fmt.Errorf("You may only pass in one of the following three flags at a time: -m -n -a -split.\n" + 
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