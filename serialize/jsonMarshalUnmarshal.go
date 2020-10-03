package serialize

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
)

func jsonMarshalUnmarshal(c *Cl) {
	cJsonByte, err := json.Marshal(c)
	if err != nil {
		log.Fatal(err)
	}
	var buf bytes.Buffer
	if err := json.Indent(&buf, cJsonByte, "", "	"); err != nil {
		log.Fatal(err)
	}

	cJson := buf.String()
	log.Println(cJson)

	outputJsonFilename := "test.json"

	if err := ioutil.WriteFile(outputJsonFilename, []byte(cJson), 0644); err != nil {
		log.Fatal(err)
	}

	var rCl Cl

	r, err := ioutil.ReadFile(outputJsonFilename)
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(r, &rCl); err != nil {
		log.Fatal(err)
	}
	log.Println(rCl)
}
