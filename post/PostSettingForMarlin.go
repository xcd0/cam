package post

import "fmt"

// CNCコントローラーごとに違いがあるので
// マシン毎にいろいろ適応したいが
// 実機使えるのが3Dプリンターくらいなので
// とりあえずMarlin向けのポスト設定を書いてみる
// 一通りできたらFunuc向けに書いてみる

type PostSettingForMarlin struct {
	//PostSetting
	Header string
	Footer string
}

func (ps *PostSettingForMarlin) GetHeader() string { return ps.Header }
func (ps *PostSettingForMarlin) SetHeader(param HeaderParam) {
	ps.Header = fmt.Sprintf(`
;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;
;FLAVOR:Marlin
;TIME:%d
;Filament used:%fm
;Layer height:%f
;MINX:%f
;MINY:%f
;MINZ:%f
;MAXX:%f
;MAXY:%f
;MAXZ:%f
;Generated with github.com/xcd0/cam
;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;
`, param.Time, param.LayerHeight, param.MinX, param.MinY, param.MinZ, param.MaxX, param.MaxY, param.MaxZ)
}

func (ps *PostSettingForMarlin) GetFooter() string { return ps.Footer }
func (ps *PostSettingForMarlin) SetFooter() {
	ps.Footer = `
;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;
;End of Gcode
;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;
`
}

func NewPostSettingForMarlin(param HeaderParam) PostSetting {
	var psfm PostSettingForMarlin
	psfm.SetHeader(param)
	psfm.SetFooter()
	return &psfm
}
