package service

const (
	KeyPrefix = "resume_resolving:department:"
	KeyZSet   = "zSet:"
	KeyString = "string:"
)

func GetKey(part string) string {
	return KeyPrefix + part
}
