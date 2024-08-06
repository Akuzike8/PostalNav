package models

// Response structure
type Response struct {
    Message string `json:"message"`
}

// job structure
type PostalCode struct {
    Id string `json:""`
    Area string `json:""`
    Postal_code string `json:""`
    District string `json:""`
    Region string `json:""`
    Settlement string `json:""`
}

type PagenatedResponse struct {
    Data []PostalCode `json:""`
    Total int64 `json:""`
    Page int `json:""`
    LastPage float64 `json:""`
}