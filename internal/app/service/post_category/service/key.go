package service

const (
	KeyPrefix                 = "resume_resolving:post_category:"
	KeyLevel1PostCategoryZSet = "level1_zSet"
	KeyLevel2PostCategoryZSet = "level2_zSet:"
	KeyString                 = "string:"
)

func GetKey(part string) string {
	return KeyPrefix + part
}
