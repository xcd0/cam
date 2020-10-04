package post

import (
	"fmt"

	_ "github.com/xcd0/cam/cl"
)

/*
↓みたいなのまねて生成する
M140 S80
M105
M190 S80
M104 S230
M105
M109 S230
M82               ; absolute extrusion mode
M107              ; Start with the fan off
G21               ; Set units to millimeters
G91               ; Change to relative positioning mode for retract filament and nozzle lifting
G1 F200 E-3       ; Retract 3mm filament for a clean start
G92 E0            ; Zero the extruded length
G1 F1000 Z5       ; Lift the nozzle 5mm before homing axes
G90               ; Absolute positioning
M82               ; Set extruder to absolute mode too
G28 X0 Y0         ; First move X/Y to min endstops
G28 Z0            ; Then move Z to min endstops
G1 F1000 Z15      ; After homing lift the nozzle 15mm before start printing

*/

func (nc *Nc) GenPreProcess(ps PostSetting) {
	var pre Code

	pre.AddComment("Cam version  : " + nc.ClData.Version)                       // camのバージョン
	pre.AddComment("Length unit  : " + nc.ClData.LengthUnit)                    // mm, inch
	pre.AddComment("Tool origin  : " + fmt.Sprintf("%v", nc.ClData.ToolOrigin)) // 工具原点
	pre.AddComment("Coolant type : " + nc.ClData.CoolantType)                   // no, air, oil ?
	pre.AddComment("Clearance Z  : " + fmt.Sprintf("%v", nc.ClData.ClearanceZ)) // クリアランスZ
	pre.AddString("")

	init := &nc.ClData.Path[0][0].Connection[0].Points[0]
	pre.AddComment("-- pre process start --\n")
	pre.SetBedTemperature(init.BedTemperature)           // M140 S%d
	pre.ReportTemperature()                              // M105
	pre.SetHeadTemperature(init.HeadTemperature)         // M104 S%d
	pre.ReportTemperature()                              // M105
	pre.SetHeadTemperatureWithWait(init.HeadTemperature) // M190 S%d
	pre.SetBedTemperatureWithWait(init.BedTemperature)   // M109 S%d
	pre.SetExtrusionModeToAbsolute()                     // G82
	pre.SetCoolant(0)                                    // M107 S0
	pre.SetLengthUnit(nc.ClData)                         // G21
	pre.PreparateionForMove()                            // G1 F200 E-3; G1 F200 E-3
	pre.AddString("")
	pre.AddComment("-- pre process end --")
	nc.PreProcess = pre

	nc.State = MachineState{
		PreMoveMode:           0,                   // 0:G90 absolute, 1:G91 relative
		PreCoord:              []float64{0, 0, 15}, // X,Y,Z
		PreAxis:               []float64{0, 0, 1},  // 工具軸方向 多軸ポストとか考えてないけど枠だけ用意
		PreCoolant:            100,                 // 0 ~ 100
		PreSpin:               0.,                  // S
		PreFeed:               1000.,               // F
		PreExtrude:            0.,                  // E
		PreCutterCompensation: 0,                   // G40, G41, G42 径補正とかとりあえずいらんけど枠だけ用意
	}

}
