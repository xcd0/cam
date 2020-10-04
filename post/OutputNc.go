package post

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	. "github.com/xcd0/cam/cl"
)

func OutputNc(inputPath *string) {

	var cl *Cl
	if e := filepath.Ext(*inputPath); e == `.json` {
		cl = JsonUnmarshal(inputPath) // シリアライズテスト json
	} else if e == `.gob` {
		cl = GobUnmarshal(inputPath) // シリアライズテスト gob
	}
	fmt.Println(cl)

	var ps PostSetting
	if cl.Flavor == "marlin" {

		param := HeaderParam{
			Time:        0,
			LayerHeight: 0.2,
			MinX:        0,
			MinY:        0,
			MinZ:        0,
			MaxX:        100,
			MaxY:        100,
			MaxZ:        100,
		}

		ps = NewPostSettingForMarlin(param)
	} else {
		log.Fatalf("%s は未実装です", cl.Flavor)
	}

	nc := NewNc(cl, ps)
	gcode := nc.String()

	output := filepath.Base(*inputPath) + ".gcode"

	if err := ioutil.WriteFile(output, []byte(gcode), 0644); err != nil {
		log.Fatal(err)
	}
}
