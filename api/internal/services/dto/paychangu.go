package dto

type PaychanguPayDepositCallback struct {
	TxrefId                  string                                     `json:"tx_ref"`
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

type PaychanguPayout struct {
	ChargeID             string `json:"charge_id"`
	Amount               int    `json:"amount"`
	Currency             string `json:"currency"`
	Country              string `json:"country"`
	StatementDescription string `json:"statement_description"`
	PaymentMethod        string `json:"payment_method"`
	Email                string `json:"email"`
	FirstName            string `json:"first_name"`
	LastName             string `json:"last_name"`
	Mobile               string `json:"mobile"`
	MobileOperator       string `json:"mobile_money_operator_ref_id"`
	TransactionStatus    string `json:"transaction_status"`
}

type PaychanguMobileMoney struct {
	Name    string `json:"name"`
	RefID   string `json:"ref_id"`
	Country string `json:"country"`
}

type PaychanguTransactionCharges struct {
	Currency string  `json:"currency"`
	Amount   float64 `json:"amount"`
}

type PaychanguAuthorization struct {
	Channel      string `json:"channel"`
	CardNumber   string `json:"card_number"`
	Expiry       string `json:"expiry"`
	Brand        string `json:"brand"`
	Provider     string `json:"provider"`
	MobileNumber string `json:"mobile_number"`
	CompletedAt  string `json:"completed_at"`
}

type Data struct {
	Amount             float64                     `json:"amount"`
	ChargeID           string                      `json:"charge_id"`
	RefID              string                      `json:"ref_id"`
	TransID            *string                     `json:"trans_id"`
	FirstName          *string                     `json:"first_name"`
	LastName           *string                     `json:"last_name"`
	Email              *string                     `json:"email"`
	Type               string                      `json:"type"`
	TraceID            *string                     `json:"trace_id"`
	Status             string                      `json:"status"`
	Mobile             string                      `json:"mobile"`
	Attempts           int                         `json:"attempts"`
	Currency           string                      `json:"currency"`
	Mode               string                      `json:"mode"`
	CreatedAt          string                      `json:"created_at"`
	CompletedAt        string                      `json:"completed_at"`
	EventType          string                      `json:"event_type"`
	MobileMoney        PaychanguMobileMoney        `json:"mobile_money"`
	TransactionCharges PaychanguTransactionCharges `json:"transaction_charges"`
	Authorization      PaychanguAuthorization      `json:"authorization"`
}

type PaychanguPayoutResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    Data   `json:"data"`
}
