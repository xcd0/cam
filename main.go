package main

import (
	//. "github.com/xcd0/cam/cl"
	"github.com/xcd0/cam/engine"
	"github.com/xcd0/cam/post"
)

func main() {

	clFileName := "test"

	// テスト用CLを作る
	jsonClFileName := clFileName + ".json"
	engine.CalcCl(&jsonClFileName)
	post.OutputNc(&jsonClFileName)

	gobClFileName := clFileName + ".gob"
	engine.CalcCl(&gobClFileName)
	post.OutputNc(&gobClFileName)

}
