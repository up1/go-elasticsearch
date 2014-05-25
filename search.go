package main

import "github.com/mattbaird/elastigo/api"
import "github.com/mattbaird/elastigo/core"
import "fmt"

func main() {

	type Address struct {
		ProvinceName string `json:"province_name"`
	}

	api.Domain = "localhost"

	searchJson := `{"query":{"match":{"province_name":"ไทย"}}}`

	out, _ := core.SearchRequest("data", "address", nil, searchJson)
	for i:= 0; i<len(out.Hits.Hits); i++ {
		fmt.Printf("%s\n", out.Hits.Hits[i].Source)
	}

}
