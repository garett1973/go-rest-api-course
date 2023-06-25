package comment

import (
	"context"
	"fmt"
)

// Comment represents a comment on an article
type Comment struct {
	ID     string
	Slug   string
	Body   string
	Author string
}

type Store interface {
	GetComment(context.Context, string) (Comment, error)
}

// Service represents the comment service. All the logic for the comment service is contained within this struct.
type Service struct {
	Store Store
}

// NewService creates a new comment service, returns a pointer to a new instance of the service
func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

func (s *Service) GetComments(ctx context.Context, id string) (Comment, error) {
	fmt.Println("Retrieving comments...")
	cmt, err := s.Store.GetComment(ctx, id)

	if err != nil {
		fmt.Println(err)
		return Comment{}, err
	}

	return cmt, nil
}
