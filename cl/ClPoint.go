package cl

const (
	Rapid = iota // G00
	Move         // G01
	Cw           // G02
	Ccw          // G03
)

type Arc struct {
	Center []float64 `json:"center"` // 円弧中心座標
	Radius float64   `json:"radius"` // 円弧半径
	Angle  float64   `json:"angle"`  // 円弧中心角
}

type ClPoint struct {
	Mode               int       `json:"mode"`               // 移動タイプ rapid, move, cw, ccw → G00, G01, G02, G03
	Coord              []float64 `json:"coord"`              // 移動終点座標       工具先端基準
	Axis               []float64 `json:"axis"`               // 工具軸方向 工具先端から工具末端への単位方向ベクトル
	Arc                *Arc      `json:"arc"`                // 円弧情報         G02,G03の時のみ G00,G01のときnil
	Feed               float64   `json:"feed"`               // 送り速度         -1の時変更しない
	CutterCompensation int       `json:"cutterCompensation"` // 径補正           -1の時指定なし
	Extrude            float64   `json:"extrude"`            // FDM 押し出し量   -1の時変更しない
	HeadTemperature    float64   `json:"head_temperature"`   // FDM ヘッド温度   -1の時変更しない
	BedTemperature     float64   `json:"bead_temperature"`   // FDM テーブル温度 -1の時変更しない
}

func NewClPointG00(coord *[]float64) *ClPoint {
	p := new(ClPoint)
	p.Mode = Rapid
	p.Coord = *coord
	p.Axis = []float64{0, 0, 1}
	p.Feed = -1
	p.CutterCompensation = -1
	p.Extrude = -1
	p.HeadTemperature = -1
	p.BedTemperature = -1
	return p
}

func NewClPointG01(coord *[]float64) *ClPoint {
	p := new(ClPoint)
	p.Mode = Move
	p.Coord = *coord
	p.Axis = []float64{0, 0, 1}
	p.Feed = -1
	p.Extrude = -1
	p.HeadTemperature = -1
	p.BedTemperature = -1
	return p
}

func NewClPointG02(coord *[]float64, center *[]float64, radius float64, angle float64) *ClPoint {
	p := new(ClPoint)
	p.Mode = Cw
	p.Coord = *coord
	p.Axis = []float64{0, 0, 1}
	p.Feed = -1
	p.Extrude = -1
	p.Arc = &Arc{*center, radius, angle}
	p.HeadTemperature = -1
	p.BedTemperature = -1
	return p
}

func NewClPointG03(coord *[]float64, center *[]float64, radius float64, angle float64) *ClPoint {
	p := new(ClPoint)
	p.Mode = Ccw
	p.Coord = *coord
	p.Axis = []float64{0, 0, 1}
	p.Feed = -1
	p.Extrude = -1
	p.Arc = &Arc{*center, radius, angle}
	p.HeadTemperature = -1
	p.BedTemperature = -1
	return p
}
