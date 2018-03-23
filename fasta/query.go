package fasta

// this file contains the code for interacting with NCBI 
// and downloding a fasta sequence for a provided accession number
// the downloaded sequences can be loaded into a Fasta struct, or 
// printed directly to a file

import(
    "fmt"
    "net/http"
    "strings"
    "log"
    "io/ioutil"
    "os"
    "io"
)


type UID struct {
    all []string
}

func buildURL(accession UID) string {
    url_front := "https://eutils.ncbi.nlm.nih.gov/entrez/eutils/efetch.fcgi?db=nucleotide&id="
    
    // build the middle of the url from the input slice
    url_middle := ""
    for _, i := range accession.all {
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
        fileseqs.AddItem(ParseFasta(entry))
    }
    return fileseqs
}

// take the query unique IDs and get string response
func Query(accession UID) Fasta {
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


func QueryToFile(accession UID, output string) error {
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

    // Write data to file
    _, err = io.Copy(out, resp.Body)
    if err != nil {
        return err
    }

    return nil
}


