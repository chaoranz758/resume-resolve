package middleware

var (
	errGetUserFromContext = "get user's information from context failed"
	errCtxValue2User      = "context don't have user's information"
	errWrongRole          = "user authority denied"
)
