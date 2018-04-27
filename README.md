# go-fasta
## Command line based fasta housekeeping (merge, split, sort and summarize)and NCBI queries via sequence identifiers.
[![Build Status](https://travis-ci.org/CNuge/go-fasta.svg?branch=master)](https://travis-ci.org/CNuge/go-fasta)
[![GoDoc](https://godoc.org/github.com/CNuge/go-fasta/fasta?status.svg)](https://godoc.org/github.com/CNuge/go-fasta/fasta)

## What is in the repository?
This repository contains 1. the go-fasta program, and 2. the fasta go package, which can be imported and used within other go programs. These two components are discussed below (documentation provided).

## 1. go-fasta program 
The go-fasta command line executable file provides the ability to efficiently execute common fasta file housekeeping tasks. The program uses concurrency to speed up process execution where possible. 
### install
To get the program up and running on your computer, download this repository and run the following commands ([you will need to have go installed](https://golang.org/) [and your gopath set up properly](https://github.com/golang/go/wiki/SettingGOPATH)): 
```
go get github.com/CNuge/go-fasta
go install github.com/CNuge/go-fasta
```
go-fasta is controlled through a series of user specified command line options which allow for:
* fasta file merger (`-m`)
* fasta file splitting (`-s`)
* fasta file creation from NCBI query (`-n`)
* fasta file summary (`-summary`)
* fasta file sorting (`-a`)
This functionality is controlled by the following command line flags (example use cases given).

### `-m` Merge Fasta files
The merge flag takes a comma delimited list of fasta filenames to be merged. The final fasta will contain the sequences in the order of the .fasta inputs.
You an also pass in a .txt filename which contains a list of filnames (all names specified on seprate lines).
Use in conjunction with the -f flag to alter the output file name (default: output.fasta).

```
go-fasta -m ./example_data/example1.fasta,./example_data/example2.fasta
```

Note below the use of a bulk read in. The file contains a list of fasta files, each listed on a separate line. In a case where you are using files not in your current directory, you will need to specify the full path to each of the files.
```
go-fasta -m ./example_data/filelist.txt -f test_out.fa
```

### `-n` Query NCBI. 
This flag takes a comma delimited list of unique NCBI IDs. The .fasta files associated with the accession IDs will be downloaded and saved to a .fasta file. You can also pass in a .txt filename which contains a list of IDs (all specified on seprate lines). 
Use in conjunction with the -f flag to alter the output file name. Note: this will run significantly faster if not called in conjunction with the -summary flag, as this requires the data to be stored in memory instead of written directly to the file.

Query NCBI for the sequence associated with the unique accession ID 'AY646679.1'. Save this to a fasta file in the current working directory named 'custom_name.fasta'
```
go-fasta -n AY646679.1 -f custom_name.fasta
```

Query NCBI for multiple accesion IDs at once. They will all be saved to the same file. We also specify we want a summary file, which will be produced under the name 'custom_name.txt' and contain tab separated summary info (see the summary flag section).
```
go-fasta -n GL949779.1,GL949780.1 -f custom_name.fasta -summary
```

You can also specify a file containing a list of multiple accession numbers (each on a separate line). The program will read the accession numbers in, and build a sinlge output with the sequence for all the passed ids.
```
go-fasta -n ./example_data/accessionlist.txt -f multi_test.fa
```

### `-a` Alphabetize Fasta
Pass this flag name in conjunction with a -f flag. Sequences in the -f specified file will be sorted alphabetically by sequence name.

```
go-fasta -f ./example_data/example1-unsorted.fasta -a
```

### `-split` Split Fasta
Pass this flag name in conjunction with a -f flag.
The Sequences in the -f specified file will be split into a set of fasta files, one for each sequence in the file.

```
go-fasta -f ./example_data/example1.fasta -split
```

### `-summary`
Make a summary file of output. Pass this flag and a summary file will be constructed which gives the following information for each sequence in the fasta produced: 
`sequence name	sequence_length	percent_gc_content sequence_type`
IMPORTANT NOTE: summary is designed for use with nucleotide based fasta files, if you call it on a protein sequence the gc content column be 0.00 and the sequence_type column will flag the sequence as an amino acid!

To get a summary while performing another task, just add the summary flag at the end of the execution
```
go-fasta -m ./example_data/filelist.txt -f test_out.fa -summary
```
If for a given file you only wanted a summary, you could do the following:
```
go-fasta -m ./example_data/example1.fasta -f ./example_data/example1.fasta -summary
```


### `-f` File name.
A .fasta or .txt filename. For use with -m -n -a -split and -summary flags to specify an output name.
If both a fasta and summary are needed, just pass in a .fasta name and it will produce a summary file with the same name and a .txt extension.


## 2. The fasta package

This package simplifies fasta file interaction within go. After running `go get github.com/CNuge/go-fasta`, this package can be imported other golang code using the command:
```
import "github.com/CNuge/go-fasta/fasta"
```

* A copy of the fasta package's documentation can be found by clicking the button below:

[![GoDoc](https://godoc.org/github.com/CNuge/go-fasta/fasta?status.svg)](https://godoc.org/github.com/CNuge/go-fasta/fasta)

