package responsemodel

type PhotoResponse struct {
	ID         string         `json:"id"`
	Name       string         `json:"name"`
	Dimensions DimensionPhoto `json:"dimensions"`
}

type DimensionPhoto struct {
	Small  SizePhoto `json:"sm"`
	Medium SizePhoto `json:"md"`
}

type SizePhoto struct {
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Url    string `json:"url"`
}

type FileResponse struct {
	ID         string `json:"_id"`
	Name       string `json:"name"`
	NameOrigin string `json:"nameOrigin"`
	Url        string `json:"url"`
}
