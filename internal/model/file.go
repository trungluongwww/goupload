package model

import "time"

// FileRaw ...
type FileRaw struct {
	ID           AppID              `bson:"_id"`
	Name         string             `bson:"name"`
	Ext          string             `bson:"ext"`
	OriginalName string             `bson:"originalName"`
	Dimension    DimensionFilePhoto `bson:"dimension"`
	CreatedAt    time.Time          `bson:"createdAt"`
	Type         string             `bson:"type"`
	UpdatedAt    time.Time          `bson:"updatedAt"`
}

// DimensionFilePhoto ...
type DimensionFilePhoto struct {
	Original FileSize `bson:"original"`
	Small    FileSize `bson:"small"`
	Medium   FileSize `bson:"medium"`
}

// FileSize ...
type FileSize struct {
	Width  int `bson:"width"`
	Height int `bson:"height"`
}

// ResizeDimension ...
type ResizeDimension struct {
	MediumWidth  int
	MediumHeight int
	SmallWidth   int
	SmallHeight  int
}
