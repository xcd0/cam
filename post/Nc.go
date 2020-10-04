package post

import (
	"github.com/xcd0/cam/cl"
)

type Nc struct {
	Header      string
	PreProcess  Code
	Moves       []Code
	PostProcess Code
	Footer      string
	ClData      *cl.Cl
	State       MachineState
}

func NewNc(cl *cl.Cl, ps PostSetting) *Nc {
	nc := new(Nc)
	nc.ClData = cl
	nc.GenHeader(ps)
	nc.GenPreProcess(ps)
	nc.GenMove(ps)
	nc.GenPostProcess(ps)
	nc.GenFooter(ps)
	return nc
}

func (nc *Nc) String() string {
	gcode := ""

	// ヘッダー
	gcode += nc.Header + "\n"

	// 初期設定 温度とか
	for _, s := range nc.PreProcess.Lines {
		gcode += s + "\n"
	}
	gcode += "\n;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;\n\n"

	// 移動
	for _, m := range nc.Moves {
		for _, s := range m.Lines {
			gcode += s + "\n"
		}
		gcode += "\n;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;\n"
	}

	// 終了処理
	for _, s := range nc.PostProcess.Lines {
		gcode += s + "\n"
	}

	// フッター
	gcode += nc.Footer + "\n"
	return gcode
}

func (nc *Nc) GenHeader(ps PostSetting) { nc.Header = ps.GetHeader() }
func (nc *Nc) GenFooter(ps PostSetting) { nc.Footer = ps.GetFooter() }
