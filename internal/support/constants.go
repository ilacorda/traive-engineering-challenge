package support

const (
	MsgTransactionCreatedSuccessfully    = "Transaction created successfully"
	ErrFailedToRetrieveTransactions      = "Failed to retrieve transactions"
	ErrFailedToEncodeResponse            = "Failed to encode response"
	ErrFailedToDecodeRequest             = "Failed to decode request body"
	ErrFailedToCreateTransaction         = "Failed to create transaction"
	ErrFailedToMarshalRequestBody        = "Failed to marshal request body: %v"
	ErrFailedToMarshalExpectedResponse   = "Failed to marshal expected response for %s: %v"
	ErrFailedToUnmarshalExpectedResponse = "Failed to unmarshal expected response JSON for %s: %v"
	DesktopWeb                           = "desktop-web"
	MobileAndroid                        = "mobile-android"
	MobileIOS                            = "mobile-ios"
	DatabaseURL                          = "DATABASE_URL"
)
