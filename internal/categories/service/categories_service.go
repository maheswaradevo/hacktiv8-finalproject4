package service

import (
	"context"
	"log"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/maheswaradevo/hacktiv8-finalproject4/internal/categories"
	"github.com/maheswaradevo/hacktiv8-finalproject4/internal/dto"
	"github.com/maheswaradevo/hacktiv8-finalproject4/internal/transaction"
	"github.com/maheswaradevo/hacktiv8-finalproject4/pkg/constants"
	"github.com/maheswaradevo/hacktiv8-finalproject4/pkg/errors"
)

type CategoriesServiceImpl struct {
	repo            categories.CategoriesRepository
	transactionRepo transaction.TransactionRepository
}

func ProvideCategoriesService(repo categories.CategoriesRepository, transactionRepo transaction.TransactionRepository) *CategoriesServiceImpl {
	return &CategoriesServiceImpl{
		repo: repo,
	}
}

func (ctg *CategoriesServiceImpl) CreateCategories(ctx context.Context, data *dto.CreateCategoriesRequest, userID uint64) (res *dto.CreateCategoriesResponse, err error) {
	categoriesData := data.ToCategoriesEntity()

	userInfo, errGetRole := ctg.transactionRepo.FindRoleByUserID(ctx, userID)
	if errGetRole != nil {
		log.Printf("[ViewUsersTransaction] failed to find role, err: %v", (errGetRole))
		return nil, errGetRole
	}
	if strings.ToLower(userInfo.Role) != constants.CustomerRole {
		errOnlyAdmin := errors.ErrOnlyAdmin
		log.Printf("[ViewUsersTransaction] only admin can acces, err: %v", (errOnlyAdmin))
		return nil, errOnlyAdmin
	}

	validate := validator.New()
	validateError := validate.Struct(data)
	if validateError != nil {
		validateError = errors.ErrInvalidRequestBody
		log.Printf("[CreateCategory] there's data that not through the validate process")
		return nil, validateError
	}
	categoryID, err := ctg.repo.CreateCategories(ctx, *categoriesData)
	if err != nil {
		log.Printf("[CreateCategory] failed to store user data to database: %v", err)
		return
	}
	return dto.NewCategoriesCreateResponse(*categoriesData, userID, categoryID), nil
}

func (ctg *CategoriesServiceImpl) ViewCategories(ctx context.Context) ([]dto.ViewCategoriesResponse, error) {
	count, err := ctg.repo.CountCategories(ctx)
	if err != nil {
		log.Printf("[ViewCategory] failed to count the category, err: %v", err)
		return nil, err
	}
	if count == 0 {
		err = errors.ErrDataNotFound
		log.Printf("[ViewCategory] no data exists in the database: %v", err)
		return nil, err
	}
	res, err := ctg.repo.ViewCategories(ctx)
	if err != nil {
		log.Printf("[ViewCategory] failed to view the category, err: %v", err)
		return nil, err
	}
	categoriesResponse := dto.NewViewCategoriesResponse(res)
	return categoriesResponse, nil
}

func (ctg *CategoriesServiceImpl) UpdateCategories(ctx context.Context, categoryID uint64, userID uint64, data *dto.EditCategoriesRequest) (*dto.EditCategoriesResponse, error) {
	editedCategories := data.ToCategoriesEntity()

	userInfo, errGetRole := ctg.transactionRepo.FindRoleByUserID(ctx, userID)
	if errGetRole != nil {
		log.Printf("[ViewUsersTransaction] failed to find role, err: %v", (errGetRole))
		return nil, errGetRole
	}
	if strings.ToLower(userInfo.Role) != constants.CustomerRole {
		errOnlyAdmin := errors.ErrOnlyAdmin
		log.Printf("[ViewUsersTransaction] only admin can acces, err: %v", (errOnlyAdmin))
		return nil, errOnlyAdmin
	}

	check, err := ctg.repo.CheckCategories(ctx, categoryID)
	if err != nil {
		log.Printf("[UpdateCategory] failed to check category, err: %v", err)
		return nil, err
	}
	if !check {
		err = errors.ErrDataNotFound
		log.Printf("[UpdateCategory] no Category found")
		return nil, err
	}
	err = ctg.repo.UpdateCategories(ctx, *editedCategories, categoryID)
	if err != nil {
		log.Printf("[UpdateCategory] failed to update category, err: %v", err)
		return nil, err
	}
	categories, err := ctg.repo.GetCategoriesByID(ctx, categoryID)
	if err != nil {
		log.Printf("[UpdateComment] failed to get photo, err: %v", err)
		return nil, err
	}
	return categories, nil
}

func (ctg *CategoriesServiceImpl) DeleteCategories(ctx context.Context, categoryID uint64, userID uint64) (*dto.DeleteCategoriesResponse, error) {
	userInfo, errGetRole := ctg.transactionRepo.FindRoleByUserID(ctx, userID)
	if errGetRole != nil {
		log.Printf("[ViewUsersTransaction] failed to find role, err: %v", (errGetRole))
		return nil, errGetRole
	}
	if strings.ToLower(userInfo.Role) != constants.CustomerRole {
		errOnlyAdmin := errors.ErrOnlyAdmin
		log.Printf("[ViewUsersTransaction] only admin can acces, err: %v", (errOnlyAdmin))
		return nil, errOnlyAdmin
	}

	check, err := ctg.repo.CheckCategories(ctx, categoryID)
	if err != nil {
		log.Printf("[DeleteCategory] failed to check category, err: %v", err)
		return nil, err
	}
	if !check {
		err = errors.ErrDataNotFound
		log.Printf("[DeleteCategory] no category")
		return nil, err
	}

	err = ctg.repo.DeleteCategories(ctx, categoryID)
	if err != nil {
		log.Printf("[DeleteCategory] failed to delete category id: %v", categoryID)
		return nil, err
	}
	message := "Category has been successfully deleted"
	return dto.NewDeleteCategoriesResponse(message), nil
}
