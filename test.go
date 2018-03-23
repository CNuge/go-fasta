
package main

import(
	"fmt"
	"net/http"
)



func buildURL(UID ... string) string {
	url_front := "https://eutils.ncbi.nlm.nih.gov/entrez/eutils/efetch.fcgi?db=nucleotide&id="
    
    // build the middle of the url from the input slice
    url_middle := ""
    for _, i := range UID{
        url_middle = Sprintf("%v,%v", url_middle, i)
    }
	
    url_end := "&rettype=fasta&retmode=text"

    url := []string{url_front, url_middle, url_end}

    return strings.Join(url, "")

}


// take the query unique IDs and write them to the output fasta
func Query( UID ... string ) {

    query_url := buildURL(UID)

    response, err := http.Get(query_url)
    if err != nil {
            log.Fatal(err)
    } else {
        return response
    }
}



func main() {
	fmt.Println(Query("AY646679.1"))
	list_of_ids := []string{"AY646679.1", "AF298042.1"}
}
