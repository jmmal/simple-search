package main

import (
	"strings"
	"os"
	"bufio"
	"fmt"
	"github.com/jmmal/simple-search/index"
)

func main() {
	i := index.NewIndex()

	doc1 := index.Doc{
		Id: 1,
		Name: "Josh Maloney",
		Comment: "Nulla justo. Aliquam quis turpis eget elit sodales scelerisque. Mauris sit amet eros. Suspendisse accumsan tortor quis turpis. Sed ante. Vivamus tortor.",
	}
	  
	doc2 := index.Doc{
		Id: 2,
		Name: "Noni Beining",
		Comment: "Fusce consequat. Nulla nisl. Nunc nisl. Duis bibendum, felis sed interdum venenatis, turpis enim blandit mi, in porttitor pede justo eu massa. Donec dapibus. Duis at velit eu est congue elementum.",
	}
	doc3 := index.Doc{
		Id: 3,
		Name: "Sella Presland",
		Comment: "Quisque id justo sit amet sapien dignissim vestibulum. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia Curae; Nulla dapibus dolor vel est. Donec odio justo, sollicitudin ut, suscipit a, feugiat et, eros. Vestibulum ac est lacinia nisi venenatis tristique. Fusce congue, diam id ornare imperdiet, sapien urna pretium nisl, ut volutpat sapien arcu sed augue. Aliquam erat volutpat. In congue. Etiam justo.",
	}
	doc4 := index.Doc{
		Id: 4,
		Name: "Cathrin Legges",
		Comment: "Fusce posuere felis sed lacus. Morbi sem mauris, laoreet ut, rhoncus aliquet, pulvinar sed, nisl. Nunc rhoncus dui vel sem.",
	}

	i.IndexDocument(&doc1)
	i.IndexDocument(&doc2)
	i.IndexDocument(&doc3)
	i.IndexDocument(&doc4)

	reader := bufio.NewReader(os.Stdin)
	
	for {
		fmt.Print("Enter search query: ")
		query, _ := reader.ReadString('\n')
		query = strings.Trim(query, "\n")
		matches, err := i.Query(query)

		if err != nil {
			fmt.Printf("Error occurred searching documents %s", err)
		}

		fmt.Printf("Found documents matching query: '%s'\n", query)
		s := index.PrettyDocs(matches)
		fmt.Println(s)
	}	
}