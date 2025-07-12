package dto

// Structs for PawaPay Callback

type PawapayDepositCallbackAddress struct {
	Value string `json:"value"`
}

type PawapayDepositCallbackPayer struct {
	Type    string                        `json:"type"`
	Address PawapayDepositCallbackAddress `json:"address"`
}

type PawapayDepositCallbackCorrespondentIds struct {
	MTNInit  string `json:"MTN_INIT"`
	MTNFinal string `json:"MTN_FINAL"`
}

type PawapayDepositCallbackSuspiciousActivity struct {
	ActivityType string `json:"activityType"`
	Comment      string `json:"comment"`
}

type PawapayDepositCallbackFailureReason struct {
	FailureCode    string `json:"failureCode"`
	FailureMessage string `json:"failureMessage"`
}

type PawapayDepositCallbackMetadata struct {
	Booking_no string `json:"booking_no"`
	User_no    string `json:"user_no"`
}

type PawaPayDepositCallback struct {
	DepositId                string                                     `json:"depositId"`
	Status                   string                                     `json:"status"`
	RequestedAmount          string                                     `json:"requestedAmount"`
	Currency                 string                                     `json:"currency"`
	Country                  string                                     `json:"country"`
	Correspondent            string                                     `json:"correspondent"`
	Payer                    PawapayDepositCallbackPayer                `json:"payer"`
	CustomerTimestamp        string                                     `json:"customerTimestamp"`
	StatementDescription     string                                     `json:"statementDescription"`
	Created                  string                                     `json:"created"`
	DepositedAmount          string                                     `json:"depositedAmount"`
	RespondedByPayer         string                                     `json:"respondedByPayer"`
	CorrespondentIds         PawapayDepositCallbackCorrespondentIds     `json:"correspondentIds"`
	SuspiciousActivityReport []PawapayDepositCallbackSuspiciousActivity `json:"suspiciousActivityReport"`
	FailureReason            PawapayDepositCallbackFailureReason        `json:"failureReason"`
	Metadata                 PawapayDepositCallbackMetadata             `json:"metadata"`
}

type PawapayAddress struct {
	Value string `json:"value"`
}

type PawapayCorrespondentIDs struct {
	SomeCorrespondentID string `json:"SOME_CORRESPONDENT_ID"`
}

type PawapayMeta struct {
	FieldName  string `json:"fieldName"`
	FieldValue string `json:"fieldValue"`
	IsPII      bool   `json:"isPII"`
}

type PawapayRecipient struct {
	Type    string         `json:"type"`
	Address PawapayAddress `json:"address"`
}

type PawapayDeposit struct {
	DepositId            string        `json:"depositId"`
	ReturnUrl            string        `json:"returnUrl"`
	StatementDescription string        `json:"statementDescription"`
	Amount               string        `json:"amount"`
	Country              string        `json:"country"`
	Language             string        `json:"language"`
	Reason               string        `json:"reason"`
	Metadata             []PawapayMeta `json:"metadata"`
}

type PawapayPayout struct {
	PayoutID             string           `json:"payoutId"`
	Amount               string           `json:"amount"`
	Currency             string           `json:"currency"`
	Country              string           `json:"country"`
	Correspondent        string           `json:"correspondent"`
	Recipient            PawapayRecipient `json:"recipient"`
	CustomerTimestamp    string           `json:"customerTimestamp"`
	StatementDescription string           `json:"statementDescription"`
	Metadata             []PawapayMeta    `json:"metadata"`
}

type PawapayPayoutResponse struct {
	PayoutID string `json:"payoutId"`
	Status   string `json:"status"`
	Created  string `json:"created"`
}

type PawapayPayoutStatus struct {
	PayoutID             string                  `json:"payoutId"`
	Status               string                  `json:"status"`
	Amount               string                  `json:"amount"`
	Currency             string                  `json:"currency"`
	Country              string                  `json:"country"`
	Correspondent        string                  `json:"correspondent"`
	Recipient            PawapayRecipient        `json:"recipient"`
	CustomerTimestamp    string                  `json:"customerTimestamp"`
	StatementDescription string                  `json:"statementDescription"`
	Created              string                  `json:"created"`
	ReceivedByRecipient  string                  `json:"receivedByRecipient"`
	CorrespondentIDs     PawapayCorrespondentIDs `json:"correspondentIds"`
	Metadata             PawapayMeta             `json:"metadata"`
}
