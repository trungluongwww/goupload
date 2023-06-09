package response

const (
	CommonInvalidExtension = "CommonInvalidExtension"
	CommonInvalidSize      = "CommonInvalidSize"
	CommonInvalidPayload   = "CommonInvalidPayload"
	CommonInvalidOption    = "CommonInvalidOption"
	CommonFileNotFound     = "CommonFileNotFound"
	CommonSuccess          = "CommonSuccess"
	CommonBadRequest       = "CommonBadRequest"
	CommonUnAuthorization  = "CommonUnAuthorization"
)

var commonCodes = []Code{
	{
		Message: "định dạng không hợp lệ",
		Code:    0,
		Key:     "CommonInvalidExtension",
	},
	{
		Message: "kích thước không hợp lệ",
		Code:    1,
		Key:     "CommonInvalidSize",
	},
	{
		Message: "payload không hợp lệ",
		Code:    2,
		Key:     "CommonInvalidPayload",
	},
	{
		Message: "option không hợp lệ",
		Code:    3,
		Key:     "CommonInvalidOption",
	},
	{
		Message: "không tìm thấy file",
		Code:    4,
		Key:     "CommonFileNotFound",
	},
}
