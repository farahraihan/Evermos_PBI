package service

import (
	"errors"
	"evermos_pbi/internal/features/stores"
	"evermos_pbi/internal/utils"
	"log"
	"mime/multipart"
)

type StoreServices struct {
	qry        stores.SQuery
	cloudinary utils.CloudinaryUtilityInterface
}

func NewStoreService(q stores.SQuery, c utils.CloudinaryUtilityInterface) stores.SService {
	return &StoreServices{
		qry:        q,
		cloudinary: c,
	}
}

func (ss *StoreServices) AddStore(newStore stores.Store, src multipart.File, filename string) error {
	ownerExist, err := ss.qry.IsOwnerExist(newStore.UserID)
	if ownerExist {
		log.Println("user already has a store: ", err)
		return errors.New("user already has a store")
	}

	if src != nil && filename != "" {
		imageURL, err := ss.cloudinary.UploadToCloudinary(src, filename)
		if err != nil {
			log.Println("image upload failed: ", err)
			return errors.New("failed to upload image, please try again later")
		}
		newStore.StoreImage = imageURL
	}

	err = ss.qry.AddStore(newStore)
	if err != nil {
		log.Println("add store query error: ", err)
		return errors.New("failed to add a store, please try again later")
	}

	return nil
}

func (ss *StoreServices) UpdateStore(userID uint, updateStore stores.Store, src multipart.File, filename string) error {
	if src != nil && filename != "" {
		imageURL, err := ss.cloudinary.UploadToCloudinary(src, filename)
		if err != nil {
			log.Println("image upload failed: ", err)
			return errors.New("failed to upload image, please try again later")
		}
		updateStore.StoreImage = imageURL
	}

	err := ss.qry.UpdateStore(userID, updateStore)
	if err != nil {
		log.Println("update store query error: ", err)
		return errors.New("failed to update a store, please try again later")
	}

	return nil

}

func (ss *StoreServices) DeleteStore(userID uint) error {
	err := ss.qry.DeleteStore(userID)
	if err != nil {
		log.Println("delete store query error: ", err)
		return errors.New("failed to delete a store, please try again later")
	}

	return nil
}

func (ss *StoreServices) GetStoreByID(storeID uint) (stores.Store, error) {
	store, err := ss.qry.GetStoreByID(storeID)
	if err != nil {
		log.Println("get store by ID query error: ", err)
		return stores.Store{}, errors.New("failed to retrieve store, please try again later")
	}

	return store, nil
}

func (ss *StoreServices) GetStoreByUserID(userID uint) (stores.Store, error) {
	store, err := ss.qry.GetStoreByUserID(userID)
	if err != nil {
		log.Println("get store by ID query error: ", err)
		return stores.Store{}, errors.New("failed to retrieve store, please try again later")
	}

	return store, nil
}

func (ss *StoreServices) GetAllStores(limit uint, page uint, search string) ([]stores.Store, uint, error) {
	stores, totalItems, err := ss.qry.GetAllStores(limit, page, search)

	if err != nil {
		log.Println("get all stores query error: ", err)
		return nil, 0, errors.New("failed to retrieve store data, please try again later")
	}

	return stores, totalItems, nil
}
