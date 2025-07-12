package dto

import "gorm.io/datatypes"

type Token_info struct {
	UserID        uint
	User_no       string
	Merchant_user []string
	Permissions   []datatypes.JSON
}
