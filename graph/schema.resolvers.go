package graph

import (
	"context"
	"fmt"

	"github.com/shakilahmmeed/gqlgen-todos/graph/generated"
	"github.com/shakilahmmeed/gqlgen-todos/graph/model"
	"github.com/shakilahmmeed/gqlgen-todos/models"
)

func (r *linkResolver) ID(ctx context.Context, obj *models.Link) (string, error) {
	return "1", nil
}

func (r *linkResolver) Address(ctx context.Context, obj *models.Link) (string, error) {
	return "Dhaka", nil
}

func (r *mutationResolver) CreateLink(ctx context.Context, input model.NewLink) (*models.Link, error) {
	link2 := models.Link{
		Title:   input.Title,
		Address: input.Address,
	}

	r.DB.Create(&link2)
	return &link2, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Links(ctx context.Context) ([]*models.Link, error) {
	var links []*models.Link
	r.DB.Preload("Links").Find(&links)

	return links, nil
}

// Link returns generated.LinkResolver implementation.
func (r *Resolver) Link() generated.LinkResolver { return &linkResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type linkResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
