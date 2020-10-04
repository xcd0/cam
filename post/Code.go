package post

import (
	"fmt"
	"strings"

	. "github.com/xcd0/cam/cl"
)

type Code struct {
	Lines []string
}

func (code *Code) AddString(str string) {
	code.Lines = append(code.Lines, str)
}
func (code *Code) Add(str string, comment string) {
	code.Lines = append(code.Lines, fmt.Sprintf("%-30s; %s", str, comment))
}

func (code *Code) Get() *[]string {
	return &code.Lines
}

func (code *Code) DisableMotors() {
	code.Add("M84", "Disable motors")
}
func (code *Code) ReportTemperature() {
	code.Add("M105", "Report temperature")
}
func (code *Code) SetBedTemperature(degree float64) {
	code.Add(fmt.Sprintf("M140 S%d", int(degree)), "Set bed temperature")
}
func (code *Code) SetBedTemperatureWithWait(degree float64) {
	code.Add(fmt.Sprintf("M190 S%d", int(degree)), "Set bed temperature. And wait until the specified value is reached")
}
func (code *Code) SetHeadTemperature(degree float64) {
	code.Add(fmt.Sprintf("M104 S%d", int(degree)), "Set head temperature")
	code.Add("M105", "Report temperature")
}
func (code *Code) SetHeadTemperatureWithWait(degree float64) {
	code.Add(fmt.Sprintf("M109 S%d", int(degree)), "Set head temperature And wait until the specified value is reached")
}
func (code *Code) SetCoolant(value int) {
	if value > 100 {
		value = 100
	}
	if value < 0 {
		value = 0
	}
	code.Add(fmt.Sprintf("M106 S%d", int(float64(value)/100.0*255)), "Set coolant value (0 ~ 255)")
}
func (code *Code) SetLengthUnit(cl *Cl) {
	if cl.LengthUnit == "mm" {
		code.Add("G21", "Set length unit mm")
	} else if cl.LengthUnit == "inch" {
		code.Add("G20", "Set length unit inch")
	}
}
func (code *Code) ReturnToOrigin() {
	code.Add("M28", "Return to the origin")
}
func (code *Code) WaitSeconds(sec int) {
	code.Add(fmt.Sprintf("M04 S%s", sec), "wait seconds")
}
func (code *Code) WaitMilliSeconds(ms int) {
	code.Add(fmt.Sprintf("M04 P%s", ms), "wait milli seconds")
}
func (code *Code) SetExtrusionModeToAbsolute() {
	code.Add("G82", "Set extrusion mode to absolute")
}

func (code *Code) MoveXYZEFS(p []*float64) {
	var a []string
	for i, v := range p {
		if v != nil {
			switch i {
			case 0:
				a = append(a, fmt.Sprintf("X %f", *v))
			case 1:
				a = append(a, fmt.Sprintf("Y %f", *v))
			case 2:
				a = append(a, fmt.Sprintf("Z %f", *v))
			case 3:
				a = append(a, fmt.Sprintf("A %f", *v))
			case 4:
				a = append(a, fmt.Sprintf("B %f", *v))
			case 5:
				a = append(a, fmt.Sprintf("C %f", *v))

			case 6:
				a = append(a, fmt.Sprintf("E %f", *v))
			case 7:
				a = append(a, fmt.Sprintf("F %f", *v))
			case 8:
				a = append(a, fmt.Sprintf("S %f", *v))
			}
		}
	}
	code.AddString(strings.Join(a, " "))
}

func (code *Code) SetMoveModeToAbsolute() {
	code.Add("G90", "Set move mode to absolute")
}
func (code *Code) SetMoveModeToRelative() {
	code.Add("G91", "Set move mode to relative")
}
func (code *Code) PreparateionForMove() {
	code.Add("G91", "Change to relative positioning mode for retract filament and nozzle lifting")
	code.Add("G1 F200 E-3", "Retract 3mm filament for a clean start")
	code.Add("G92 E0", "Zero the extruded length")
	code.Add("G1 F1000 Z5", "Lift the nozzle 5mm before homing axes")
	code.Add("G90", "Absolute positioning")
	code.Add("M82", "Set extruder to absolute mode too")
	code.Add("G28 X0 Y0", "First move X/Y to min endstops")
	code.Add("G28 Z0", "Then move Z to min endstops")
	code.Add("G1 F1000 Z15", "After homing lift the nozzle 15mm before start printing")
}
func (code *Code) CleanUpAfterMove() {
	code.Add("G91", "Change to relative positioning mode for filament retraction and nozzle lifting")
	code.Add("G1 F200 E-4", "Retract the filament a bit before lifting the nozzle")
	code.Add("G1 F1000 Z5", "Lift nozzle 5mm")
	code.Add("G90", "Change to absolute positioning mode to prepare for part rermoval")
	code.Add("G1 X0 Y400", "Move the print to max y pos for part rermoval`)")
}

func (code *Code) AddComment(str string) {
	code.AddString(fmt.Sprintf("; %s", str))
}
