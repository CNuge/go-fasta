
package main

import(
	"fmt"
	"net/http"
	"strings"
	"log"
	"io/ioutil"
)


type UID struct {
	all []string
}


func (accession UID) buildURL() string {
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


// take the query unique IDs and write them to the output fasta
func (accession UID) Query() string {

    query_url := accession.buildURL()

    resp, err := http.Get(query_url)
    if err != nil {
            log.Fatal(err)
    } 
		
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return string(body)

}



func main() {
	single := UID{[]string{"AY646679.1"}}
	out1 := single.Query()
	fmt.Println(out1)
	list_of_ids := UID{[]string{"AY646679.1", "AF298042.1"}}
	out2 := list_of_ids.Query()
	fmt.Println(out2)
}
