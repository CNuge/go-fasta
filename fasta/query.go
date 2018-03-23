package fasta

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


// take the query unique IDs and get string response
func Query(accession UID) seq {
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

    // parse data to string, pass to parser
    return ParseFasta(string(body))

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


