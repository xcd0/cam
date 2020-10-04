package cl

import (
	"bytes"
	"encoding/gob"
	"io/ioutil"
	"log"
	"os"
)

func GobMarshal(c *Cl, outputPath *string) {
	output := *outputPath
	buf := bytes.NewBuffer(nil)
	if err := gob.NewEncoder(buf).Encode(&c); err != nil {
		log.Fatal(err)
	}
	if err := ioutil.WriteFile(output, buf.Bytes(), 0644); err != nil {
		log.Fatal(err)
	}
}

func GobUnmarshal(outputPath *string) *Cl {
	cl := new(Cl)
	r, err := os.Open(*outputPath)
	if err != nil {
		log.Fatal(err)
	}
	if err := gob.NewDecoder(r).Decode(cl); err != nil {
		log.Fatal(err)
	}
	return cl
}

func GobMarshalUnmarshal(c *Cl, outputPath *string) {
	GobMarshal(c, outputPath)
	_ = GobUnmarshal(outputPath)
	//cl := GobUnmarshal(outputPath)
	//log.Println(cl)
}
