package rest

type establishMode string

var (
	SecuredMode establishMode = "secured"
	UnsafeMode  establishMode = "unsafe"
)
