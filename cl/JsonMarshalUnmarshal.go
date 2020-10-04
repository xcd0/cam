package cl

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
)

func JsonMarshal(c *Cl, outputPath *string) {
	output := *outputPath
	cJsonByte, err := json.Marshal(c)
	if err != nil {
		log.Fatal(err)
	}
	var buf bytes.Buffer
	if err := json.Indent(&buf, cJsonByte, "", "	"); err != nil {
		log.Fatal(err)
	}
	cJson := buf.String()
	//log.Println(cJson)
	if err := ioutil.WriteFile(output, []byte(cJson), 0644); err != nil {
		log.Fatal(err)
	}
}

func JsonUnmarshal(outputPath *string) *Cl {
	output := *outputPath
	cl := new(Cl)
	r, err := ioutil.ReadFile(output)
	if err != nil {
		log.Fatal(err)
	}
	if err := json.Unmarshal(r, cl); err != nil {
		log.Fatal(err)
	}
	//log.Println(cl)
	return cl
}

func JsonMarshalUnmarshal(c *Cl, outputPath *string) {
	JsonMarshal(c, outputPath)
	_ = JsonUnmarshal(outputPath)
	//cl := JsonUnmarshal(outputPath)
	//log.Println(*cl)
}
