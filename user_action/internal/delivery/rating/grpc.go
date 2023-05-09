package rating

import (
	"context"

	"github.com/dranikpg/dto-mapper"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/user_action/pkg/domain"
	ratingProto "github.com/go-park-mail-ru/2023_1_ContentDealers/user_action/pkg/proto/rating"
)

type Grpc struct {
	ratingProto.UnimplementedRatingServiceServer
	ratingUseCase RatingUseCase
	logger        logging.Logger
}

func NewGrpc(ratingUseCase RatingUseCase, logger logging.Logger) *Grpc {
	return &Grpc{ratingUseCase: ratingUseCase, logger: logger}
}

func (service *Grpc) DeleteRating(ctx context.Context, rateRequest *ratingProto.Rating) (*ratingProto.Nothing, error) {
	rate := domain.Rating{}
	err := dto.Map(&rate, rateRequest)
	if err != nil {
		service.logger.WithRequestID(ctx).Trace(err)
		return nil, err
	}
	err = service.ratingUseCase.Delete(ctx, rate)
	if err != nil {
		return nil, err
	}
	return &ratingProto.Nothing{}, nil
}

func (service *Grpc) AddRating(ctx context.Context, rateRequest *ratingProto.Rating) (*ratingProto.Nothing, error) {
	rate := domain.Rating{}
	err := dto.Map(&rate, rateRequest)
	if err != nil {
		service.logger.WithRequestID(ctx).Trace(err)
		return nil, err
	}
	err = service.ratingUseCase.Add(ctx, rate)
	if err != nil {
		return nil, err
	}
	return &ratingProto.Nothing{}, nil
}

func (service *Grpc) HasRating(ctx context.Context, rateRequest *ratingProto.Rating) (*ratingProto.HasRate, error) {
	rate := domain.Rating{}
	err := dto.Map(&rate, rateRequest)
	if err != nil {
		service.logger.WithRequestID(ctx).Trace(err)
		return nil, err
	}
	hasRating, err := service.ratingUseCase.Has(ctx, rate)
	if err != nil {
		return nil, err
	}
	hasRateResponse := ratingProto.HasRate{}
	err = dto.Map(&hasRateResponse, hasRating)
	if err != nil {
		service.logger.WithRequestID(ctx).Trace(err)
		return nil, err
	}
	return &hasRateResponse, nil
}

func (service *Grpc) GetRatingByUser(ctx context.Context, rateOptionsRequest *ratingProto.RatingsOptions) (*ratingProto.Ratings, error) {
	rateOptions := domain.RatingsOptions{}
	err := dto.Map(&rateOptions, rateOptionsRequest)
	if err != nil {
		service.logger.WithRequestID(ctx).Trace(err)
		return nil, err
	}

	ratings, err := service.ratingUseCase.GetByUser(ctx, rateOptions)
	if err != nil {
		return nil, err
	}
	ratingsResponse := ratingProto.Ratings{}
	err = dto.Map(&ratingsResponse, ratings)
	if err != nil {
		service.logger.WithRequestID(ctx).Trace(err)
		return nil, err
	}
	return &ratingsResponse, nil
}

func (service *Grpc) GetRatingByContent(ctx context.Context, rateOptionsRequest *ratingProto.RatingsOptions) (*ratingProto.Ratings, error) {
	rateOptions := domain.RatingsOptions{}
	err := dto.Map(&rateOptions, rateOptionsRequest)
	if err != nil {
		service.logger.WithRequestID(ctx).Trace(err)
		return nil, err
	}

	ratings, err := service.ratingUseCase.GetByContent(ctx, rateOptions)

	if err != nil {
		return nil, err
	}

	ratingsResponse := ratingProto.Ratings{}
	err = dto.Map(&ratingsResponse, ratings)

	if err != nil {
		service.logger.WithRequestID(ctx).Trace(err)
		return nil, err
	}

	return &ratingsResponse, nil
}
