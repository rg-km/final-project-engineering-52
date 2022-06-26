package usecase

import (
	"context"
	"go-scholarship/api/models"
	"log"
	"time"

	"golang.org/x/sync/errgroup"
)

type scholarUseCase struct {
	scholarRepo  models.ScholarshipRepository
	userRepo     models.UserRepository
	categoryRepo models.CategoryRepository
	timeout      time.Duration
}

func NewScholarshipUseCase(scholarRepo models.ScholarshipRepository, userRepo models.UserRepository, categoryRepo models.CategoryRepository, timeout time.Duration) models.ScholarshipUseCase {
	return &scholarUseCase{
		scholarRepo:  scholarRepo,
		userRepo:     userRepo,
		categoryRepo: categoryRepo,
		timeout:      timeout,
	}
}

// fill user details at scholarship
func (s *scholarUseCase) fillUserDetails(ctx context.Context, scholars []models.ScholarResponse) ([]models.ScholarResponse, error) {
	g, ctx := errgroup.WithContext(ctx)

	users := map[int64]models.UserResponse{}

	for _, scholar := range scholars {
		users[scholar.User.ID] = models.UserResponse{}
	}

	userChan := make(chan models.UserResponse)
	for id := range users {
		id := id
		g.Go(func() error {
			user, err := s.userRepo.FetchById(ctx, id)
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

	for i, s := range scholars {
		if u, ok := users[s.User.ID]; ok {
			scholars[i].User = u
		}
	}

	return scholars, nil
}

// fill category details at scholarship
func (s *scholarUseCase) fillCategoryDetails(ctx context.Context, scholars []models.ScholarResponse) ([]models.ScholarResponse, error) {
	g, ctx := errgroup.WithContext(ctx)

	categories := map[int64]models.Category{}

	for _, scholar := range scholars {
		categories[scholar.Category.ID] = models.Category{}
	}

	categoryChan := make(chan models.Category)
	for id := range categories {
		id := id
		g.Go(func() error {
			category, err := s.categoryRepo.FetchById(ctx, id)
			if err != nil {
				return err
			}

			categoryChan <- category
			return nil
		})
	}

	go func() {
		if err := g.Wait(); err != nil {
			log.Println(err)
			return
		}

		close(categoryChan)
	}()

	for category := range categoryChan {
		if category != (models.Category{}) {
			categories[category.ID] = category
		}
	}

	if err := g.Wait(); err != nil {
		return nil, err
	}

	for i, s := range scholars {
		if c, ok := categories[s.Category.ID]; ok {
			scholars[i].Category = c
		}
	}

	return scholars, nil
}

// fetch all scholarships
func (s *scholarUseCase) Fetch(ctx context.Context) ([]models.ScholarResponse, error) {
	c, cancel := context.WithTimeout(ctx, s.timeout)

	defer cancel()

	scholars, err := s.scholarRepo.Fetch(c)
	if err != nil {
		return nil, err
	}

	scholars, err = s.fillCategoryDetails(c, scholars)
	if err != nil {
		return nil, err
	}

	scholars, err = s.fillUserDetails(c, scholars)
	if err != nil {
		return nil, err
	}

	return scholars, nil
}

// fetch by id scholarship
func (s *scholarUseCase) FetchById(ctx context.Context, id int64) (models.ScholarResponse, error) {
	c, cancel := context.WithTimeout(ctx, s.timeout)

	defer cancel()

	scholar, err := s.scholarRepo.FetchById(c, id)
	if err != nil {
		return models.ScholarResponse{}, err
	}

	category, err := s.categoryRepo.FetchById(c, scholar.Category.ID)
	if err != nil {
		return models.ScholarResponse{}, err
	}

	user, err := s.userRepo.FetchById(c, scholar.User.ID)
	if err != nil {
		return models.ScholarResponse{}, err
	}

	scholar.Category = category // fill category at scholar
	scholar.User = user         // fill user at scholar

	return scholar, nil
}

// create scholarship
func (s *scholarUseCase) Create(ctx context.Context, scholarReq *models.ScholarRequest) (models.ScholarResponse, error) {
	c, cancel := context.WithTimeout(ctx, s.timeout)

	defer cancel()

	scholarResp, err := s.scholarRepo.Create(c, scholarReq)
	if err != nil {
		return scholarResp, err
	}

	category, err := s.categoryRepo.FetchById(c, scholarReq.CategoryID)
	if err != nil {
		return scholarResp, err
	}

	user, err := s.userRepo.FetchById(c, scholarReq.UserID)
	if err != nil {
		return scholarResp, err
	}

	scholarResp.Category = category // fill user at scholar
	scholarResp.User = user // fill user at scholar

	return scholarResp, nil
}

// update scholarship
func (s *scholarUseCase) Update(ctx context.Context, id int64, scholarReq *models.ScholarRequest) (models.ScholarResponse, error) {
	c, cancel := context.WithTimeout(ctx, s.timeout)

	defer cancel()

	scholarResp, err := s.scholarRepo.Update(c, id, scholarReq)
	if err != nil {
		return scholarResp, err
	}

	category, err := s.categoryRepo.FetchById(c, scholarReq.CategoryID)
	if err != nil {
		return scholarResp, err
	}

	user, err := s.userRepo.FetchById(c, scholarReq.UserID)
	if err != nil {
		return scholarResp, err
	}

	scholarResp.Category = category // fill user at scholar
	scholarResp.User = user // fill user at scholar

	return scholarResp, nil
}

// delete scholarship
func (s *scholarUseCase) Delete(ctx context.Context, id int64) error {
	c, cancel := context.WithTimeout(ctx, s.timeout)

	defer cancel()

	return s.scholarRepo.Delete(c, id)
}
