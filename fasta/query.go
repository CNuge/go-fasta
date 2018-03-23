package fasta

// take an accession number, query NCBI and then return and use the data
// treat the return from NCBI like a file, send it to



// this is done using eutils - get a sense of the parts of the records needed
// https://www.ncbi.nlm.nih.gov/books/NBK25500/#chapter1.Downloading_Full_Records

// base url:
fetch := "https://eutils.ncbi.nlm.nih.gov/entrez/eutils/efetch.fcgi"


/*
Required Parameters
db
Database from which to retrieve records. 
The value must be a valid Entrez database name (default = pubmed). 
Currently EFetch does not support all Entrez databases. 
Please see Table 1 in Chapter 2 for a list of available databases.

Required Parameter – Used only when input is from a UID list
id
UID list. Either a single UID or a comma-delimited list of UIDs may be provided. 
All of the UIDs must be from the database specified by db. 
There is no set maximum for the number of UIDs that can be passed to EFetch, 
but if more than about 200 UIDs are to be provided, the request should be made 
using the HTTP POST method.

efetch.fcgi?db=protein&id=15718680,157427902,119703751
Required Parameters – Used only when input is from the Entrez History server
query_key
Query key. This integer specifies which of the UID lists attached to the given 
Web Environment will be used as input to EFetch. Query keys are obtained from 
the output of previous ESearch, EPost or ELInk calls. The query_key parameter 
must be used in conjunction with WebEnv.

WebEnv
Web Environment. This parameter specifies the Web Environment that contains 
the UID list to be provided as input to EFetch. Usually this WebEnv value is 
obtained from the output of a previous ESearch, EPost or ELink call. 
The WebEnv parameter must be used in conjunction with query_key.

efetch.fcgi?db=protein&query_key=<key>&WebEnv=<webenv string>
Optional Parameters – Retrieval
retmode
Retrieval mode. This parameter specifies the data format of the records returned, 
such as plain text, HMTL or XML. See Table 1 for a full list of allowed values for 
each database.

rettype
Retrieval type. This parameter specifies the record view returned, such as Abstract or MEDLINE from PubMed, or GenPept or FASTA from protein. Please see Table 1 for a full list of allowed values for each database.


*/