package repository

import (
	"github.com/JulesMike/api-er/entity"
	"github.com/JulesMike/api-er/service"
	"github.com/jinzhu/gorm"
)

// User represents the user repository
type User struct {
	db          *gorm.DB
	securitySvc *service.Security
}

// NewUser returns a new User
func NewUser(db *gorm.DB, securitySvc *service.Security) *User {
	return &User{db: db, securitySvc: securitySvc}
}

// Create creates a new user
func (r *User) Create(user *entity.User) (*entity.User, error) {
	if err := r.db.Create(user).Error; err != nil {
		return nil, ErrCreateRecord
	}

	user, err := r.Get(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// Get retrieves a single user
func (r *User) Get(user *entity.User) (*entity.User, error) {
	if err := r.db.Where(user).First(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, ErrRecordNotFound
		}
		return nil, ErrGetRecord
	}

	return user, nil
}

// List retrieves a list of user
func (r *User) List() ([]entity.User, error) {
	users := []entity.User{}

	if err := r.db.Find(&users).Error; err != nil {
		return nil, ErrListRecords
	}

	return users, nil
}

// Update updates a user
func (r *User) Update(user *entity.User, values interface{}) (*entity.User, error) {
	user, err := r.Get(user)
	if err != nil {
		return nil, err
	}

	// Hash the password if not hashed already
	if user.Password != "" && !r.securitySvc.IsPasswordHashed([]byte(user.Password)) {
		if hashedPassword, err := r.securitySvc.HashPassword([]byte(user.Password)); err == nil {
			user.Password = string(hashedPassword)
		}
	}

	if err := r.db.Model(user).Updates(values).Error; err != nil {
		return nil, ErrUpdateRecord
	}

	return user, nil
}

// Delete deletes a user
func (r *User) Delete(user *entity.User) (*entity.User, error) {
	user, err := r.Get(user)
	if err != nil {
		return nil, err
	}

	if err := r.db.Delete(user).Error; err != nil {
		return nil, ErrDeleteRecord
	}

	return user, nil
}
