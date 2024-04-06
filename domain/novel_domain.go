package domain

import "CRUD-REDIS/model"

type NovelRepo interface {
	CreateNovel(createNovel model.Novel) error
	GetNovelyById(id int) (model.Novel, error)
}

type NovelUsecase interface {
	CreateNovel(createNovel model.Novel) error
	GetNovelyById(id int) (model.Novel, error)
}
