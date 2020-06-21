package user

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// CollectionName is the mongodb collection name
const CollectionName = "users"

// Repository related errors
var (
	ErrCreateUser   = errors.New("can't create user")
	ErrGetUser      = errors.New("can't retrieve user")
	ErrListUsers    = errors.New("can't retrieve users")
	ErrUpdateUser   = errors.New("can't update user")
	ErrDeleteUser   = errors.New("can't delete user")
	ErrUserNotFound = errors.New("user not found")
)

// Repository represents the user repository
type Repository struct {
	col     *mongo.Collection
	userSvc *Service
}

// NewRepository returns a new Repository
func NewRepository(db *mongo.Database, userSvc *Service) *Repository {
	return &Repository{col: db.Collection(CollectionName), userSvc: userSvc}
}

// Create creates a new user
func (r *Repository) Create(ctx context.Context, user *Model) (*Model, error) {
	user.SetDefaults()

	// Hash the password if not hashed already
	if user.Password != "" && !r.userSvc.IsPasswordHashed([]byte(user.Password)) {
		if hashedPassword, err := r.userSvc.HashPassword([]byte(user.Password)); err == nil {
			user.Password = string(hashedPassword)
		}
	}

	res, err := r.col.InsertOne(ctx, user)
	if err != nil {
		return nil, ErrCreateUser
	}

	userID, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, ErrCreateUser
	}

	user.ID = userID

	user, err = r.Get(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// Get retrieves a single user
func (r *Repository) Get(ctx context.Context, user *Model) (*Model, error) {
	if err := r.col.FindOne(ctx, Model{ID: user.ID}).Decode(user); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, ErrUserNotFound
		}
		return nil, ErrGetUser
	}

	return user, nil
}

// List retrieves a list of user
func (r *Repository) List(ctx context.Context, filter interface{}) ([]Model, error) {
	users := []Model{}

	cur, err := r.col.Find(ctx, filter)
	if err != nil {
		return nil, ErrListUsers
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var user Model
		cur.Decode(&user)
		users = append(users, user)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

// Update updates a user
func (r *Repository) Update(ctx context.Context, user *Model) (*Model, error) {
	user.SetDefaults()

	// Hash the password if not hashed already
	if user.Password != "" && !r.userSvc.IsPasswordHashed([]byte(user.Password)) {
		if hashedPassword, err := r.userSvc.HashPassword([]byte(user.Password)); err == nil {
			user.Password = string(hashedPassword)
		}
	}

	if err := r.col.FindOneAndReplace(ctx, Model{ID: user.ID}, user).Decode(user); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, ErrUserNotFound
		}

		return nil, ErrUpdateUser
	}

	return user, nil
}

// Delete deletes a user
func (r *Repository) Delete(ctx context.Context, user *Model) (*Model, error) {
	if err := r.col.FindOneAndDelete(ctx, Model{ID: user.ID}).Decode(&user); err != nil {
		return nil, ErrDeleteUser
	}

	return user, nil
}
