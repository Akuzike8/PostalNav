package dto

type Travel_Ticket struct {
	Source				string
	Source_code 		string
	Destination 		string
	Destination_code	string
	Duration			string
	CompanyName			string		`json:"Company"`
	Passenger			string
	Date				string
	Seat				string
	Departure			string
	Arrival				string
	Data				string
	Logo				string		`json:"Company_logo"`
	Transport_logo		string
	Validity			string
}
