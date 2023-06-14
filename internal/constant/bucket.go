package constant

var BucketShort = struct {
	B0 string
	B1 string
	B2 string
	B3 string
	B4 string
	B5 string
	B6 string
	B7 string
	B8 string
}{
	B0: "b00",
	B1: "b01",
	B2: "b02",
	B3: "b03",
	B4: "b04",
	B5: "b05",
	B6: "b06",
	B7: "b07",
	B8: "b08",
}

var (
	ListBucketShort = []string{
		BucketShort.B0, BucketShort.B1, BucketShort.B2, BucketShort.B3, BucketShort.B4, BucketShort.B5, BucketShort.B6,
		BucketShort.B7, BucketShort.B8,
	}
	BucketShortInterfaces = []interface{}{
		BucketShort.B0, BucketShort.B1, BucketShort.B2, BucketShort.B3, BucketShort.B4, BucketShort.B5, BucketShort.B6,
		BucketShort.B7, BucketShort.B8,
	}
)
