# go-fasta
## command line based fasta housekeeping
## query NCBI to obtain the .fasta for a given accession number
## merge, split, clean and utilize fasta files in other go programs.


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
- fasta.Summary() - get a 

- special fasta.Seq and fasta.File structures.
- sequences are