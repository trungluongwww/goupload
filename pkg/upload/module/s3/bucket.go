package s3

import "github.com/trungluongwww/goupload/internal/constant"

type BucketOption struct {
	Name         string
	MaxSizePhoto int64
}

func GetBucketOption(name string) BucketOption {
	switch name {
	case constant.BucketShort.B0:
		return BucketOption{
			Name:         "public-avatar",
			MaxSizePhoto: constant.Size2MB,
		}
	case constant.BucketShort.B1:
		return BucketOption{
			Name:         "public-media",
			MaxSizePhoto: constant.Size5MB,
		}
	case constant.BucketShort.B2:
		return BucketOption{
			Name:         "public-product",
			MaxSizePhoto: constant.Size2MB,
		}
	case constant.BucketShort.B3:
		return BucketOption{
			Name:         "public-certification",
			MaxSizePhoto: constant.Size5MB,
		}
	case constant.BucketShort.B4:
		return BucketOption{
			Name:         "public-admin",
			MaxSizePhoto: constant.Size10MB,
		}
	case constant.BucketShort.B5:
		return BucketOption{
			Name:         "private-identification",
			MaxSizePhoto: constant.Size5MB,
		}
	case constant.BucketShort.B6:
		return BucketOption{
			Name:         "private-contract",
			MaxSizePhoto: constant.Size10MB,
		}
	case constant.BucketShort.B7:
		return BucketOption{
			Name:         "private-chat",
			MaxSizePhoto: constant.Size2MB,
		}
	case constant.BucketShort.B8:
		return BucketOption{
			Name:         "private-help-center",
			MaxSizePhoto: constant.Size2MB,
		}
	}

	return BucketOption{}
}
