// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/k0kishima/golang-realworld-example-app/ent/user"
	"github.com/k0kishima/golang-realworld-example-app/ent/userfollow"
)

// UserFollow is the model entity for the UserFollow schema.
type UserFollow struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// FollowerID holds the value of the "follower_id" field.
	FollowerID uuid.UUID `json:"follower_id,omitempty"`
	// FolloweeID holds the value of the "followee_id" field.
	FolloweeID uuid.UUID `json:"followee_id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserFollowQuery when eager-loading is set.
	Edges        UserFollowEdges `json:"edges"`
	selectValues sql.SelectValues
}

// UserFollowEdges holds the relations/edges for other nodes in the graph.
type UserFollowEdges struct {
	// Follower holds the value of the follower edge.
	Follower *User `json:"follower,omitempty"`
	// Followee holds the value of the followee edge.
	Followee *User `json:"followee,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// FollowerOrErr returns the Follower value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserFollowEdges) FollowerOrErr() (*User, error) {
	if e.Follower != nil {
		return e.Follower, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "follower"}
}

// FolloweeOrErr returns the Followee value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserFollowEdges) FolloweeOrErr() (*User, error) {
	if e.Followee != nil {
		return e.Followee, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "followee"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*UserFollow) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case userfollow.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case userfollow.FieldID, userfollow.FieldFollowerID, userfollow.FieldFolloweeID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the UserFollow fields.
func (uf *UserFollow) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case userfollow.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				uf.ID = *value
			}
		case userfollow.FieldFollowerID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field follower_id", values[i])
			} else if value != nil {
				uf.FollowerID = *value
			}
		case userfollow.FieldFolloweeID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field followee_id", values[i])
			} else if value != nil {
				uf.FolloweeID = *value
			}
		case userfollow.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				uf.CreatedAt = value.Time
			}
		default:
			uf.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the UserFollow.
// This includes values selected through modifiers, order, etc.
func (uf *UserFollow) Value(name string) (ent.Value, error) {
	return uf.selectValues.Get(name)
}

// QueryFollower queries the "follower" edge of the UserFollow entity.
func (uf *UserFollow) QueryFollower() *UserQuery {
	return NewUserFollowClient(uf.config).QueryFollower(uf)
}

// QueryFollowee queries the "followee" edge of the UserFollow entity.
func (uf *UserFollow) QueryFollowee() *UserQuery {
	return NewUserFollowClient(uf.config).QueryFollowee(uf)
}

// Update returns a builder for updating this UserFollow.
// Note that you need to call UserFollow.Unwrap() before calling this method if this UserFollow
// was returned from a transaction, and the transaction was committed or rolled back.
func (uf *UserFollow) Update() *UserFollowUpdateOne {
	return NewUserFollowClient(uf.config).UpdateOne(uf)
}

// Unwrap unwraps the UserFollow entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (uf *UserFollow) Unwrap() *UserFollow {
	_tx, ok := uf.config.driver.(*txDriver)
	if !ok {
		panic("ent: UserFollow is not a transactional entity")
	}
	uf.config.driver = _tx.drv
	return uf
}

// String implements the fmt.Stringer.
func (uf *UserFollow) String() string {
	var builder strings.Builder
	builder.WriteString("UserFollow(")
	builder.WriteString(fmt.Sprintf("id=%v, ", uf.ID))
	builder.WriteString("follower_id=")
	builder.WriteString(fmt.Sprintf("%v", uf.FollowerID))
	builder.WriteString(", ")
	builder.WriteString("followee_id=")
	builder.WriteString(fmt.Sprintf("%v", uf.FolloweeID))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(uf.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// UserFollows is a parsable slice of UserFollow.
type UserFollows []*UserFollow
