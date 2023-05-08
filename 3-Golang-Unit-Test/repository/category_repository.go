package repository

import "3-Golang-Unit-Test/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}