package post

import (
	_ "github.com/xcd0/cam/cl"
)

/*
↓みたいなのを出す
   G1 F2400 E293.45557
   M140 S0
   M204 S4000
   M205 X20 Y20
   M107
   G91            ;Change to relative positioning mode for filament retraction and nozzle lifting
   G1 F200 E-4    ; Retract the filament a bit before lifting the nozzle
   G1 F1000 Z5    ; Lift nozzle 5mm
   G90            ; Change to absolute positioning mode to prepare for part rermoval
   G1 X0 Y400     ; Move the print to max y pos for part rermoval
   M104 S0        ; Turn off hotend
   M106 S0        ; Turn off cooling fan
   M140 S0        ; Turn off bed
   M84            ; Disable motors
   M82            ; absolute extrusion mode
   M104 S0
*/
func (nc *Nc) GenPostProcess(ps PostSetting) {
	var post Code
	post.AddString("\n; -- post process start --\n")
	post.SetLengthUnit(nc.ClData)     // G21
	post.CleanUpAfterMove()           // G1 F200 E-3; G1 F200 E-3
	post.SetHeadTemperature(0)        // M104 S0
	post.SetCoolant(0)                // M107 S0
	post.SetBedTemperature(0)         // M140 S0
	post.DisableMotors()              // M84
	post.SetExtrusionModeToAbsolute() // G82
	post.AddString("\n; -- post process end --")
	nc.PostProcess = post
}
