package handler

import "evermos_pbi/internal/features/address"

type AddOrUpdateAddressRequest struct {
	RcpName  string `json:"rcp_name" form:"rcp_name"`
	Phone    string `json:"phone" form:"phone"`
	Detail   string `json:"detail" form:"detail"`
	Province string `json:"province" form:"province"`
	Regency  string `json:"regency" form:"regency"`
	District string `json:"district" form:"district"`
	Village  string `json:"village" form:"village"`
	UserID   uint   `json:"user_id" form:"user_id"`
}

func AddToAddress(ar AddOrUpdateAddressRequest) address.Address {
	return address.Address{
		RcpName:  ar.RcpName,
		Phone:    ar.Phone,
		Detail:   ar.Detail,
		Province: ar.Province,
		Regency:  ar.Regency,
		District: ar.District,
		Village:  ar.Village,
		UserID:   ar.UserID,
	}
}
