package dto

type PayChangu struct{}
type CreditCard struct{}

type CollectionTransaction struct{
	Trip_no			string `json:"trip_no"`
	Amount          string `json:"amount"`
	Currency        string `json:"currency"`
	Method          string `json:"method"`
	Booking_no      string `json:"booking_no"`
	Status			string `json:"status"`
	Transaction_id  string `jsom:"transaction_id"`
	Customer        string `json:"customer"`
	Route			string `json:"route"`
	Date			string `json:"date"`
}

type PayoutTransaction struct{
	Trip_no			string `json:"trip_no"`
	Amount          string `json:"amount"`
	Currency        string `json:"currency"`
	Method          string `json:"method"`
	Booking_no      string `json:"booking_no"`
	Status			string `json:"status"`
	Transaction_id  string `jsom:"transaction_id"`
	Customer        string `json:"customer"`
	Route			string `json:"route"`
	Date			string `json:"date"`
}

type ManagementTransaction struct{
	Payment_type	string `json:"payment_type"`
	Amount          string `json:"amount"`
	Currency        string `json:"currency"`
	Method          string `json:"method"`
	Booking_no      string `json:"booking_no"`
	Status			string `json:"status"`
	Transaction_id  string `jsom:"transaction_id"`
	Customer        string `json:"customer"`
	Merchant		string `json:"merchant"`
	Date			string `json:"date"`
}

type Payment struct {
	Route       	string `json:"route"`
	Amount      	string `json:"amount"`
	User_no     	string `json:"user_no"`
	Merchants_no 	string `json:"merchants_no"`
	Booking_sno 	string `json:"booking_sno"`
	Booking_no  	string `json:"booking_no"`
	Payment     	string `json:"payment"`
	Category    	string `json:"category"`
	Event			string `json:"event"`

}

type VerifyPayment struct {
	Payment 	string `json:"payment"`
	BookingId	string `json:"booking_id"`
	DepositId   string `json:"deposit_id"`
	TxRef		string `json:"tx_ref"`
	Status		string `json:"status"`
}
