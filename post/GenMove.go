package post

import (
	"github.com/xcd0/cam/cl"
)

func (nc *Nc) SetMove(moves *[]Code, e *cl.ClElement, coolant int, spin float64) {

	var move Code
	for _, p := range e.Points {

		// とりあえず3軸

		var diff []*float64 = make([]*float64, 9) // XYZ ABC EFS
		for i := 0; i < len(p.Coord); i++ {
			if nc.State.PreCoord[i] != p.Coord[i] {
				nc.State.PreCoord[i] = p.Coord[i] // XYZ
				diff[i] = &p.Coord[i]
			}
		}
		/* // とりあえず3軸なので
		for i := 0; i < len(p.Axis); i++ {
			if nc.State.PreAxis[i] != p.Axis[i] {
				nc.State.PreAxis[i] = p.Axis[i] // XYZ
				p[i] = p.Axis[i]
			}
		}
		*/
		if p.Extrude != -1 && nc.State.PreExtrude != p.Extrude {
			nc.State.PreExtrude = p.Extrude // E
			diff[6] = &p.Extrude
		}
		if p.Feed != -1 && nc.State.PreFeed != p.Feed {
			nc.State.PreFeed = p.Feed // F
			diff[7] = &p.Feed
		}
		if spin != -1 && nc.State.PreSpin != spin {
			nc.State.PreSpin = spin // S
			diff[8] = &spin
		}

		if p.HeadTemperature != -1 && nc.State.PreHeadTemperature != p.HeadTemperature {
			nc.State.PreHeadTemperature = p.HeadTemperature
			move.SetHeadTemperature(p.HeadTemperature)
		}
		if p.BedTemperature != -1 && nc.State.PreBedTemperature != p.BedTemperature {
			nc.State.PreBedTemperature = p.BedTemperature
			move.SetBedTemperature(p.BedTemperature)
		}

		move.MoveXYZEFS(diff)

		//nc.State.PreAxis = []float64{0, 0, 1}   // 工具軸方向 多軸ポストとか考えてないけど枠だけ用意
		if nc.State.PreCoolant != coolant {
			nc.State.PreCoolant = coolant // 0 ~ 100
			move.SetCoolant(coolant)
		}
		//nc.State.PreCutterCompensation = p.CutterCompensation // G40 G41 G42 径補正とかとりあえずいらんけど枠だけ用意
	}
	*moves = append(*moves, move)
}

func (nc *Nc) GenMove(ps PostSetting) {
	var moves []Code

	for _, segments := range nc.ClData.Path {
		for _, segment := range segments {
			for _, e := range segment.Connection {
				nc.SetMove(&moves, &e, segment.Coolant, segment.Spin)
			}
			for _, e := range segment.Approach {
				nc.SetMove(&moves, &e, segment.Coolant, segment.Spin)
			}
			for _, e := range segment.Cut {
				nc.SetMove(&moves, &e, segment.Coolant, segment.Spin)
			}
			for _, e := range segment.Escape {
				nc.SetMove(&moves, &e, segment.Coolant, segment.Spin)
			}
		}
	}

	nc.Moves = moves
}
