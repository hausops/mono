package credential

import (
	"context"
	"errors"
	"fmt"
	"net/mail"

	"golang.org/x/crypto/bcrypt"
)

// Service provides the logic for saving and validating credentials.
type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) ExistByEmail(ctx context.Context, email mail.Address) bool {
	_, err := s.repo.FindByEmail(ctx, email)
	return err != nil
}

// Save upserts the credential to the credential repo.
func (s *Service) Save(ctx context.Context, email mail.Address, password []byte) error {
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("hash password: %w", err)
	}

	cred := Credential{
		Email:    email,
		Password: hashedPassword,
	}
	err = s.repo.Upsert(ctx, cred)
	if err != nil {
		return fmt.Errorf("repo.Upsert(%s): %w", email.Address, err)
	}
	return nil
}

// Validate checks if the provided credential is valid.
func (s *Service) Validate(ctx context.Context, email mail.Address, password []byte) error {
	cred, err := s.repo.FindByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return err
		}
		return fmt.Errorf("repo.FindByEmail(%s): %w", email.Address, err)
	}

	err = bcrypt.CompareHashAndPassword(cred.Password, password)
	if err != nil {
		return ErrInvalidPassword
	}

	return nil
}
