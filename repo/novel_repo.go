package repo

import (
	"CRUD-REDIS/domain"
	"CRUD-REDIS/model"
	"database/sql"
	"github.com/redis/go-redis/v9"
)

type novelRepo struct {
	db  *sql.DB
	rdb *redis.Client
}

func (n *novelRepo) GetNovelyById(id int) (model.Novel, error) {
	var novel model.Novel
	query := "SELECT * FROM novels WHERE id = ?"
	err := n.db.QueryRow(query, id).Scan(&novel.Name, &novel.Author, &novel.Description)
	return novel, err
}

func (n *novelRepo) CreateNovel(createNovel model.Novel) error {

	query := "INSERT INTO novels (title, author, published_at) VALUES (?, ?, ?)"
	_, err := n.db.Exec(query, createNovel.Name, createNovel.Author, createNovel.Description)
	return err
}

func NewNovelRepo(db *sql.DB, rdb *redis.Client) domain.NovelRepo {
	return &novelRepo{
		db:  db,
		rdb: rdb,
	}
}
