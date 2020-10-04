
;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;
;FLAVOR:Marlin
;TIME:0
;Filament used:0.200000m
;Layer height:0.000000
;MINX:0.000000
;MINY:0.000000
;MINZ:100.000000
;MAXX:100.000000
;MAXY:100.000000
;MAXZ:%!f(MISSING)
;Generated with github.com/xcd0/cam
;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;

; Cam version  : 0.0.0.0
; Length unit  : mm
; Tool origin  : [0 0 0]
; Coolant type : air
; Clearance Z  : 0

; -- pre process start --

M140 S80                      ; Set bed temperature
M105                          ; Report temperature
M104 S230                     ; Set head temperature
M105                          ; Report temperature
M105                          ; Report temperature
M109 S230                     ; Set head temperature And wait until the specified value is reached
M190 S80                      ; Set bed temperature. And wait until the specified value is reached
G82                           ; Set extrusion mode to absolute
M106 S0                       ; Set coolant value (0 ~ 255)
G21                           ; Set length unit mm
G91                           ; Change to relative positioning mode for retract filament and nozzle lifting
G1 F200 E-3                   ; Retract 3mm filament for a clean start
G92 E0                        ; Zero the extruded length
G1 F1000 Z5                   ; Lift the nozzle 5mm before homing axes
G90                           ; Absolute positioning
M82                           ; Set extruder to absolute mode too
G28 X0 Y0                     ; First move X/Y to min endstops
G28 Z0                        ; Then move Z to min endstops
G1 F1000 Z15                  ; After homing lift the nozzle 15mm before start printing

; -- pre process end --

;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;

M104 S230                     ; Set head temperature
M105                          ; Report temperature
M140 S80                      ; Set bed temperature
Z 0.000000
X 10.000000 Y 10.000000
Z 10.000000

;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;

Z 0.200000
X 20.000000 Y 20.000000

;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;

X 30.000000
Y 30.000000
X 20.000000
Y 20.000000

;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;

X 10.000000 Y 10.000000
Z 20.000000
Z 0.000000
X 0.000000 Y 0.000000

;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;

; -- post process start --

G21                           ; Set length unit mm
G91                           ; Change to relative positioning mode for filament retraction and nozzle lifting
G1 F200 E-4                   ; Retract the filament a bit before lifting the nozzle
G1 F1000 Z5                   ; Lift nozzle 5mm
G90                           ; Change to absolute positioning mode to prepare for part rermoval
G1 X0 Y400                    ; Move the print to max y pos for part rermoval`)
M104 S0                       ; Set head temperature
M105                          ; Report temperature
M106 S0                       ; Set coolant value (0 ~ 255)
M140 S0                       ; Set bed temperature
M84                           ; Disable motors
G82                           ; Set extrusion mode to absolute

; -- post process end --

;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;
;End of Gcode
;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;

