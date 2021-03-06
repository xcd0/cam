package engine

import (
	"log"
	_ "math"
	"path/filepath"
	"time"

	"github.com/hschendel/stl"
	. "github.com/xcd0/cam/cl"
)

type EngineInputSetting struct {
	MeshPitch float32
}

type EngineSetting struct {
	InputSetting *EngineInputSetting // 入力パラメータ
	BoundingBox  [2]stl.Vec3         // 計算したバウンディングボックス
	MapSize      [3]int              // 指定ピッチでモデルが入る最小のマップサイズ
	Voxel        Cells               // [z][y][x]な感じで3x3x3のセルを並べたボクセル[0][0][0]のボクセルがバウンディングボックスの最小値を含む
}

func Engine(solid *stl.Solid, outputClPath *string, eis *EngineInputSetting) {

	var es EngineSetting
	es.InputSetting = eis
	es.BoundingBox = GetBoundingBox(solid)
	log.Println(es.BoundingBox)
	log.Println(es.MapSize)

	NewCells(&es.BoundingBox, es.InputSetting.MeshPitch)

	// stl → voxel
	StlToVoxel(solid, &es)

	time.Sleep(time.Second * 3)

	return
}

func StlToVoxel(solid *stl.Solid, es *EngineSetting) {
	// MapSizeを3で割った値分ボクセルを用意
	es.Voxel = make([][][]Cell, es.MapSize[2]/3)
	for z := 0; z < es.MapSize[2]/3; z++ {
		es.Voxel[z] = make([][]Cell, es.MapSize[1]/3)
		for y := 0; y < es.MapSize[1]/3; y++ {
			es.Voxel[z][y] = make([]Cell, es.MapSize[0]/3)
		}
	}
	//
	for i, s := range solid.Triangles {
		//
	}
}

// とりあえずテストCLを作るだけ
func CalcCl(outputClPath *string) {

	c := makeTestCl()

	if e := filepath.Ext(*outputClPath); e == ".gob" {
		GobMarshal(c, outputClPath) // シリアライズテスト gob
	} else if e == ".json" {
		JsonMarshal(c, outputClPath) // シリアライズテスト json
	}

}

func GetBoundingBox(solid *stl.Solid) [2]stl.Vec3 {
	bb := [2]stl.Vec3{}
	for _, t3 := range solid.Triangles {
		for _, t := range t3.Vertices {
			if bb[0][0] < t[0] {
				bb[0][0] = t[0]
			}
			if bb[1][0] > t[0] {
				bb[1][0] = t[0]
			}
			if bb[0][1] < t[1] {
				bb[0][1] = t[1]
			}
			if bb[1][1] > t[1] {
				bb[1][1] = t[1]
			}
			if bb[0][2] < t[2] {
				bb[0][2] = t[2]
			}
			if bb[1][2] > t[2] {
				bb[1][2] = t[2]
			}
		}
	}
	return bb
}
