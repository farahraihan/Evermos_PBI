package repository

import (
	"evermos_pbi/internal/features/stores"

	"gorm.io/gorm"
)

type StoreQuery struct {
	db *gorm.DB
}

func NewStoreQuery(connect *gorm.DB) stores.SQuery {
	return &StoreQuery{
		db: connect,
	}
}

func (sq *StoreQuery) AddStore(newStore stores.Store) error {
	cnvData := ToStoreQuery(newStore)
	qry := sq.db.Create(&cnvData)

	if qry.Error != nil {
		return qry.Error
	}

	return nil
}

func (sq *StoreQuery) UpdateStore(userID uint, updateStore stores.Store) error {
	cnvData := ToStoreQuery(updateStore)

	qry := sq.db.Where("user_id = ?", userID).Updates(&cnvData)

	if qry.Error != nil {
		return qry.Error
	}

	if qry.RowsAffected < 1 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (sq *StoreQuery) DeleteStore(userID uint) error {
	qry := sq.db.Where("user_id = ?", userID).Delete(&Store{})

	if qry.Error != nil {
		return qry.Error
	}

	if qry.RowsAffected < 1 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (sq *StoreQuery) GetStoreByID(storeID uint) (stores.Store, error) {
	var store stores.Store

	err := sq.db.First(&store, storeID).Error
	if err != nil {
		return stores.Store{}, err
	}
	return store, nil
}

func (sq *StoreQuery) GetStoreByUserID(userID uint) (stores.Store, error) {
	var store stores.Store

	err := sq.db.Where("user_id = ?", userID).First(&store).Error
	if err != nil {
		return stores.Store{}, err
	}

	return store, nil
}

func (sq *StoreQuery) GetAllStores(limit uint, page uint, search string) ([]stores.Store, uint, error) {
	var storesList []Store
	var totalItems int64

	offset := (page - 1) * limit

	qry := sq.db.Model(&Store{})
	if search != "" {
		qry = qry.Where("store_name ILIKE ?", "%"+search+"%")
	}
	if err := qry.Count(&totalItems).Error; err != nil {
		return nil, 0, err
	}

	qry = sq.db.Preload("User").Limit(int(limit)).Offset(int(offset)).Find(&storesList)
	if search != "" {
		qry = qry.Where("store_name ILIKE ?", "%"+search+"%")
	}
	if qry.Error != nil {
		return nil, 0, qry.Error
	}

	storesEntities := make([]stores.Store, len(storesList))
	for i, store := range storesList {
		storesEntities[i] = store.ToStoreEntity()
	}

	return storesEntities, uint(totalItems), nil
}

func (sq *StoreQuery) IsOwnerExist(userID uint) (bool, error) {
	var count int64
	if err := sq.db.Model(&stores.Store{}).Where("user_id = ?", userID).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}
