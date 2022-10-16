package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"project-3/graph/generated"
	"project-3/graph/model"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.UserInput) (*model.User, error) {
	var user model.User = r.userService.CreateUser(input)
	return &user, nil
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, id int, input model.UserInput) (string, error) {
	// panic(fmt.Errorf("not implemented: UpdateUser - updateUser"))
    err := r.userService.UpdateUser(&input, id)
	if err != nil {
		return "nil", err
	}
	pesan := "data berhasil di update !"
	return pesan, nil

}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, id int) (string, error) {
    err := r.userService.DeleteUser(id)
	if err != nil {
		return "", err
	}
	pesan := "user berhasil dihapus !"
	return pesan, nil
}

// CreateBook is the resolver for the createBook field.
func (r *mutationResolver) CreateBook(ctx context.Context, input model.BookInput) (*model.Book, error) {
	var book model.Book = r.bookService.CreateBook(input)
	return &book, nil
}

// UpdateBook is the resolver for the updateBook field.
func (r *mutationResolver) UpdateBook(ctx context.Context, id int, input model.BookInput) (string, error) {
    err := r.userService.UpdateBook(&input, id)
	if err != nil {
		return "nil", err
	}
	pesan := "data berhasil di update !"
	return pesan, nil
}

// DeleteBook is the resolver for the deleteBook field.
func (r *mutationResolver) DeleteBook(ctx context.Context, id int) (string, error) {
    err := r.userService.DeleteBook(id)
	if err != nil {
		return "", err
	}
	pesan := "user berhasil dihapus !"
	return pesan, nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	// panic(fmt.Errorf("not implemented: Users - users"))
	var user []*model.User = r.userService.GetAllUsers()
	return user, nil
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	user, err := r.userService.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Books is the resolver for the books field.
func (r *queryResolver) Books(ctx context.Context) ([]*model.Book, error) {
	// panic(fmt.Errorf("not implemented: Books - books"))
	var book []*model.Book = r.bookService.GetAllBooks()
	return book, nil
}

// Book is the resolver for the book field.
func (r *queryResolver) Book(ctx context.Context, id string) (*model.Book, error) {
	book, err := r.bookService.GetBookByID(id)
	if err != nil {
		return nil, err
	}
	return &book, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
