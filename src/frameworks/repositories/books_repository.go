package repositories

import (
	"database/sql"
	"restapi/src/applications/protocols"
	"restapi/src/entities"
)

// booksRepository ...
type booksRepository struct {
	db *sql.DB
}

// Create ...
func (repo *booksRepository) Create(entity *entities.BookEntity) (*entities.BookEntity, error) {
	sql := "INSERT INTO books (title, author, publishing_company, edition) VALUES ($1, $2, $3, $4) RETURNING *"

	prepare, err := repo.db.Prepare(sql)
	if err != nil {
		return nil, err
	}

	err = prepare.QueryRow(
		entity.Title,
		entity.Author,
		entity.PublishingCompany,
		entity.Edition,
	).Scan(
		&entity.ID,
		&entity.Title,
		&entity.Author,
		&entity.PublishingCompany,
		&entity.Edition,
		&entity.CreatedAt,
		&entity.UpdatedAt,
		&entity.DeletedAt,
	)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

// FindByTitle ...
func (b *booksRepository) FindByTitle(title string) (*entities.BookEntity, error) {
	sql := "SELECT id, title, author, publishing_company, edition, created_at, updated_at, deleted_at FROM books WHERE title = $1"

	prepare, err := b.db.Prepare(sql)
	if err != nil {
		return nil, err
	}
	entity := &entities.BookEntity{}
	err = prepare.QueryRow(title).Scan(
		&entity.ID,
		&entity.Title,
		&entity.Author,
		&entity.PublishingCompany,
		&entity.Edition,
		&entity.CreatedAt,
		&entity.UpdatedAt,
		&entity.DeletedAt,
	)

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, nil
		}
		return nil, err
	}

	return entity, nil
}

// FindByID ...
func (b *booksRepository) FindByID(id uint64) (*entities.BookEntity, error) {
	sql := "SELECT id, title, author, publishing_company, edition, created_at, updated_at, deleted_at FROM books WHERE id = $1"

	prepare, err := b.db.Prepare(sql)
	if err != nil {
		return nil, err
	}
	entity := &entities.BookEntity{}
	err = prepare.QueryRow(id).Scan(
		&entity.ID,
		&entity.Title,
		&entity.Author,
		&entity.PublishingCompany,
		&entity.Edition,
		&entity.CreatedAt,
		&entity.UpdatedAt,
		&entity.DeletedAt,
	)

	if err != nil {
		return nil, err
	}

	return entity, nil
}

// FindAll ...
func (b *booksRepository) FindAll() ([]entities.BookEntity, error) {
	sql := "SELECT id, title, author, publishing_company, edition, created_at, updated_at, deleted_at FROM books"

	entity := entities.BookEntity{}
	entitySlice := []entities.BookEntity{}

	rows, err := b.db.Query(sql)

	for rows.Next() {
		err = rows.Scan(
			&entity.ID,
			&entity.Title,
			&entity.Author,
			&entity.PublishingCompany,
			&entity.Edition,
			&entity.CreatedAt,
			&entity.UpdatedAt,
			&entity.DeletedAt,
		)
		entitySlice = append(entitySlice, entity)

		if err != nil {
			break
		}
	}

	if err != nil {
		return nil, err
	}

	return entitySlice, nil
}

// // Update ...
// func (b *booksRepository) Update(id uint64, book *entities.InputCreateBook) (*entities.BookEntity, error) {
// 	var set string

// 	if book.Author != "" {
// 		set += fmt.Sprintf("author = '%s', ", book.Author)
// 	}
// 	if book.Edition != 0 {
// 		set += fmt.Sprintf("edition = %d, ", book.Edition)
// 	}
// 	if book.PublishingCompany != "" {
// 		set += fmt.Sprintf("publishing_company = '%s', ", book.PublishingCompany)
// 	}
// 	if book.Title != "" {
// 		set += fmt.Sprintf("title = '%s', ", book.Title)
// 	}

// 	set = set[:len(set)-2]
// 	sql := "UPDATE books SET " + set + " WHERE id = ? RETURNING *"

// 	prepare, err := b.db.Prepare(sql)
// 	if err != nil {
// 		fmt.Println(sql, err)
// 		return nil, err
// 	}

// 	entity := &entities.BookEntity{}
// 	err = prepare.QueryRow(id).Scan(
// 		&entity.ID,
// 		&entity.Title,
// 		&entity.Author,
// 		&entity.PublishingCompany,
// 		&entity.Edition,
// 		&entity.CreatedAt,
// 		&entity.UpdatedAt,
// 		&entity.DeletedAt,
// 	)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return entity, nil
// }

// Delete ...
func (b *booksRepository) Delete(id uint64) (*entities.BookEntity, error) {
	sql := "DELETE FROM books WHERE id = ? RETURNING *"

	prepare, err := b.db.Prepare(sql)
	if err != nil {
		return nil, err
	}

	entity := &entities.BookEntity{}
	err = prepare.QueryRow(id).Scan(
		&entity.ID,
		&entity.Title,
		&entity.Author,
		&entity.PublishingCompany,
		&entity.Edition,
		&entity.CreatedAt,
		&entity.UpdatedAt,
		&entity.DeletedAt,
	)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

// // Delete ...
// func (b *booksRepository) SoftDelete(id uint64) (*entities.BookEntity, error) {

// 	return nil, nil
// }

// NewBooksRepository ...
func NewBooksRepository(dbConnection *sql.DB) protocols.IBooksRepository {
	return &booksRepository{db: dbConnection}
}
