package main

import (
	"flag"
	"log"
	"path/filepath"

	"github.com/hschendel/stl"
	"github.com/xcd0/cam/engine"
	"github.com/xcd0/cam/post"
)

func main() {

	flag.Parse()
	args := flag.Args()

	var solid *stl.Solid
	if len(args) != 1 {
		log.Fatal("引数にstlファイルを１つ渡してください")

	} else {
		if e := filepath.Ext(args[0]); e == ".stl" {
			if s, err := stl.ReadFile(args[0]); err != nil {
				log.Fatal(err)
			} else {
				solid = s
			}
			//solid.Scale(25.4) // Convert from Inches to mm
			//errWrite := solid.WriteFile(outputFilename)
		}
	}
	//log.Printf("Triangles     : %v", solid.Triangles)
	log.Printf("Name          : %s", solid.Name)
	log.Printf("Triangles Num : %d", len(solid.Triangles))
	log.Printf("IsAscii       : %v", solid.IsAscii)
	outputClName := "test_output_cl.json"

	eis := engine.EngineInputSetting{0.1}
	log.Printf("ZmapPitch     : %f  // rough for test", eis.MeshPitch)
	engine.Engine(solid, &outputClName, &eis)
	return

	clFileName := "test"

	// テスト用CLを作る
	jsonClFileName := clFileName + ".json"
	engine.CalcCl(&jsonClFileName)
	post.OutputNc(&jsonClFileName)

	gobClFileName := clFileName + ".gob"
	engine.CalcCl(&gobClFileName)
	post.OutputNc(&gobClFileName)

}
