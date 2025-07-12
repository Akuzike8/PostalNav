package dto

type Bus_manager_dashboard struct {
	Routes     int64
	Drivers    int64
	Attendants int64
	Terminals  int64

	Buses struct {
		Total       int64
		Available   int64
		Maintenance int64
	}

	Bookings struct {
		Month     int64
		LastMonth int64
	}

	MonthlyTickets struct {
		Month int
		Count int64
	}

	Tickets struct {
		Year      int64
		Month     int64
		LastMonth int64
		Jan       int64
		Feb       int64
		Mar       int64
		Apr       int64
		May       int64
		Jun       int64
		Jul       int64
		Aug       int64
		Sep       int64
		Oct       int64
		Nov       int64
		Dec       int64
	}

	Sales struct {
		Year      float64
		Month     float64
		LastMonth float64
		Jan       float64
		Feb       float64
		Mar       float64
		Apr       float64
		May       float64
		Jun       float64
		Jul       float64
		Aug       float64
		Sep       float64
		Oct       float64
		Nov       float64
		Dec       float64
	}

	TopRoutes []struct {
		Origin_terminal_city      string
		Destination_terminal_city string
		Total                     string
	}

	Transactions []struct {
		Method string
		Total  string
	}

	RouteMonthlyTickets []struct {
		RouteID             int
		OriginTerminal      string
		DestinationTerminal string
		Month               int
		TicketsSold         int64
	}
}

type Management_dashboard struct {
	Users     int64
	Merchants int64

	MonthlyTickets struct {
		Month int
		Count int64
	}

	Transactions []struct {
		Method string
		Total  string
	}

	Collections struct {
		Year      float64
		Month     float64
		LastMonth float64
		Jan       float64
		Feb       float64
		Mar       float64
		Apr       float64
		May       float64
		Jun       float64
		Jul       float64
		Aug       float64
		Sep       float64
		Oct       float64
		Nov       float64
		Dec       float64
	}

	Payouts struct {
		Year      float64
		Month     float64
		LastMonth float64
		Jan       float64
		Feb       float64
		Mar       float64
		Apr       float64
		May       float64
		Jun       float64
		Jul       float64
		Aug       float64
		Sep       float64
		Oct       float64
		Nov       float64
		Dec       float64
	}

	Tickets struct {
		Year      int64
		Month     int64
		LastMonth int64
		Jan       int64
		Feb       int64
		Mar       int64
		Apr       int64
		May       int64
		Jun       int64
		Jul       int64
		Aug       int64
		Sep       int64
		Oct       int64
		Nov       int64
		Dec       int64
	}

	Sales struct {
		Year      float64
		Month     float64
		LastMonth float64
		Jan       float64
		Feb       float64
		Mar       float64
		Apr       float64
		May       float64
		Jun       float64
		Jul       float64
		Aug       float64
		Sep       float64
		Oct       float64
		Nov       float64
		Dec       float64
	}
}
