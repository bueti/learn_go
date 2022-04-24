package comment

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrFetchingComment = errors.New("failed fetching comment by id")
	ErrNotImplemented  = errors.New("not implemented")
)

// a representation of the comment strcture for our service
type Comment struct {
	ID     string
	Slug   string
	Body   string
	Author string
}

// Store - this interface defines all of the methods that our service needs
// in order to operate
type Store interface {
	GetComment(context.Context, string) (Comment, error)
}

// is the struct on which all our logic will be built on to of
type Service struct {
	Store Store
}

// returns a pointer to a new service
func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	fmt.Println("Retrieving a comment")

	cmt, err := s.Store.GetComment(ctx, id)
	if err != nil {
		// log detailed error
		fmt.Println(err)
		// return only a "meaningful" error to the user
		return Comment{}, ErrFetchingComment
	}

	return cmt, nil
}

func (s *Service) UpdateComment(ctx context.Context, cmt Comment) error {
	return ErrNotImplemented
}

func (s *Service) DeleteComment(ctx context.Context, id string) error {
	return ErrNotImplemented
}

func (s *Service) CreateComment(ctx context.Context, cmt Comment) (Comment, error) {
	return Comment{}, ErrNotImplemented
}
