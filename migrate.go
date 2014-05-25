package main

import "github.com/mattbaird/elastigo/api"
import "github.com/mattbaird/elastigo/core"
import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
)

func main() {

	api.Domain = "localhost"

	db, err := sql.Open("mysql", "user1:user1password@/todo")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	sql := `
	SELECT p1.province_id, p1.province_code, p1.province_name
		,a1.amphur_id, a1.amphur_code, a1.amphur_name
		,d1.district_id, d1.district_code, d1.district_name
		, zipcode
	FROM  provinces p1, amphures a1
	    , districts d1, zipcodes z1
	WHERE p1.province_id = a1.province_id
	  AND a1.amphur_id = d1.amphur_id
	  AND z1.district_code = d1.DISTRICT_CODE
	`

	statementQuery, err := db.Prepare(sql)
	if err != nil {
		panic(err.Error())
	}
	defer statementQuery.Close()

	rows, err := statementQuery.Query()
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	type Address struct {
		ProvinceID   int    `json:"province_id"`
		ProvinceCode string `json:"province_code"`
		ProvinceName string `json:"province_name"`
		AmphurID     int    `json:"amphur_id"`
		AmphurCode   string `json:"amphur_code"`
		AmphurName   string `json:"amphur_name"`
		DistrictID   int    `json:"district_id"`
		DistrictCode string `json:"district_code"`
		DistrictName string `json:"district_name"`
		ZipCode      string `json:"zipcode"`
	}
	index := 1
	for rows.Next() {
		var province_id int
		var province_code string
		var province_name string
		var amphur_id int
		var amphur_code string
		var amphur_name string
		var district_id int
		var district_code string
		var district_name string
		var zipcode string
		rows.Scan(&province_id, &province_code, &province_name,
			&amphur_id, &amphur_code, &amphur_name,
			&district_id, &district_code, &district_name, &zipcode)

		core.Index("data", "address", strconv.Itoa(index), nil,
			Address{province_id, province_code, province_name,
				amphur_id, amphur_code, amphur_name,
				district_id, district_code, district_name, zipcode})
		index++
	}

}
