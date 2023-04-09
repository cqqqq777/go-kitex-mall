package errz

const (
	Success    int64 = 0
	SuccessMsg       = "success"
)

const (
	CodeInvalidParam int64 = 10000 + iota
	CodeTokenInvalid

	CodeNoPermission
)

const (
	CodeUserService int64 = 20000 + iota

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
	ErrUserInternal = NewErrZ(WithCode(CodeUserService), WithMsg("user service busy"))

	ErrSentVerification  = NewErrZ(WithCode(CodeSentVerification), WithMsg("send verification failed"))
	ErrGetVerification   = NewErrZ(WithCode(CodeGetVerification), WithMsg("verification expired "))
	ErrWrongVerification = NewErrZ(WithCode(CodeWrongVerification), WithMsg("verification wrong"))
	ErrGenerateUid       = NewErrZ(WithCode(CodeGenerateUid), WithMsg("generate user id failed"))
	ErrUserExist         = NewErrZ(WithCode(CodeUserExist), WithMsg("user exist"))
	ErrGenerateToken     = NewErrZ(WithCode(CodeGenerateToken), WithMsg("generate user token failed"))
	ErrWrongPassword     = NewErrZ(WithCode(CodeWrongPassword), WithMsg("wrong user password"))
	ErrGetUserInfo       = NewErrZ(WithCode(CodeGetUserInfo), WithMsg("get user info failed"))
	ErrChangeAvatar      = NewErrZ(WithCode(CodeChangeAvatar), WithMsg("change avatar failed"))
	ErrChangeBackground  = NewErrZ(WithCode(CodeChangeBackground), WithMsg("change background failed"))
	ErrPublishMsgInNsq   = NewErrZ(WithCode(CodePublishMsgInNsq), WithMsg("publish message in nsq failed"))
)

const (
	CodeMerchantService int64 = 30000 + iota

	CodeGenerateMerchantId
	CodeMerchantExist
	CodeGenerateMToken
	CodeWrongPwd
	CodeGetMerchantInfo
)

var (
	ErrMerchantInternal = NewErrZ(WithCode(CodeMerchantService), WithMsg("merchant service busy"))

	ErrGenerateMerchantId = NewErrZ(WithCode(CodeGenerateMerchantId), WithMsg("generate merchant id failed"))
	ErrMerchantExist      = NewErrZ(WithCode(CodeMerchantExist), WithMsg("merchant exist"))
	ErrGenerateMToken     = NewErrZ(WithCode(CodeGenerateMToken), WithMsg("generate merchant token failed"))
	ErrWrongPwd           = NewErrZ(WithCode(CodeWrongPwd), WithMsg("wrong merchant password"))
	ErrGetMerchantInfo    = NewErrZ(WithCode(CodeGetMerchantInfo), WithMsg("get merchant info failed"))
)

const (
	CodeProductService int64 = 40000 + iota

	CodeGenerateProductId
	CodeCreateProduct
	CodeNoProduct
)

var (
	ErrProductInternal = NewErrZ(WithCode(CodeProductService), WithMsg("product service busy"))

	ErrGenerateProductId = NewErrZ(WithCode(CodeGenerateProductId), WithMsg("generate product id failed"))
	ErrCreateProduct     = NewErrZ(WithCode(CodeCreateProduct), WithMsg("create product failed"))
	ErrNoProduct         = NewErrZ(WithCode(CodeNoProduct), WithMsg("no such product"))
)

const (
	CodeOperateService int64 = 50000 + iota

	CodeGetOperateInfo
	CodeFavoriteProduct
	CodeGetFavoriteStatus
	CodeGetCommentNum
	CodeGetSaleNum
)

var (
	ErrOperateInternal = NewErrZ(WithCode(CodeOperateService), WithMsg("operate service busy"))

	ErrGetOperateInfo    = NewErrZ(WithCode(CodeGetOperateInfo), WithMsg("get operate info failed"))
	ErrFavoriteProduct   = NewErrZ(WithCode(CodeFavoriteProduct), WithMsg("favorite product failed"))
	ErrGetFavoriteStatus = NewErrZ(WithCode(CodeGetFavoriteStatus), WithMsg("get user favorite status failed"))
	ErrGetCommentNum     = NewErrZ(WithCode(CodeGetCommentNum), WithMsg("get comment num failed"))
	ErrGetSaleNum        = NewErrZ(WithCode(CodeGetSaleNum), WithMsg("get sale num failed"))
)
