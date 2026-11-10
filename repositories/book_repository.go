package repositories

import (
	"context"
	"fmt"
	"go-postgres-api/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type BookRepository struct{ DB *pgxpool.Pool }

func NewBookRepository(pool *pgxpool.Pool) *BookRepository {
	return &BookRepository{DB: pool}
}

func (r *BookRepository) GetAll(ctx context.Context) ([]models.Book, error) {
	const query = `
		SELECT id, title, author, price, created_at
		FROM books
		ORDER BY id
	`

	rows, err := r.DB.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("olishda xatolik: %w ❗️❗️❗️", err)
	}
	defer rows.Close()

	books := make([]models.Book, 0)

	for rows.Next() {
		var book models.Book

		if err := rows.Scan(
			&book.ID,
			&book.Title,
			&book.Author,
			&book.Price,
			&book.CreatedAt,
		); err != nil {
			return nil, fmt.Errorf("malumotlarini o'qishda xatolik: %w ❗️❗️❗️", err)
		}

		books = append(books, book)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("kitoblarni o'qishda xatolik: %w ❗️❗️❗️", err)
	}

	return books, nil
}

func (r *BookRepository) GetByID(ctx context.Context, id int64) (models.Book, error) {
	const query = `
		SELECT id, title, author, price, created_at
		FROM books
		WHERE id = $1
	`

	var book models.Book

	err := r.DB.QueryRow(ctx, query, id).Scan(
		&book.ID,
		&book.Title,
		&book.Author,
		&book.Price,
		&book.CreatedAt,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return models.Book{}, pgx.ErrNoRows
		}

		return models.Book{}, fmt.Errorf("id bo'yicha olishda xatolik: %w ❗️❗️❗️", err)
	}

	return book, nil
}

func (r *BookRepository) Create(ctx context.Context, book models.Book) (models.Book, error) {
	const query = `
        INSERT INTO books (title, author, price) 
        VALUES ($1, $2, $3) 
        RETURNING id, title, author, price, created_at;
    `

	var createdBook models.Book

	err := r.DB.QueryRow(ctx, query, book.Title, book.Author, book.Price).Scan(
		&createdBook.ID,
		&createdBook.Title,
		&createdBook.Author,
		&createdBook.Price,
		&createdBook.CreatedAt,
	)
	if err != nil {
		return models.Book{}, fmt.Errorf("yaratishda xatolik: %w ❗️❗️❗️", err)
	}

	return createdBook, nil
}

func (r *BookRepository) Update(ctx context.Context, id int64, book models.Book) (models.Book, error) {
	const query = `
		UPDATE books
		SET title = $1, author = $2, price = $3
		WHERE id = $4
		RETURNING id, title, author, price, created_at;
	`
	var createdBook models.Book
	err := r.DB.QueryRow(ctx, query, book.Title, book.Author, book.Price, id).Scan(
		&createdBook.ID,
		&createdBook.Title,
		&createdBook.Author,
		&createdBook.Price,
		&createdBook.CreatedAt,
	)
	if err != nil {
		return models.Book{}, fmt.Errorf("kitob yaratishda xatolik: %w❗️❗️❗️", err)
	}

	return createdBook, nil
}

func (r *BookRepository) Delete(ctx context.Context, id int64) error {
	const query = `
		DELETE FROM books
		WHERE id = $1;
	`

	result, err := r.DB.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("xatolik > > h: %w", err)
	}

	if result.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	return nil
}