package main

import "github.com/mattbaird/elastigo/api"
import "github.com/mattbaird/elastigo/core"

func main() {

	type Address struct {
		ProvinceName string `json:"province_name"`
	}

	api.Domain = "localhost"
	//api.Port = "9300"

	core.Index("data", "address", "1", nil, Address{"ประเทศไทย"})
	core.Index("data", "address", "2", nil, Address{"เพลงไทย"})

}
