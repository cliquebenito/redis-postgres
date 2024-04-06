package usecase

import (
	"CRUD-REDIS/domain"
	"CRUD-REDIS/model"
)

type novelUseCase struct {
	novelRepo domain.NovelRepo
}

func (n *novelUseCase) CreateNovel(createNovel model.Novel) error {

	err := n.novelRepo.CreateNovel(createNovel)
	return err
}

func NewNovelUsecase(novelRepo domain.NovelRepo) domain.NovelUsecase {
	return &novelUseCase{novelRepo: novelRepo}
}
