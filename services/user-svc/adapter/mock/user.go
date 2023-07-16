package mock

import (
	"context"
	"net/mail"

	"github.com/hausops/mono/services/user-svc/domain/user"
)

type userRepository struct {
	// byID maps user ID to user.User.
	byID map[user.ID]user.User

	// byEmail maps email address to user ID.
	byEmail map[mail.Address]user.ID
}

// NewUserRepository creates a new instance of the userRepository with
// an initial empty state.
//
// The returned user repository can be used to store and retrieve user information.
func NewUserRepository() *userRepository {
	return &userRepository{
		byID:    make(map[user.ID]user.User),
		byEmail: make(map[mail.Address]user.ID),
	}
}

// Ensure userRepository implements the user.Repository interface.
var _ user.Repository = (*userRepository)(nil)

func (r *userRepository) Delete(ctx context.Context, id user.ID) (user.User, error) {
	u, ok := r.byID[id]
	if !ok {
		return user.User{}, user.ErrNotFound
	}
	delete(r.byID, id)
	delete(r.byEmail, u.Email)
	return u, nil
}

func (r *userRepository) FindByID(ctx context.Context, id user.ID) (user.User, error) {
	u, ok := r.byID[id]
	if !ok {
		return user.User{}, user.ErrNotFound
	}
	return u, nil
}

func (r *userRepository) FindByEmail(ctx context.Context, email mail.Address) (user.User, error) {
	id, ok := r.byEmail[email]
	if !ok {
		return user.User{}, user.ErrNotFound
	}

	u, ok := r.byID[id]
	if !ok {
		return user.User{}, user.ErrNotFound
	}
	return u, nil
}

func (r *userRepository) Upsert(ctx context.Context, u user.User) (user.User, error) {
	// Check if the email address is already associated with another user.
	if prevID, ok := r.byEmail[u.Email]; ok {
		// email must be unique
		if u.ID != prevID {
			return user.User{}, user.ErrEmailTaken
		}
	}

	// If updating an existing user
	if prev, ok := r.byID[u.ID]; ok {
		delete(r.byEmail, prev.Email)
	}

	r.byID[u.ID] = u
	r.byEmail[u.Email] = u.ID
	return u, nil
}
