package usecase

import (
	"context"
	"go-scholarship/api/models"
	"log"
	"time"

	"golang.org/x/sync/errgroup"
)

type commentUseCase struct {
	commentRepo models.CommentRepository
	userRepo    models.UserRepository
	timeout     time.Duration
}

func NewCommentUseCase(commentRepo models.CommentRepository, userRepo models.UserRepository, timeout time.Duration) models.CommentUseCase {
	return &commentUseCase{
		commentRepo: commentRepo,
		userRepo:    userRepo,
		timeout:     timeout,
	}
}

// fill user details at comment
func (co *commentUseCase) fillUserDetails(ctx context.Context, comments []models.CommentResponse) ([]models.CommentResponse, error) {
	g, ctx := errgroup.WithContext(ctx)

	users := map[int64]models.UserResponse{}

	for _, scholar := range comments {
		users[scholar.User.ID] = models.UserResponse{}
	}

	userChan := make(chan models.UserResponse)
	for id := range users {
		id := id
		g.Go(func() error {
			user, err := co.userRepo.FetchById(ctx, id)
			if err != nil {
				return err
			}

			userChan <- user
			return nil
		})
	}

	go func() {
		if err := g.Wait(); err != nil {
			log.Println(err)
			return
		}

		close(userChan)
	}()

	for user := range userChan {
		if user != (models.UserResponse{}) {
			users[user.ID] = user
		}
	}

	if err := g.Wait(); err != nil {
		return nil, err
	}

	for i, s := range comments {
		if u, ok := users[s.User.ID]; ok {
			comments[i].User = u
		}
	}

	return comments, nil
}

// fetch comments
func (co *commentUseCase) Fetch(ctx context.Context) ([]models.CommentResponse, error) {
	c, cancel := context.WithTimeout(ctx, co.timeout)

	defer cancel()

	comments, err := co.commentRepo.Fetch(c)
	if err != nil {
		return nil, err
	}

	comments, err = co.fillUserDetails(c, comments)
	if err != nil {
		return nil, err
	}

	return comments, nil
}

// fetch by id comment
func (co *commentUseCase) FetchById(ctx context.Context, id int64) (models.CommentResponse, error) {
	c, cancel := context.WithTimeout(ctx, co.timeout)

	defer cancel()

	commentResp, err := co.commentRepo.FetchById(c, id)
	if err != nil {
		return models.CommentResponse{}, err
	}

	user, err := co.userRepo.FetchById(c, commentResp.User.ID)
	if err != nil {
		return models.CommentResponse{}, err
	}

	commentResp.User = user

	return commentResp, nil
}

// create comment
func (co *commentUseCase) Create(ctx context.Context, commentReq *models.CommentRequest) (models.CommentResponse, error) {
	c, cancel := context.WithTimeout(ctx, co.timeout)

	defer cancel()

	commentResp, err := co.commentRepo.Create(c, commentReq)
	if err != nil {
		return models.CommentResponse{}, err
	}

	user, err := co.userRepo.FetchById(c, commentResp.User.ID)
	if err != nil {
		return models.CommentResponse{}, err
	}

	commentResp.User = user

	return commentResp, nil
}

// update comment
func (co *commentUseCase) Update(ctx context.Context, id int64, commentReq *models.CommentRequest) (models.CommentResponse, error) {
	c, cancel := context.WithTimeout(ctx, co.timeout)

	defer cancel()

	commentResp, err := co.commentRepo.Update(c, id, commentReq)
	if err != nil {
		return models.CommentResponse{}, err
	}

	user, err := co.userRepo.FetchById(c, commentResp.User.ID)
	if err != nil {
		return models.CommentResponse{}, err
	}

	commentResp.User = user

	return commentResp, nil
}

// delete comment
func (co *commentUseCase) Delete(ctx context.Context, id int64) error {
	c, cancel := context.WithTimeout(ctx, co.timeout)

	defer cancel()

	return co.commentRepo.Delete(c, id)
}
