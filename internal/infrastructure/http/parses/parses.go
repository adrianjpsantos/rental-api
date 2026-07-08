package parses

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/adrianjpsantos/rental-api/internal/domain/item"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type Request[T any] struct {
	Data T `json:"data"`
}

var (
	ErrEmptyQueryParam = errors.New("query parameter is empty")
	ErrInvalidUUID     = errors.New("invalid uuid parameter")
	ErrInvalidBody     = errors.New("invalid request body")
	ErrInvalidInt      = errors.New("invalid int query parameter")
	ErrInvalidBool     = errors.New("invalid bool query parameter")
	ErrInvalidFloat    = errors.New("invalid float query parameter")
	ErrInvalidDate     = errors.New("invalid date query parameter")
)

func ParseBody[T any](c fiber.Ctx) (T, error) {
	var req Request[T]

	if err := c.Bind().Body(&req); err != nil {
		var zero T
		return zero, err
	}
	validate := validator.New()

	if err := validate.Struct(req.Data); err != nil {
		var zero T
		return zero, fmt.Errorf("validation error: %w", err)
	}

	return req.Data, nil
}

func ParseStringQuery(c fiber.Ctx, key string) *string {
	value := strings.TrimSpace(c.Query(key))

	if value == "" {
		return nil
	}

	return &value
}

func ParseUUIDParam(c fiber.Ctx, key string) (*uuid.UUID, error) {
	value := strings.TrimSpace(c.Query(key))

	if value == "" {
		return nil, nil
	}

	parsed, err := uuid.Parse(value)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrInvalidUUID, key)
	}

	return &parsed, nil
}

func ParseIntQuery(c fiber.Ctx, key string) (*int, error) {
	value := strings.TrimSpace(c.Query(key))

	if value == "" {
		return nil, nil
	}

	parsed, err := strconv.Atoi(value)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrInvalidInt, key)
	}

	return &parsed, nil
}

func ParseBoolQuery(c fiber.Ctx, key string) (*bool, error) {
	value := strings.TrimSpace(c.Query(key))

	if value == "" {
		return nil, nil
	}

	parsed, err := strconv.ParseBool(value)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrInvalidBool, key)
	}

	return &parsed, nil
}

func ParseFloatQuery(c fiber.Ctx, key string) (*float64, error) {
	value := strings.TrimSpace(c.Query(key))

	if value == "" {
		return nil, nil
	}

	parsed, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrInvalidFloat, key)
	}

	return &parsed, nil
}

func ParseDateQuery(c fiber.Ctx, key, layout string) (*time.Time, error) {
	value := strings.TrimSpace(c.Query(key))

	if value == "" {
		return nil, nil
	}

	parsed, err := time.Parse(layout, value)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrInvalidDate, key)
	}

	return &parsed, nil
}

func ParseItemFilter(c fiber.Ctx) (*item.ItemFilter, error) {
	ownerID, err := ParseUUIDParam(c, "owner_id")
	if err != nil {
		return nil, err
	}

	categoryID, err := ParseUUIDParam(c, "category_id")
	if err != nil {
		return nil, err
	}

	minPrice, err := ParseFloatQuery(c, "min_price")
	if err != nil {
		return nil, err
	}

	maxPrice, err := ParseFloatQuery(c, "max_price")
	if err != nil {
		return nil, err
	}

	limit, err := ParseIntQuery(c, "limit")
	if err != nil {
		return nil, err
	}

	offset, err := ParseIntQuery(c, "offset")
	if err != nil {
		return nil, err
	}

	location := ParseStringQuery(c, "location")

	return &item.ItemFilter{
		OwnerID:    ownerID,
		CategoryID: categoryID,
		MinPrice:   minPrice,
		MaxPrice:   maxPrice,
		Location:   location,
		Limit:      limit,
		Offset:     offset,
	}, nil
}
