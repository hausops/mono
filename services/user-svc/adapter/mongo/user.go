package mongo

import (
	"context"
	"errors"
	"fmt"
	"net/mail"
	"time"

	"github.com/google/uuid"
	"github.com/hausops/mono/services/user-svc/domain/user"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type userRepository struct {
	collection *mongo.Collection
}

// NewUserRepository creates a new userRepository instance
// and create a email index if one is not already set up.
func NewUserRepository(
	ctx context.Context,
	// Take collection so we can use a different database when testing.
	userCollection *mongo.Collection,
) (*userRepository, error) {
	if name := userCollection.Name(); name != "users" {
		return nil, fmt.Errorf("wrong collection name: %s", name)
	}

	if err := createEmailIndex(ctx, userCollection); err != nil {
		return nil, fmt.Errorf("create email index: %w", err)
	}

	return &userRepository{collection: userCollection}, nil
}

// createEmailIndex creates a unique index on the "email" field of userCollection.
func createEmailIndex(ctx context.Context, userCollection *mongo.Collection) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := userCollection.Indexes().CreateOne(
		ctx,
		mongo.IndexModel{
			Keys:    bson.M{"email": 1},
			Options: options.Index().SetUnique(true),
		},
	)
	if err != nil {
		return err
	}
	return nil
}

// Ensure userRepository implements the user.Repository interface.
var _ user.Repository = (*userRepository)(nil)

func (r *userRepository) Delete(ctx context.Context, id uuid.UUID) (user.User, error) {
	var u userBSON
	err := r.collection.FindOneAndDelete(ctx, bson.M{"_id": id}).Decode(&u)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return user.User{}, user.ErrNotFound
		}
		return user.User{},
			fmt.Errorf("users.FindOneAndDelete(_id: %s).Decode() err = %w", id, err)
	}

	return u.toUser(), nil
}

func (r *userRepository) FindByID(ctx context.Context, id uuid.UUID) (user.User, error) {
	var u userBSON
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&u)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return user.User{}, user.ErrNotFound
		}
		return user.User{},
			fmt.Errorf("users.FindOne(_id: %s).Decode() err = %w", id, err)
	}

	return u.toUser(), nil
}

func (r *userRepository) FindByEmail(ctx context.Context, email mail.Address) (user.User, error) {
	var u userBSON
	err := r.collection.FindOne(ctx, bson.M{"email": email}).Decode(&u)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return user.User{}, user.ErrNotFound
		}
		return user.User{},
			fmt.Errorf("users.FindOne(email: %s).Decode() err = %w", email, err)
	}

	return u.toUser(), nil
}

func (r *userRepository) Upsert(ctx context.Context, u user.User) (user.User, error) {
	if u.ID == (uuid.UUID{}) {
		return user.User{}, user.ErrMissingID
	}

	// Check if the email address is already associated with another user.
	//
	// Despite having a unique index on the email field, we perform this check
	// because the MongoDB duplicate key error doesn't provide access to
	// the specific key causing the issue.
	{
		existing, err := r.FindByEmail(ctx, u.Email)
		switch {
		// There's nothing to check the email is not used by anyone.
		case errors.Is(err, user.ErrNotFound):
			break
		case err != nil:
			return user.User{}, fmt.Errorf("find exising user by email: %w", err)
		case existing.ID != u.ID:
			return user.User{}, user.ErrEmailTaken
		}
	}

	up := toUserBSON(u)
	_, err := r.collection.UpdateByID(ctx, u.ID, bson.M{"$set": up},
		options.Update().SetUpsert(true))

	if err != nil {
		return user.User{},
			fmt.Errorf("users.UpdateByID(_id: %s) err = %w", u.ID, err)
	}
	return u, nil
}

type userBSON struct {
	ID          uuid.UUID    `bson:"_id"`
	Email       mail.Address `bson:"email"`
	DateCreated time.Time    `bson:"date_created"`
	DateUpdated time.Time    `bson:"date_updated"`
}

func (b *userBSON) toUser() user.User {
	return user.User{
		ID:          b.ID,
		Email:       b.Email,
		DateCreated: b.DateCreated,
		DateUpdated: b.DateUpdated,
	}
}

func toUserBSON(u user.User) userBSON {
	return userBSON{
		ID:          u.ID,
		Email:       u.Email,
		DateCreated: u.DateCreated,
		DateUpdated: u.DateUpdated,
	}
}
