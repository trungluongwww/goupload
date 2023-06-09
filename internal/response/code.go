package response

type Code struct {
	Message string
	Code    int
	Key     string
}

var (
	list        []Code
	defaultCode = Code{
		Message: "không tìm thấy",
		Code:    -1,
		Key:     "",
	}
)

func Init() {
	list = append(list, commonCodes...)
}

func getCodeByKey(key string) Code {
	for _, item := range list {
		if item.Key == key {
			return item
		}
	}

	return defaultCode
}
