package service

import (
	"belajar-unit-test/entity"
	"belajar-unit-test/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var CategoryRepository = repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: &CategoryRepository}
func TestCategoryService_Get(t *testing.T){
	CategoryRepository.Mock.On("FindById","1").Return(nil)

category, error :=	categoryService.Get("1")

assert.Nil(t, category)
assert.NotNil(t, error)
}

func TestCategoryService_GetSuccess(t *testing.T){
	category := entity.Category{
		Id :"1",
		Name :"Laptop",
	}
	CategoryRepository.Mock.On("FindById", "2").Return(category)

	result, err :=	categoryService.Get("2")
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, category.Id, result.Id)
	assert.Equal(t, category.Name, result.Name)
}