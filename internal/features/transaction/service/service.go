package service

import (
	"errors"
	"evermos_pbi/internal/features/detailtransaction"
	"evermos_pbi/internal/features/transaction"
	"log"
)

type TransactionServices struct {
	qry      transaction.TQuery
	dService detailtransaction.DService
}

func NewTransactionService(q transaction.TQuery, d detailtransaction.DService) transaction.TService {
	return &TransactionServices{
		qry:      q,
		dService: d,
	}
}

func (ts *TransactionServices) AddTransaction(newTransaction transaction.Transaction, detail transaction.DetailTransaction2) error {
	trxID, err := ts.qry.CheckTransactionInCart(newTransaction.UserID)
	if err != nil {
		log.Println("failed to check in cart: ", err)
		return errors.New("failed to check in cart")
	}

	if trxID == 0 {
		err = ts.qry.AddTransaction(&newTransaction)
		if err != nil {
			log.Println("add transaction query error: ", err)
			return errors.New("failed to add transaction, please try again later")
		}
	}

	newDetailTransaction := detailtransaction.DetailTransaction{
		Quantity:      detail.Quantity,
		StoreID:       detail.StoreID,
		ProductID:     detail.ProductID,
		TransactionID: newTransaction.ID,
	}

	isProduct, err := ts.dService.IsProductInDetail(detail.ProductID)
	if err != nil {
		log.Println("failed to check product in detail transaction: ", err)
		return errors.New("failed to check product in cart")
	}

	if isProduct {
		err = ts.dService.UpdateDetailTransaction(detail.ProductID, newDetailTransaction.ID, newDetailTransaction)
		log.Println("update detail transaction query error: ", err)
		return errors.New("failed to update detail transaction, please try again later")
	}

	err = ts.dService.AddDetailTransaction(newDetailTransaction)
	if err != nil {
		log.Println("add detail transaction query error: ", err)
		return errors.New("failed to add detail transaction, please try again later")
	}

	return nil
}

func (ts *TransactionServices) UpdateDetailTransaction(userID uint, transactionID uint, productID uint, quantity uint) error {
	isOwnerTransaction, err := ts.qry.IsTransactionOwner(userID, transactionID)
	if err != nil {
		log.Println("failed to check user in transaction: ", err)
		return errors.New("failed to check user")
	}

	if !isOwnerTransaction {
		log.Println("user not the owner of this transaction")
		return errors.New("user not the owner of this transaction")
	}

	updateDetailTransaction := detailtransaction.DetailTransaction{
		Quantity: quantity,
	}

	err = ts.dService.UpdateDetailTransaction(productID, transactionID, updateDetailTransaction)
	if err != nil {
		log.Println("update detail transaction query error : ", err)
		return errors.New("failed to update detail transaction, please try again later")
	}

	return nil
}

func (ts *TransactionServices) DeleteTransaction(userID uint, transactionID uint, productID uint) error {
	isOwnerTransaction, err := ts.qry.IsTransactionOwner(userID, transactionID)
	if err != nil {
		log.Println("failed to check user in transaction: ", err)
		return errors.New("failed to check user")
	}

	if !isOwnerTransaction {
		log.Println("user not the owner of this transaction")
		return errors.New("user not the owner of this transaction")
	}

	err = ts.dService.DeleteDetailTransaction(productID, transactionID)
	if err != nil {
		log.Println("delete detail transaction query error : ", err)
		return errors.New("failed to delete detail transaction, please try again later")
	}

	return nil
}

func (ts *TransactionServices) UpdateTransaction(userID uint, transactionID uint, status string) error {
	isOwnerTransaction, err := ts.qry.IsTransactionOwner(userID, transactionID)
	if err != nil {
		log.Println("failed to check user in transaction: ", err)
		return errors.New("failed to check user")
	}

	if !isOwnerTransaction {
		log.Println("user not the owner of this transaction")
		return errors.New("user not the owner of this transaction")
	}

	err = ts.qry.UpdateTransaction(transactionID, status)
	if err != nil {
		log.Println("update status transaction query error : ", err)
		return errors.New("failed to update transaction, please try again later")
	}

	return nil
}

func (ts *TransactionServices) GetTransactionByStatusCart(userID uint, transactionID uint) (transaction.TransactionWithDetail, error) {
	isOwnerTransaction, err := ts.qry.IsTransactionOwner(userID, transactionID)
	if err != nil {
		log.Println("failed to check user in transaction: ", err)
		return transaction.TransactionWithDetail{}, errors.New("failed to check user")
	}

	if !isOwnerTransaction {
		log.Println("user not the owner of this transaction")
		return transaction.TransactionWithDetail{}, errors.New("user not the owner of this transaction")
	}

	cart, err := ts.qry.GetTransactionByStatusCart(transactionID)
	if err != nil {
		log.Println("get transaction by status cart error : ", err)
		return transaction.TransactionWithDetail{}, errors.New("failed to retrieve cart, please try again later")
	}

	return cart, nil
}

func (ts *TransactionServices) GetTransactionHistory(userID uint, limit uint, page uint) ([]transaction.TransactionWithDetail, uint, error) {
	transactionHistory, totalItems, err := ts.qry.GetTransactionHistory(userID, limit, page)
	if err != nil {
		log.Println("get transcation history query error : ", err)
		return []transaction.TransactionWithDetail{}, 0, errors.New("failed to retrieved transaction history, please try again later")
	}

	return transactionHistory, totalItems, nil
}

func (ts *TransactionServices) GetTransactionByID(userID uint, transactionID uint) (transaction.TransactionWithDetail, error) {
	isOwnerTransaction, err := ts.qry.IsTransactionOwner(userID, transactionID)
	if err != nil {
		log.Println("failed to check user in transaction: ", err)
		return transaction.TransactionWithDetail{}, errors.New("failed to check user")
	}

	if !isOwnerTransaction {
		log.Println("user not the owner of this transaction")
		return transaction.TransactionWithDetail{}, errors.New("user not the owner of this transaction")
	}

	invoice, err := ts.qry.GetTransactionByID(transactionID)
	if err != nil {
		log.Println("get transaction by ID query error : ", err)
		return transaction.TransactionWithDetail{}, nil
	}

	return invoice, nil
}
