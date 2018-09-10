package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

type Version struct {
	Version string `xml:"version"`
}

func main() {

	dat, err := ioutil.ReadFile("./pom.xml")
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	v := &Version{}
	// fmt.Print(string(dat))
	err = xml.Unmarshal(dat, &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	fmt.Println(v.Version)
}
