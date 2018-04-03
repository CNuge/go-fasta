# go-fasta
## Command line based fasta housekeeping (merge, split, sort and summarize)and NCBI queries via sequence identifiers.
[![Build Status](https://travis-ci.org/CNuge/go-fasta.svg?branch=master)](https://travis-ci.org/CNuge/go-fasta)

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

Note below the use of a bulk read in. mention the fmt of the file, it must have the paths
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
`sequence name	sequence_length	percent_gc_content`
IMPORTANT NOTE: summary is designed for use with nucleotide based fasta files, if you call it on a protein sequence fasta file the gc content column will be nonsense!

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
This package simplifies fasta file interaction within go. It can be imported other golang code using the command:
```import "github.com/CNuge/go-fasta/fasta"```
* A copy of the fasta package's documentation is shown below. It can also be obtained using `godoc`, follwed by the path to the package (i.e. `godoc ./fasta from this directory`)

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

