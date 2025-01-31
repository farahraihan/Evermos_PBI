package handler

import (
	"evermos_pbi/internal/features/address"
	"time"
)

type AddressResponse struct {
	ID        uint      `json:"id"`
	RcpName   string    `json:"rcp_name"`
	Phone     string    `json:"phone"`
	Detail    string    `json:"detail"`
	Province  string    `json:"province"`
	Regency   string    `json:"regency"`
	District  string    `json:"district"`
	Village   string    `json:"village"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ToAddressResponse(input address.Address) AddressResponse {
	return AddressResponse{
		ID:        input.ID,
		RcpName:   input.RcpName,
		Phone:     input.Phone,
		Detail:    input.Detail,
		Province:  input.Province,
		Regency:   input.Regency,
		District:  input.District,
		Village:   input.Village,
		UserID:    input.UserID,
		CreatedAt: input.CreatedAt,
		UpdatedAt: input.UpdatedAt,
	}
}

func ToAddressResponses(address []address.Address) []AddressResponse {
	responses := make([]AddressResponse, len(address))
	for i, addres := range address {
		responses[i] = ToAddressResponse(addres)
	}
	return responses
}
