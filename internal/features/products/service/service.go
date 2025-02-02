package service

import (
	"errors"
	"evermos_pbi/internal/features/logproduct"
	"evermos_pbi/internal/features/products"
	"evermos_pbi/internal/features/stores"
	"evermos_pbi/internal/utils"
	"log"
	"mime/multipart"
)

type ProductServices struct {
	qry        products.PQuery
	cloudinary utils.CloudinaryUtilityInterface
	sService   stores.SService
	lService   logproduct.LService
}

func NewProductService(q products.PQuery, c utils.CloudinaryUtilityInterface, s stores.SService, l logproduct.LService) products.PService {
	return &ProductServices{
		qry:        q,
		cloudinary: c,
		sService:   s,
		lService:   l,
	}
}

func (ps *ProductServices) AddProduct(userID uint, newProduct products.Product, src multipart.File, filename string) error {
	storeOwnedByUser, err := ps.sService.IsStoreOwnedByUser(newProduct.StoreID, userID)
	if err != nil {
		log.Println("failed to get owner: ", err)
		return errors.New("failed to get owner")
	}
	if !storeOwnedByUser {
		log.Println("user not the owner of this store")
		return errors.New("user not the owner of this store")
	}

	if src != nil && filename != "" {
		imageURL, err := ps.cloudinary.UploadToCloudinary(src, filename)
		if err != nil {
			log.Println("image upload failed: ", err)
			return errors.New("failed to upload image, please try again later")
		}
		newProduct.ProductImage = imageURL
	}

	err = ps.qry.AddProduct(&newProduct)
	if err != nil {
		log.Println("add product query error: ", err)
		return errors.New("failed to add a product, please try again later")
	}

	if newProduct.ID == 0 {
		return errors.New("failed to generate product ID")
	}

	newLogProduct := logproduct.LogProduct{
		ProductName:   newProduct.ProductName,
		ProductImage:  newProduct.ProductImage,
		Slug:          newProduct.Slug,
		ResellerPrice: newProduct.ResellerPrice,
		ConsumenPrice: newProduct.ConsumenPrice,
		Stock:         newProduct.Stock,
		Description:   newProduct.Description,
		ProductID:     newProduct.ID,
	}

	err = ps.lService.AddLogProduct(newLogProduct)
	if err != nil {
		log.Println("add log product query error: ", err)
		return errors.New("failed to add log product, please try again later")
	}

	return nil
}

func (ps *ProductServices) UpdateProduct(userID uint, productID uint, updateProduct products.Product, src multipart.File, filename string) error {
	productOwnedByUser, err := ps.qry.IsProductOwnedByUser(productID, userID)
	if err != nil {
		log.Println("failed to get owner: ", err)
		return errors.New("failed to get owner")
	}
	if !productOwnedByUser {
		log.Println("user not the owner of this product")
		return errors.New("user not the owner of this product")
	}

	if src != nil && filename != "" {
		imageURL, err := ps.cloudinary.UploadToCloudinary(src, filename)
		if err != nil {
			log.Println("image upload failed: ", err)
			return errors.New("failed to upload image, please try again later")
		}
		updateProduct.ProductImage = imageURL
	}

	err = ps.qry.UpdateProduct(productID, updateProduct)
	if err != nil {
		log.Println("update product query error: ", err)
		return errors.New("failed to update product, please try again later")
	}

	newLogProduct := logproduct.LogProduct{
		ProductName:   updateProduct.ProductName,
		ProductImage:  updateProduct.ProductImage,
		Slug:          updateProduct.Slug,
		ResellerPrice: updateProduct.ResellerPrice,
		ConsumenPrice: updateProduct.ConsumenPrice,
		Stock:         updateProduct.Stock,
		Description:   updateProduct.Description,
		ProductID:     productID,
	}

	err = ps.lService.AddLogProduct(newLogProduct)
	if err != nil {
		log.Println("add log product query error: ", err)
		return errors.New("failed to add log product, please try again later")
	}

	return nil
}

func (ps *ProductServices) DeleteProduct(userID uint, productID uint) error {
	productOwnedByUser, err := ps.qry.IsProductOwnedByUser(productID, userID)
	if err != nil {
		log.Println("failed to get owner: ", err)
		return errors.New("failed to get owner")
	}
	if !productOwnedByUser {
		log.Println("user not the owner of this product")
		return errors.New("user not the owner of this product")
	}

	err = ps.qry.DeleteProduct(productID)
	if err != nil {
		log.Println("delete product query error: ", err)
		return errors.New("failed to delete product, please try again later")
	}

	return nil
}

func (ps *ProductServices) GetProductByID(productID uint) (products.Product, error) {
	product, err := ps.qry.GetProductByID(productID)
	if err != nil {
		log.Println("get product by ID query error: ", err)
		return products.Product{}, errors.New("failed to retrieve product, please try again later")
	}

	return product, nil
}

func (ps *ProductServices) GetAllProducts(limit uint, page uint, search string) ([]products.Product, uint, error) {
	products, totalItems, err := ps.qry.GetAllProducts(limit, page, search)

	if err != nil {
		log.Println("get all products query error: ", err)
		return nil, 0, errors.New("failed to retrieve product data, please try again later")
	}

	return products, totalItems, nil
}

func (ps *ProductServices) GetProductsByStoreID(storeID uint, limit uint, page uint, search string) ([]products.Product, uint, error) {
	products, totalItems, err := ps.qry.GetProductsByStoreID(storeID, limit, page, search)

	if err != nil {
		log.Println("get products by store id query error: ", err)
		return nil, 0, errors.New("failed to retrieve product data, please try again later")
	}

	return products, totalItems, nil
}
