package engine

import (
	. "github.com/xcd0/cam/cl"
)

func makeTestCl() *Cl {

	// テスト用CLを作ってみる

	c := NewCl("0.0.0.0")
	a := ClAttribute{"contour"}

	// 最初の点の情報にいろいろ詰める
	initPoint := NewClPointG00(&c.ToolOrigin)
	initPoint.HeadTemperature = 230
	initPoint.BedTemperature = 80.0

	zSeg1 := 0.2
	seg1 := NewClSegment(
		100,
		0,
		[]ClElement{ClElement{a, []ClPoint{
			*initPoint,
			*NewClPointG00(&[]float64{10, 10, c.ToolOrigin[2]}), // XY con
			*NewClPointG01(&[]float64{10, 10, 10}),              // Z con
		}}},
		[]ClElement{ClElement{a, []ClPoint{
			*NewClPointG01(&[]float64{10, 10, 10}),    // Z con end
			*NewClPointG01(&[]float64{10, 10, zSeg1}), // Z app
			*NewClPointG01(&[]float64{20, 20, zSeg1}), // XY app
		}}},
		[]ClElement{ClElement{a, []ClPoint{
			*NewClPointG01(&[]float64{20, 20, zSeg1}), // XY app end
			*NewClPointG01(&[]float64{30, 20, zSeg1}),
			*NewClPointG01(&[]float64{30, 30, zSeg1}),
			*NewClPointG01(&[]float64{20, 30, zSeg1}),
			*NewClPointG01(&[]float64{20, 20, zSeg1}),
		}}},
		[]ClElement{ClElement{a, []ClPoint{
			*NewClPointG01(&[]float64{20, 20, zSeg1}),           // cut end
			*NewClPointG01(&[]float64{10, 10, zSeg1}),           // XY esc
			*NewClPointG01(&[]float64{10, 10, 20}),              // Z esc
			*NewClPointG00(&[]float64{10, 10, c.ToolOrigin[2]}), // rapid Z esc
			*NewClPointG00(&c.ToolOrigin),                       // move to origin G28
		}}},
	)

	c.Path = append(c.Path, []ClSegment{*seg1})
	return c
}
