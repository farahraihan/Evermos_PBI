package service

import (
	"errors"
	"evermos_pbi/internal/features/categories"
	"evermos_pbi/internal/features/users"
	"log"
)

type CategoryServices struct {
	qry      categories.CQuery
	uService users.UService
}

func NewCategoryService(q categories.CQuery, u users.UService) categories.CService {
	return &CategoryServices{
		qry:      q,
		uService: u,
	}
}

func (cs *CategoryServices) AddCategory(userID uint, newCategory categories.Category) error {
	isAdmin, err := cs.uService.IsAdmin(userID)
	if err != nil || !isAdmin {
		log.Println("add category permission error: ", err)
		return errors.New("access denied")
	}

	err = cs.qry.AddCategory(newCategory)
	if err != nil {
		log.Println("add category query error: ", err)
		return errors.New("failed to add category, please try again later")
	}

	return nil
}

func (cs *CategoryServices) UpdateCategory(userID uint, categoryID uint, updateCategory categories.Category) error {
	isAdmin, err := cs.uService.IsAdmin(userID)
	if err != nil || !isAdmin {
		log.Println("update category permission error: ", err)
		return errors.New("access denied")
	}

	err = cs.qry.UpdateCategory(categoryID, updateCategory)
	if err != nil {
		log.Println("update category query error: ", err)
		return errors.New("failed to update category, please try again later")
	}

	return nil
}

func (cs *CategoryServices) DeleteCategory(userID uint, categoryID uint) error {
	isAdmin, err := cs.uService.IsAdmin(userID)
	if err != nil || !isAdmin {
		log.Println("delete category permission error: ", err)
		return errors.New("access denied")
	}

	err = cs.qry.DeleteCategory(categoryID)
	if err != nil {
		log.Println("delete category query error: ", err)
		return errors.New("failed to delete category, please try again later")
	}

	return nil

}

func (cs *CategoryServices) GetCategoryByID(categoryID uint) (categories.Category, error) {
	category, err := cs.qry.GetCategoryByID(categoryID)
	if err != nil {
		log.Println("get category by ID query error: ", err)
		return categories.Category{}, errors.New("failed to retrieve category, please try again later")
	}

	return category, nil
}

func (cs *CategoryServices) GetAllCategories(limit uint, page uint, search string) ([]categories.Category, uint, error) {
	categories, totalItems, err := cs.qry.GetAllCategories(limit, page, search)

	if err != nil {
		log.Println("get all categories query error: ", err)
		return nil, 0, errors.New("failed to retrieve category data, please try again later")
	}

	return categories, totalItems, nil
}
