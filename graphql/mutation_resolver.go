package main

import (
	"context"
	"errors"
	productModel "github.com/wignn/micro-3/order/model"
	"log"
	"time"
)

var (
	ErrInvalidParameter = errors.New("invalid parameter")
)

type mutationResolver struct {
	server *GraphQLServer
}

func (r *mutationResolver) CreateAccount(c context.Context, in AccountInput) (*Account, error) {
	c, cancel := context.WithTimeout(c, 3*time.Second)
	defer cancel()

	a, err := r.server.accountClient.PostAccount(c, in.Name, in.Email, in.Password)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	account := &Account{
		ID:    a.ID,
		Name:  a.Name,
		Email: a.Email,
	}

	return account, nil
}

func (r *mutationResolver) CreateProduct(c context.Context, in ProductInput) (*Product, error) {
	c, cancel := context.WithTimeout(c, 3*time.Second)

	defer cancel()

	p, err := r.server.catalogClient.PostProduct(c, in.Name, in.Description, in.Price, in.Image)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &Product{
		ID:          p.Id,
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
		Image:       p.Image,
	}, nil
}

func (r *mutationResolver) CreateOrder(ctx context.Context, in OrderInput) (*Order, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var products []*productModel.OrderedProduct
	for _, p := range in.Products {
		if p.Quantity <= 0 {
			return nil, ErrInvalidParameter
		}
		products = append(products, &productModel.OrderedProduct{
			ID:       p.ID,
			Quantity: uint32(p.Quantity),
		})
	}
	o, err := r.server.orderClient.PostOrder(ctx, in.AccountID, products)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var orderProducts []*OrderedProduct
	for _, p := range o.Products {
		productID := p.ID
		prodDetail, err := r.server.catalogClient.GetProduct(ctx, productID)
		if err != nil {
			log.Printf("failed to get product %s: %v", productID, err)
			continue
		}
		orderProducts = append(orderProducts, &OrderedProduct{
			ID:          prodDetail.Id,
			Name:        prodDetail.Name,
			Description: prodDetail.Description,
			Price:       prodDetail.Price,
			Quantity:    int(p.Quantity),
		})
	}

	return &Order{
		ID:         o.ID,
		CreatedAt:  o.CreatedAt,
		TotalPrice: o.TotalPrice,
		Products:   orderProducts,
	}, nil
}

func (r *mutationResolver) CreateReview(ctx context.Context, in ReviewInput) (*Review, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	if in.Rating < 1 || in.Rating > 5 {
		return nil, ErrInvalidParameter
	}
	review, err := r.server.reviewClient.PostReview(ctx, in.ProductID, in.AccountID, *in.Content, int32(in.Rating))

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &Review{
		ID:        review.ID,
		Rating:    int(review.Rating),
		Content:   &review.Content,
		CreatedAt: review.CreatedAt,
	}, nil
}

func (r *mutationResolver) DeleteProduct(c context.Context, id string) (*DeleteProductResponse, error) {
	result, err := r.server.catalogClient.DeleteProduct(c, id)

	if err != nil {
		log.Printf("failed to delete product: %+v\n", err)
		message := "Failed to delete product"
		return &DeleteProductResponse{Success: false, Message: &message, DeletedID: result.DeletedID}, err
	}
	return &DeleteProductResponse{Success: result.Success, Message: &result.Message, DeletedID: result.DeletedID}, nil
}


func (r *mutationResolver) Login(c context.Context, in LoginInput) (*AuthResponse, error) {
	ctx, cancel := context.WithTimeout(c, 3*time.Second)
	defer cancel()

	token, err := r.server.authClient.Login(ctx, in.Email, in.Password)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &AuthResponse{
		ID:           token.Auth.Id,
		Email:        token.Auth.Email,
		BackendToken: &Token{
			AccessToken:  token.Auth.Token.AccessToken,
			RefreshToken: token.Auth.Token.RefreshToken,
			ExpiresIn:    int(token.Auth.Token.ExpiresAt),
		},
	}, nil
}


func (r *mutationResolver) RefreshToken(c context.Context, refreshToken string) (*Token, error) {
	ctx, cancel := context.WithTimeout(c, 3*time.Second)
	defer cancel()

	newToken, err := r.server.authClient.RefreshToken(ctx, refreshToken)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &Token{
		AccessToken:  newToken.AccessToken,
		RefreshToken: newToken.RefreshToken,
		ExpiresIn:    int(newToken.ExpiresAt),
	}, nil
}