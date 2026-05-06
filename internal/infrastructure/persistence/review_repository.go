package persistence

import (
	"context"
	"database/sql"

	"github.com/adrianjpsantos/rental-api/internal/domain/review"
	"github.com/google/uuid"
)

type ReviewRepository struct {
	db *sql.DB
}

func NewReviewRepository(db *sql.DB) review.InterfaceReviewRepository {
	return &ReviewRepository{db: db}
}

func (r *ReviewRepository) Create(ctx context.Context, review *review.Review) error {
	query := `
		INSERT INTO reviews (
			id, rental_id, reviewer_id, reviewed_id,
			item_id, rating, comment, review_type, created_at
		) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)
	`

	_, err := r.db.ExecContext(ctx, query,
		review.Id,
		review.RentalID,
		review.ReviewerID,
		review.ReviewedID,
		review.ItemID,
		review.Rating,
		review.Comment,
		review.ReviewType,
		review.CreatedAt,
	)

	return err
}

func (r *ReviewRepository) GetByID(ctx context.Context, id uuid.UUID) (*review.Review, error) {
	query := `
		SELECT id, rental_id, reviewer_id, reviewed_id,
		       item_id, rating, comment, review_type, created_at
		FROM reviews
		WHERE id = $1
	`

	var rv review.Review

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&rv.Id,
		&rv.RentalID,
		&rv.ReviewerID,
		&rv.ReviewedID,
		&rv.ItemID,
		&rv.Rating,
		&rv.Comment,
		&rv.ReviewType,
		&rv.CreatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, review.ErrReviewNotFound
	}
	if err != nil {
		return nil, err
	}

	return &rv, nil
}
func (r *ReviewRepository) GetByRentalID(ctx context.Context, rentalID uuid.UUID) ([]*review.Review, error) {
	query := `
		SELECT id, rental_id, reviewer_id, reviewed_id,
		       item_id, rating, comment, review_type, created_at
		FROM reviews
		WHERE rental_id = $1
	`

	rows, err := r.db.QueryContext(ctx, query, rentalID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []*review.Review

	for rows.Next() {
		var rv review.Review

		err := rows.Scan(
			&rv.Id,
			&rv.RentalID,
			&rv.ReviewerID,
			&rv.ReviewedID,
			&rv.ItemID,
			&rv.Rating,
			&rv.Comment,
			&rv.ReviewType,
			&rv.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		list = append(list, &rv)
	}

	if len(list) == 0 {
		return nil, review.ErrReviewNotFound
	}

	return list, nil
}

func (r *ReviewRepository) GetByReviewedID(ctx context.Context, reviewedID uuid.UUID, reviewType *review.ReviewType) ([]*review.Review, error) {
	query := `
		SELECT id, rental_id, reviewer_id, reviewed_id,
		       item_id, rating, comment, review_type, created_at
		FROM reviews
		WHERE reviewed_id = $1
	`

	args := []interface{}{reviewedID}

	if reviewType != nil {
		query += " AND review_type = $2"
		args = append(args, *reviewType)
	}

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []*review.Review

	for rows.Next() {
		var rv review.Review

		err := rows.Scan(
			&rv.Id,
			&rv.RentalID,
			&rv.ReviewerID,
			&rv.ReviewedID,
			&rv.ItemID,
			&rv.Rating,
			&rv.Comment,
			&rv.ReviewType,
			&rv.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		list = append(list, &rv)
	}

	if len(list) == 0 {
		return nil, review.ErrReviewNotFound
	}

	return list, nil
}
func (r *ReviewRepository) ExistsByRentalAndReviewer(ctx context.Context, rentalID, reviewerID uuid.UUID) (bool, error) {
	query := `
		SELECT EXISTS (
			SELECT 1 FROM reviews
			WHERE rental_id = $1 AND reviewer_id = $2
		)
	`

	var exists bool

	err := r.db.QueryRowContext(ctx, query, rentalID, reviewerID).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}
func (r *ReviewRepository) ListUserReviews(ctx context.Context, userID uuid.UUID) ([]*review.Review, error) {
	query := `
		SELECT id, rental_id, reviewer_id, reviewed_id,
		       item_id, rating, comment, review_type, created_at
		FROM reviews
		WHERE reviewer_id = $1
	`

	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []*review.Review

	for rows.Next() {
		var rv review.Review

		err := rows.Scan(
			&rv.Id,
			&rv.RentalID,
			&rv.ReviewerID,
			&rv.ReviewedID,
			&rv.ItemID,
			&rv.Rating,
			&rv.Comment,
			&rv.ReviewType,
			&rv.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		list = append(list, &rv)
	}

	if len(list) == 0 {
		return nil, review.ErrReviewNotFound
	}

	return list, nil
}
