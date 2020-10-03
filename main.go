package main

import (
	. "github.com/xcd0/cam/cl"
	. "github.com/xcd0/cam/serialize"
)

func main() {

	c := makeTestCl()
	//gobMarshalUnmarshal(c)  // シリアライズテスト gob
	//jsonMarshalUnmarshal(c) // シリアライズテスト json

}
