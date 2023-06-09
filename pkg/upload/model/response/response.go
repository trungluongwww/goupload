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
