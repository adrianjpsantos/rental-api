package repositories

import (
	"context"
	"database/sql"

	"github.com/adrianjpsantos/rental-api/internal/domain/profile"
	"github.com/google/uuid"
)

type ProfileRepository struct {
	db DBTX
}

// Create implements [profile.Repository].
func (p *ProfileRepository) Create(
	ctx context.Context,
	input *profile.CreateInput,
) error {

	query := `
		INSERT INTO profiles (
			user_id,
			first_name,
			last_name,
			cpf,
			phone,
			birth_date,
			avatar_url
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	_, err := p.db.ExecContext(ctx, query,
		input.UserID,
		input.FirstName,
		input.LastName,
		input.CPF,
		input.Phone,
		input.BirthDate,
		input.AvatarURL,
	)

	if err != nil {
		return err
	}

	return nil
}

// Delete implements [profile.Repository].
func (p *ProfileRepository) Delete(
	ctx context.Context,
	userID uuid.UUID,
) error {
	_, err := p.db.ExecContext(ctx, `
		DELETE FROM profiles
		WHERE user_id = $1
	`, userID)

	return err
}

// FindByUserID implements [profile.Repository].
func (p *ProfileRepository) FindByUserID(
	ctx context.Context,
	userID uuid.UUID,
) (*profile.Profile, error) {
	result := new(profile.Profile)

	err := p.db.QueryRowContext(ctx, `
		SELECT
			user_id,
			first_name,
			last_name,
			cpf,
			phone,
			birth_date,
			avatar_url
		FROM profiles
		WHERE user_id = $1
	`, userID).Scan(
		&result.UserID,
		&result.FirstName,
		&result.LastName,
		&result.CPF,
		&result.Phone,
		&result.BirthDate,
		&result.AvatarURL,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, profile.ErrProfileNotFound
		}

		return nil, err
	}

	return result, nil
}

// Update implements [profile.Repository].
func (p *ProfileRepository) Update(
	ctx context.Context,
	input *profile.Profile,
) error {
	_, err := p.db.ExecContext(ctx, `
		UPDATE profiles
		SET
			first_name = $1,
			last_name = $2,
			phone = $3,
			birth_date = $4,
			avatar_url = $5
		WHERE user_id = $6
	`,
		input.FirstName,
		input.LastName,
		input.Phone,
		input.BirthDate,
		input.AvatarURL,
		input.UserID,
	)

	return err
}

func NewProfileRepository(db DBTX) profile.Repository {
	return &ProfileRepository{db: db}
}
