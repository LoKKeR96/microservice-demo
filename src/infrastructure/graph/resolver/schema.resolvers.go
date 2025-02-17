package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.64

import (
	"context"

	"github.com/lokker96/microservice_example/application/query"
	"github.com/lokker96/microservice_example/infrastructure/graph"
	"github.com/lokker96/microservice_example/infrastructure/graph/model"
	"github.com/palantir/stacktrace"
)

// Login is the resolver for the Login field.
func (r *mutationResolver) Login(ctx context.Context, username string, password string) (string, error) {
	userAuthenticationService := r.C.GetUserAuthenticationService()

	token, err := userAuthenticationService.Authenticate(username, password)

	if err != nil {
		return "", err
	}

	return token, nil
}

// Messages is the resolver for the messages field.
func (r *queryResolver) Messages(ctx context.Context) ([]*model.Message, error) {
	getAllMessagesQuery := r.C.GetMessagesByFilterQuery(ctx)

	messageFilters := query.GetMessagesByFilterRequest{}

	messages, err := getAllMessagesQuery.Do(ctx, messageFilters)
	if err != nil {
		return nil, stacktrace.Propagate(err, "error on retrieving messages")
	}

	response := make([]*model.Message, 0)

	for _, message := range messages {
		// de-referencing and appending as repository returns pointers to handle nil returns
		response = append(response, NewMessageResponse(*message))
	}

	return response, nil
}

// Mutation returns graph.MutationResolver implementation.
func (r *Resolver) Mutation() graph.MutationResolver { return &mutationResolver{r} }

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
