package post

type MachineState struct {
	PreMoveMode           int       `json:"move_mode"`           // 0:G90 absolute, 1:G91 relative
	PreCoord              []float64 `json:"coord"`               // X,Y,Z
	PreAxis               []float64 `json:"axis"`                // 工具軸方向 多軸ポストとか考えてないけど枠だけ用意
	PreCoolant            int       `json:"coolant"`             // 0 ~ 100
	PreSpin               float64   `json:"spin"`                // S
	PreFeed               float64   `json:"feed"`                // F
	PreExtrude            float64   `json:"extrude"`             // E
	PreCutterCompensation int       `json:"cutter_compensation"` // G40, G41, G42 径補正とかとりあえずいらんけど枠だけ用意
	PreHeadTemperature    float64   `json:"head_temperature"`    // FDM ヘッド温度   -1の時変更しない
	PreBedTemperature     float64   `json:"bead_temperature"`    // FDM テーブル温度 -1の時変更しない
}
