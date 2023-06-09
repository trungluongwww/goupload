package requestmodel

type FileInfoPayload struct {
	OriginName string `json:"originName"`
	Path       string `json:"path"`
	Name       string `json:"name"`
	Size       int64  `json:"size"`
	Width      int64  `json:"width"`
	Height     int64  `json:"height"`
	Ext        string `json:"ext"`
}
