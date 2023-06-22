package constant

var ResizeName = struct {
	Size720x720   string
	Size1280x1124 string
	Size430x600   string
	Size200x200   string
	Size1280x1280 string
}{
	Size720x720:   "720x720",   // product, post, comment
	Size1280x1124: "1280x1124", // newsgroup, home, banner
	Size430x600:   "430x600",   // banner portrait
	Size200x200:   "200x200",   // logo, avatar
	Size1280x1280: "1280x1280", // identification, config
}

var (
	ListResizeName = []string{
		ResizeName.Size720x720, ResizeName.Size1280x1124, ResizeName.Size1280x1280, ResizeName.Size430x600,
		ResizeName.Size200x200,
	}
	ListResizeNameInterfaces = []interface{}{
		ResizeName.Size720x720, ResizeName.Size1280x1124, ResizeName.Size1280x1280, ResizeName.Size430x600,
		ResizeName.Size200x200,
	}
)

// extension
var (
	ListExtensionPhotoValid    = []string{"JPEG", "JPG", "PNG", "jpg", "png", "jpeg"}
	ZipExtension               = ".zip"
	ListCompressPhotoExtension = []string{"jpg", "jpeg", "JPEG", "JPG"}
)

var TypeFile = struct {
	Photo string
	File  string
}{
	Photo: "photo",
	File:  "file",
}

var PrefixDimension = struct {
	Small  string
	Medium string
}{
	Small:  "sm",
	Medium: "md",
}

const (
	Size1MB  = 1000000
	Size2MB  = 2000000
	Size5MB  = 5000000
	Size10MB = 10000000
)
