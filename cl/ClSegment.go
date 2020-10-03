package cl

// 1つのセグメントは接続移動→接近移動→切削移動→離脱移動で構成される。
// CL末尾のセグメントの終了後、必ず工具原点に戻る
// 最後の原点復帰はポストが自動でつける
type ClSegment struct {
	Coolant    int         `json:"coolant"`    // 0 ~ 100 の値で指定 0の時off
	Spin       float64     `json:"spin"`       // Endmill 回転数
	Connection []ClElement `json:"connection"` // 接続移動
	Approach   []ClElement `json:"approach"`   // 接近移動
	Cut        []ClElement `json:"cut"`        // 切削移動
	Escape     []ClElement `json:"escape"`     // 離脱移動
}

func NewClSegment(
	coolant int,
	spin float64,
	connection []ClElement,
	approach []ClElement,
	cut []ClElement,
	escape []ClElement,
) *ClSegment {
	seg := new(ClSegment)
	seg.Coolant = 100
	seg.Spin = 0
	seg.Connection = connection
	seg.Approach = approach
	seg.Cut = cut
	seg.Escape = escape
	return seg
}
