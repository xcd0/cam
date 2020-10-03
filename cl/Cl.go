package cl

type Cl struct {
	Version     string        `json:"version"`      // CLのバージョン
	LengthUnit  string        `json:"length_unit"`  // mm, inch
	ToolOrigin  []float64     `json:"tool_origin"`  // 工具原点
	CoolantType string        `json:"coolant_type"` // no, air, oil ?
	ClearanceZ  float64       `json:"clearance_z"`  // クリアランスZ
	Data        [][]ClSegment `json:"data"`         // CLデータ本体
}

func NewCl(version string) *Cl {
	cl := new(Cl)
	cl.Version = version               // CLのバージョン
	cl.LengthUnit = "mm"               // mm, inch
	cl.ToolOrigin = []float64{0, 0, 0} // 工具原点
	cl.CoolantType = "air"             // no, air, oil ?
	cl.Data = make([][]ClSegment, 1)
	return cl
}

/*

データ構造

Cl                                 // Cl構造体
* Version     string               //  CLのバージョン
* LengthUnit  string               //  mm, inch
* ToolOrigin  []float64            //  工具原点
* CoolantType string               //  no, air, oil ?
* ClearanceZ  float64              //  クリアランスZ
* Data        [][]ClSegment        //  CLデータ本体
	ClSegment                      // ClSegment構造体
	* Coolant    int               //  0 ~ 100 の値で指定 0の時off
	* Spin       float64           //  Endmill 回転数
	* Connection []ClElement       //  接続移動
	* Approach   []ClElement       //  接近移動
	* Cut        []ClElement       //  切削移動
	* Escape     []ClElement       //  離脱移動
		ClElement                  // ClElement構造体
		* Attribute ClAttribute    //  属性 このエレメントに共通なパラメータ
			ClAttribute            // ClAttribute構造体
			* CalcMode string      //  このCLを計算した演算モード
		* Points    []ClPoint      //  1点目は座標のみ信用する
			ClPoint                // ClPoint構造体
			* Mode    int          //  移動タイプ rapid, move, cw, ccw → G00, G01, G02, G03
			* Coord   []float64    //  移動終点座標       工具先端基準
			* Axis    []float64    //  工具軸方向 工具先端から工具末端への単位方向ベクトル
			* Feed    float64      //  送り速度        -1の時変更しない
			* Extrude float64      //  FDM 押し出し量  -1の時変更しない
			* Arc     *Arc         //  円弧情報 G02,G03の時のみ
				Arc                // Arc構造体
				* Center []float64 //  円弧中心座標
				* Radius float64   //  円弧半径
				* Angle  float64   //  円弧中心角

*/
