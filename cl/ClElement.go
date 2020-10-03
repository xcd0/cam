package cl

type ClElement struct {
	Attribute ClAttribute `json:"attribute"` // 属性 このエレメントに共通なパラメータ
	Points    []ClPoint   `json:"points"`    // 1点目は座標のみ信用する
}
