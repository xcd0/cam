package engine

import (
	"math"

	"github.com/hschendel/stl"
)

type Cell struct {
	// uint64の値で3x3x3の立方体を表す 先頭5bitを何かしらの情報を持たせる
	// 5bitで32の情報持たせる 暫定
	// 0: 1ならすべて空
	// 1: 1ならすべて満
	// 2: 3: 4: 5: 6: 7:
	Bit *uint32
}
type Cells struct {
	MinCoord [3]float32
	Pitch    float32
	Data     *[][][]Cell
}

func NewCells(bb *[2]stl.Vec3, pitch float32) {

	// 計算したバウンディングボックス
	// バウンディングボックスからボクセル作る
	tmp := stl.Vec3{
		bb[0][0] - bb[1][0],
		bb[0][1] - bb[1][1],
		bb[0][2] - bb[1][2],
	}
	var mapSize [3]int
	for i, t := range tmp {
		mapSize[i] = int(math.Ceil(float64(t / pitch)))
		l := mapSize[i] % 3
		if l != 0 {
			mapSize[i] += 3 - l
		}
	}
	cells := new(Cells)
	cells.MinCoord = bb[0]
	cells.Pitch = pitch
	cells.Data = new([][][]Cell)
	for z, cc := range *cells.Data {
		cc = make([][]Cell, mapSize[2])
		for y, c := range cc {
			c = make([]Cell, mapSize[1])
		}
	}
}

func (cs *Cells) GetCell(x int, y int, z int) *Cell {
	return &(*cs.Data)[z][y][x]
}
func (cs *Cells) GetCoord(x int, y int, z int) {

}
