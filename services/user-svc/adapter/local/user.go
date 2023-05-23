package local

import (
	"context"
	"net/mail"

	"github.com/google/uuid"
	"github.com/hausops/mono/services/user-svc/domain/user"
)

type userRepository struct {
	byID map[uuid.UUID]user.User

	// TODO: change byEmail to map[uuid.UUID]uuid.UUID (user ID)
	byEmail map[string]user.User
}

// NewUserRepository creates a new instance of the userRepository with
// an initial empty state.
//
// The returned user repository can be used to store and retrieve user information.
func NewUserRepository() *userRepository {
	return &userRepository{
		byID:    make(map[uuid.UUID]user.User),
		byEmail: make(map[string]user.User),
	}
}

// Ensure userRepository implements the user.Repository interface.
var _ user.Repository = (*userRepository)(nil)

func (r *userRepository) Delete(ctx context.Context, id uuid.UUID) (user.User, error) {
	u, ok := r.byID[id]
	if !ok {
		return user.User{}, user.ErrNotFound
	}
	delete(r.byID, id)
	delete(r.byEmail, u.Email.Address)
	return u, nil
}

func (r *userRepository) FindByID(ctx context.Context, id uuid.UUID) (user.User, error) {
	u, ok := r.byID[id]
	if !ok {
		return user.User{}, user.ErrNotFound
	}
	return u, nil
}

func (r *userRepository) FindByEmail(ctx context.Context, email mail.Address) (user.User, error) {
	u, ok := r.byEmail[email.Address]
	if !ok {
		return user.User{}, user.ErrNotFound
	}
	return u, nil
}

func (r *userRepository) Upsert(ctx context.Context, u user.User) (user.User, error) {
	if prev, ok := r.byEmail[u.Email.Address]; ok {
		// email must be unique
		if u.ID != prev.ID {
			return user.User{}, user.ErrEmailAlreadyUsed
		}
	}

	// if we're updating
	if prev, ok := r.byID[u.ID]; ok {
		// if changing email, need to update the "index"
		if u.Email.Address != prev.Email.Address {
			delete(r.byEmail, prev.Email.Address)
		}
	}

	r.byID[u.ID] = u
	r.byEmail[u.Email.Address] = u
	return u, nil
}