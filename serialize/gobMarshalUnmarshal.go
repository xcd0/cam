package serialize

import (
	"bytes"
	"encoding/gob"
	"io/ioutil"
	"log"
	"os"
)

func gobMarshalUnmarshal(c *Cl) {
	output := "test.gob"
	{
		//var buf bytes.Buffer
		buf := bytes.NewBuffer(nil)
		if err := gob.NewEncoder(buf).Encode(&c); err != nil {
			log.Fatal(err)
		}

		if err := ioutil.WriteFile(output, buf.Bytes(), 0644); err != nil {
			log.Fatal(err)
		}

	}

	{
		var cl Cl
		//buf := bytes.NewBuffer(nil)
		r, err := os.Open(output)
		//r, err := ioutil.ReadFile(output)
		if err != nil {
			log.Fatal(err)
		}

		if err := gob.NewDecoder(r).Decode(&cl); err != nil {
			log.Fatal(err)
		}
		log.Println(cl)
	}
}
