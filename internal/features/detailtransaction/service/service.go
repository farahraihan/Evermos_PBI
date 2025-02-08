package service

import (
	"errors"
	"evermos_pbi/internal/features/detailtransaction"
	"log"
)

type DetailTransactionServices struct {
	qry detailtransaction.DQuery
}

func NewDetailTransactionService(q detailtransaction.DQuery) detailtransaction.DService {
	return &DetailTransactionServices{
		qry: q,
	}
}

func (ds *DetailTransactionServices) AddDetailTransaction(newDetailTransaction detailtransaction.DetailTransaction) error {
	err := ds.qry.AddDetailTransaction(newDetailTransaction)
	if err != nil {
		log.Println("add detail transaction query error: ", err)
		return errors.New("failed to add detail transaction, please try again later")
	}

	return nil
}

func (ds *DetailTransactionServices) UpdateDetailTransaction(productID uint, transactionID uint, updateDetailTransaction detailtransaction.DetailTransaction) error {
	err := ds.qry.UpdateDetailTransaction(productID, transactionID, updateDetailTransaction)
	if err != nil {
		log.Println("update detail transaction query error: ", err)
		return errors.New("failed to update detail transaction, please try again later")
	}

	return nil
}

func (ds *DetailTransactionServices) DeleteDetailTransaction(productID uint, transactionID uint) error {
	err := ds.qry.DeleteDetailTransaction(productID, transactionID)
	if err != nil {
		log.Println("delete detail transaction query error: ", err)
		return errors.New("failed to delete detail transaction, please try again later")
	}

	return nil
}

func (ds *DetailTransactionServices) IsProductInDetail(productID uint) (bool, error) {
	return ds.qry.IsProductInDetail(productID)
}
