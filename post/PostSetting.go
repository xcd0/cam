package post

type HeaderParam struct {
	Time        int
	LayerHeight float64
	MinX        float64
	MinY        float64
	MinZ        float64
	MaxX        float64
	MaxY        float64
	MaxZ        float64
}

type PostSetting interface {
	SetHeader(HeaderParam)
	GetHeader() string
	SetFooter()
	GetFooter() string
}
