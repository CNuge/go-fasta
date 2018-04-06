package fasta

/*
	this file contains the code for interacting with NCBI
	and downloding a fasta sequence for a provided accession number
	the downloaded sequences can be loaded into a Fasta struct, or
	printed directly to a file
*/

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func buildURL(accession []string) string {
	url_front := "https://eutils.ncbi.nlm.nih.gov/entrez/eutils/efetch.fcgi?db=nucleotide&id="

	// build the middle of the url from the input slice
	url_middle := ""
	for _, i := range accession {
		url_middle = fmt.Sprintf("%v,%v", url_middle, i)
	}

	url_end := "&rettype=fasta&retmode=text"

	url := []string{url_front, url_middle, url_end}

	return strings.Join(url, "")

}

func fastaFromQuery(raw_data string) Fasta {
	// parse data to string, pass to parser
	fileseqs := Fasta{}
	// split the input file on the new seq characters
	data := strings.Split(raw_data, ">")
	// the first position is empty because of the leading >
	// so we iterate from 1:end and get the sequence
	// here we parse the fasta and add it to the slice of seq
	for _, entry := range data[1:] {
		fileseqs.AddItem(parseSeq(entry))
	}
	return fileseqs
}

// This function takes a slice of strings as an argument, where each of the strings is an NCBI accession number.
// It will query NCBI for these accession numbers, and return a Fasta type instance containing the a Seq struct corresponding to each of the accession numbers.
func Query(accession []string) Fasta {
	//construct the url
	query_url := buildURL(accession)

	// make the http request
	resp, err := http.Get(query_url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// the the response data to variable
	body, err := ioutil.ReadAll(resp.Body)

	// parse response body to a fasta structure
	return fastaFromQuery(string(body))

}

// This function can be used in lieu of the Query function in instances where the data are not required
// in memory, they can then be written directly to a file (this is more efficient as the data does not
// need to be processed into the Fasta structure and the string can be written straight to the file).
// The function takes two argumens. The first argument is a slice of strings where each of the strings
// is an NCBI accession number. The second argument is a string containing the desired output file name
// to which the sequences obtained in the NCBI query will be written.
func QueryToFile(accession []string, output string) error {
	// construct the url
	query_url := buildURL(accession)

	//make the file
	out, err := os.Create(output)
	if err != nil {
		return err
	}
	defer out.Close()

	// make the http request
	resp, err := http.Get(query_url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Write data direct to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
