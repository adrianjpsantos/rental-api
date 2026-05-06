package category

import "strings"

func (c *Category) Validate() error {
	if err := c.validateName(); err != nil {
		return err
	}
	if err := c.validateSlug(); err != nil {
		return err
	}
	if err := c.validateDescription(); err != nil {
		return err
	}
	return nil
}

func (c *Category) validateName() error {
	name := strings.TrimSpace(c.Name)
	if name == "" {
		return ErrInvalidName
	}
	if len(name) < 3 {
		return ErrNameTooShort
	}
	if len(name) > 80 {
		return ErrNameTooLong
	}
	return nil
}

func (c *Category) validateSlug() error {
	if c.Slug == "" {
		return ErrInvalidSlug
	}
	if !isValidSlug(c.Slug) {
		return ErrSlugInvalidFormat
	}
	return nil
}

func (c *Category) validateDescription() error {
	if len(strings.TrimSpace(c.Description)) > 500 {
		return ErrDescriptionTooLong
	}
	return nil
}
