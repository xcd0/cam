package engine

import (
	"path/filepath"

	. "github.com/xcd0/cam/cl"
)

// とりあえずテストCLを作るだけ
func CalcCl(outputClPath *string) {

	c := makeTestCl()

	if e := filepath.Ext(*outputClPath); e == ".gob" {
		GobMarshal(c, outputClPath) // シリアライズテスト gob
	} else if e == ".json" {
		JsonMarshal(c, outputClPath) // シリアライズテスト json
	}

}
