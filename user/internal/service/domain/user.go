package domain

import (
	"fmt"
	"time"

	"github.com/rs/xid"
)

type User struct {
	ID       string
	UserID   string
	Username string
	Email    string
	Timestamps
	IsAdmin bool
}

func NewUser(email, username string, isAdmin bool) *User {
	t := &Timestamps{}

	return &User{
		UserID:     fmt.Sprintf("user_%s", xid.New().String()),
		Email:      email,
		Username:   username,
		IsAdmin:    isAdmin,
		Timestamps: t.CreatedNow(),
	}
}

func (u *User) SetEmail(email string) error {
	if email == "" {
		return EmailEmptyError
	}

	// some further checks

	u.Email = email
	u.Updated = time.Now()

	return nil
}

type Timestamps struct {
	Created time.Time
	Updated time.Time
	Deleted time.Time
}

// CreatedNow initializes the Timestamps for a newly created entity.
func (t Timestamps) CreatedNow() Timestamps {
	return Timestamps{
		Created: time.Now(),
		Updated: time.Now(),
		Deleted: time.Time{},
	}
}
