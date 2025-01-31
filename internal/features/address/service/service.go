package service

import (
	"context"
	"encoding/json"
	"errors"
	"evermos_pbi/internal/features/address"
	"fmt"
	"log"
	"net/http"
	"time"
)

type AddressServices struct {
	qry address.AQuery
}

func NewAdreessService(q address.AQuery) address.AService {
	return &AddressServices{
		qry: q,
	}
}

func (as *AddressServices) AddAddress(newAddress address.Address) error {
	err := as.qry.AddAddress(newAddress)
	if err != nil {
		log.Println("add address query error: ", err)
		return errors.New("failed to add new address, please try again later")
	}

	return nil
}

func (as *AddressServices) UpdateAddress(userID uint, addressID uint, updateAddress address.Address) error {
	err := as.qry.UpdateAddress(userID, addressID, updateAddress)
	if err != nil {
		log.Println("update address query error: ", err)
		return errors.New("failed to update address, please try again later")
	}

	return nil
}

func (as *AddressServices) DeleteAddress(userID uint, addressID uint) error {
	err := as.qry.DeleteAddress(userID, addressID)
	if err != nil {
		log.Println("delete address query error: ", err)
		return errors.New("failed to delete address, please try again later")
	}

	return nil
}

func (as *AddressServices) GetAddressByUserID(userID uint) ([]address.Address, error) {
	address, err := as.qry.GetAddressByUserID(userID)

	if err != nil {
		log.Println("get address by user ID query error: ", err)
		return nil, errors.New("failed to retrieve address, please try again later")
	}

	return address, nil
}

func (as *AddressServices) GetProvince() ([]map[string]interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "https://www.emsifa.com/api-wilayah-indonesia/api/provinces.json", nil)
	if err != nil {
		log.Println("Error creating request:", err)
		return nil, fmt.Errorf("failed to create request")
	}

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Error fetching data from external API:", err)
		return nil, fmt.Errorf("failed to fetch provinces")
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		log.Println("Unexpected HTTP status:", response.Status)
		return nil, fmt.Errorf("unexpected response status: %d", response.StatusCode)
	}

	var provinces []map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&provinces)
	if err != nil {
		log.Println("Error decoding JSON:", err)
		return nil, fmt.Errorf("failed to decode response")
	}

	return provinces, nil
}

func (as *AddressServices) GetRegency(provinceID uint) ([]map[string]interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	url := fmt.Sprintf("https://www.emsifa.com/api-wilayah-indonesia/api/regencies/%d.json", provinceID)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		log.Println("Error creating request:", err)
		return nil, fmt.Errorf("failed to create request")
	}

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Error fetching data from external API:", err)
		return nil, fmt.Errorf("failed to fetch regencies")
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		log.Println("Unexpected HTTP status:", response.Status)
		return nil, fmt.Errorf("unexpected response status: %d", response.StatusCode)
	}

	var regencies []map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&regencies)
	if err != nil {
		log.Println("Error decoding JSON:", err)
		return nil, fmt.Errorf("failed to decode response")
	}

	return regencies, nil
}

func (as *AddressServices) GetDistrict(regenciesID uint) ([]map[string]interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	url := fmt.Sprintf("https://www.emsifa.com/api-wilayah-indonesia/api/districts/%d.json", regenciesID)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		log.Println("Error creating request:", err)
		return nil, fmt.Errorf("failed to create request")
	}

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Error fetching data from external API:", err)
		return nil, fmt.Errorf("failed to fetch districts")
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		log.Println("Unexpected HTTP status:", response.Status)
		return nil, fmt.Errorf("unexpected response status: %d", response.StatusCode)
	}

	var districts []map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&districts)
	if err != nil {
		log.Println("Error decoding JSON:", err)
		return nil, fmt.Errorf("failed to decode response")
	}

	return districts, nil
}

func (as *AddressServices) GetVillage(districtsID uint) ([]map[string]interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	url := fmt.Sprintf("https://www.emsifa.com/api-wilayah-indonesia/api/villages/%d.json", districtsID)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		log.Println("Error creating request:", err)
		return nil, fmt.Errorf("failed to create request")
	}

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Error fetching data from external API:", err)
		return nil, fmt.Errorf("failed to fetch villages")
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		log.Println("Unexpected HTTP status:", response.Status)
		return nil, fmt.Errorf("unexpected response status: %d", response.StatusCode)
	}

	var villages []map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&villages)
	if err != nil {
		log.Println("Error decoding JSON:", err)
		return nil, fmt.Errorf("failed to decode response")
	}

	return villages, nil
}
