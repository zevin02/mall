package e

const (
	SUCCESS       = 200
	ERROR         = 500
	InvalidParams = 400
	//user module error
	ErrorExistName     = 10001
	ErrorFailEncoding  = 10002
	ErrorExistUserName = 10003
	ErrorNotCompare    = 1004
	ErrorAuthToken     = 1005
	ErrorExpiredToken  = 1006
	ErrorUploadFail    = 1007
	ErrorSendEmail     = 1008

	//product
	ErrorUploadProduct = 2001

	//favorite
	ErrorExistFavorite = 3001
)
