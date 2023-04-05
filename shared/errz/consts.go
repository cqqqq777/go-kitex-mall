package errz

const (
	Success    int64 = 0
	SuccessMsg       = "success"
)

const (
	CodeService int64 = 20000 + iota

	CodeSentVerification
	CodeGetVerification
	CodeWrongVerification
	CodeGenerateUid
	CodeUserExist
	CodeGenerateToken
	CodeWrongPassword
	CodeGetUserInfo
	CodeChangeAvatar
	CodeChangeBackground
	CodePublishMsgInNsq
)

var (
	ErrInternal = NewErrZ(WithCode(CodeService), WithMsg("user service busy"))

	ErrSentVerification  = NewErrZ(WithCode(CodeSentVerification), WithMsg("send verification failed"))
	ErrGetVerification   = NewErrZ(WithCode(CodeGetVerification), WithMsg("verification expired "))
	ErrWrongVerification = NewErrZ(WithCode(CodeWrongVerification), WithMsg("verification wrong"))
	ErrGenerateUid       = NewErrZ(WithCode(CodeGenerateUid), WithMsg("generate user id failed"))
	ErrUserExist         = NewErrZ(WithCode(CodeUserExist), WithMsg("user exist"))
	ErrGenerateToken     = NewErrZ(WithCode(CodeGenerateToken), WithMsg("generate token failed"))
	ErrWrongPassword     = NewErrZ(WithCode(CodeWrongPassword), WithMsg("wrong password"))
	ErrGetUserInfo       = NewErrZ(WithCode(CodeGetUserInfo), WithMsg("get user info failed"))
	ErrChangeAvatar      = NewErrZ(WithCode(CodeChangeAvatar), WithMsg("change avatar failed"))
	ErrChangeBackground  = NewErrZ(WithCode(CodeChangeBackground), WithMsg("change background failed"))
	ErrPublishMsgInNsq   = NewErrZ(WithCode(CodePublishMsgInNsq), WithMsg("publish message in nsq failed"))
)
