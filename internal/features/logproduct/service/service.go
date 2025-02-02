package service

import (
	"errors"
	"evermos_pbi/internal/features/logproduct"
	"log"
)

type LogProductServices struct {
	qry logproduct.LQuery
}

func NewLogProductService(q logproduct.LQuery) logproduct.LService {
	return &LogProductServices{
		qry: q,
	}
}

func (ls *LogProductServices) AddLogProduct(newLogProduct logproduct.LogProduct) error {
	err := ls.qry.AddLogProduct(newLogProduct)
	if err != nil {
		log.Println("add log_product query error: ", err)
		return errors.New("failed to add log product, please try again later")
	}

	return nil
}

func (ls *LogProductServices) GetLogProductByID(logProductID uint) (logproduct.LogProduct, error) {
	logProduct, err := ls.qry.GetLogProductByID(logProductID)
	if err != nil {
		log.Println("get log_product by ID query error: ", err)
		return logproduct.LogProduct{}, errors.New("failed to retrieve log_product, please try again later")
	}

	return logProduct, nil
}

func (ls *LogProductServices) GetAllLogProduct(limit uint, page uint, search string) ([]logproduct.LogProduct, uint, error) {
	logProduct, totalItems, err := ls.qry.GetAllLogProduct(limit, page, search)

	if err != nil {
		log.Println("get all log_product query error: ", err)
		return nil, 0, errors.New("failed to retrieve log_product, please try again later")
	}

	return logProduct, totalItems, nil
}
