package user_action

import (
	"context"

	"github.com/dranikpg/dto-mapper"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/user_action/pkg/domain"
	rateProto "github.com/go-park-mail-ru/2023_1_ContentDealers/user_action/pkg/proto/rating"
	"google.golang.org/grpc/metadata"
)

func (gate *Gateway) DeleteRating(ctx context.Context, rating domain.Rating) (domain.Rating, error) {
	md := metadata.Pairs(
		"requestID", ctx.Value("requestID").(string),
	)
	ctx = metadata.NewOutgoingContext(ctx, md)
	rateRequest := rateProto.Rating{}
	err := dto.Map(&rateRequest, rating)
	if err != nil {
		gate.logger.WithRequestID(ctx).Trace(err)
		return rating, err
	}
	ratingResponse, err := gate.ratingManager.DeleteRating(ctx, &rateRequest)
	if err != nil {
		gate.logger.WithRequestID(ctx).Trace(err)
		return rating, err
	}
	err = dto.Map(&rating, ratingResponse)
	if err != nil {
		gate.logger.WithRequestID(ctx).Trace(err)
		return rating, err
	}
	return rating, nil
}

func (gate *Gateway) AddRating(ctx context.Context, rating domain.Rating) error {
	md := metadata.Pairs(
		"requestID", ctx.Value("requestID").(string),
	)
	ctx = metadata.NewOutgoingContext(ctx, md)
	rateRequest := rateProto.Rating{}
	err := dto.Map(&rateRequest, rating)
	if err != nil {
		gate.logger.WithRequestID(ctx).Trace(err)
		return err
	}
	_, err = gate.ratingManager.AddRating(ctx, &rateRequest)
	if err != nil {
		gate.logger.WithRequestID(ctx).Trace(err)
		return err
	}
	return nil
}

func (gate *Gateway) HasRating(ctx context.Context, rating domain.Rating) (domain.HasRating, error) {
	md := metadata.Pairs(
		"requestID", ctx.Value("requestID").(string),
	)
	ctx = metadata.NewOutgoingContext(ctx, md)
	rateRequest := rateProto.Rating{}
	err := dto.Map(&rateRequest, rating)
	if err != nil {
		gate.logger.WithRequestID(ctx).Trace(err)
		return domain.HasRating{}, err
	}
	hasRate, err := gate.ratingManager.HasRating(ctx, &rateRequest)
	if err != nil {
		gate.logger.WithRequestID(ctx).Trace(err)
		return domain.HasRating{}, err
	}
	hasRateResponse := domain.HasRating{}
	err = dto.Map(&hasRateResponse, hasRate)
	if err != nil {
		gate.logger.WithRequestID(ctx).Trace(err)
		return domain.HasRating{}, err
	}
	return hasRateResponse, nil
}

func (gate *Gateway) GetRatingByUser(ctx context.Context, options domain.RatingsOptions) (domain.Ratings, error) {
	md := metadata.Pairs(
		"requestID", ctx.Value("requestID").(string),
	)
	ctx = metadata.NewOutgoingContext(ctx, md)

	rateOptionsRequest := rateProto.RatingsOptions{}
	err := dto.Map(&rateOptionsRequest, options)

	if err != nil {
		gate.logger.WithRequestID(ctx).Trace(err)
		return domain.Ratings{}, err
	}
	ratesResponse, err := gate.ratingManager.GetRatingByUser(ctx, &rateOptionsRequest)
	if err != nil {
		gate.logger.WithRequestID(ctx).Trace(err)
		return domain.Ratings{}, err
	}
	rates := domain.Ratings{}
	err = dto.Map(&rates, ratesResponse)
	if err != nil {
		gate.logger.WithRequestID(ctx).Trace(err)
		return domain.Ratings{}, err
	}
	return rates, nil
}

func (gate *Gateway) GetRatingByContent(ctx context.Context, options domain.RatingsOptions) (domain.Ratings, error) {
	md := metadata.Pairs(
		"requestID", ctx.Value("requestID").(string),
	)
	ctx = metadata.NewOutgoingContext(ctx, md)

	rateOptionsRequest := rateProto.RatingsOptions{}
	err := dto.Map(&rateOptionsRequest, options)

	if err != nil {
		gate.logger.WithRequestID(ctx).Trace(err)
		return domain.Ratings{}, err
	}
	ratesResponse, err := gate.ratingManager.GetRatingByContent(ctx, &rateOptionsRequest)
	if err != nil {
		gate.logger.WithRequestID(ctx).Trace(err)
		return domain.Ratings{}, err
	}
	rates := domain.Ratings{}
	err = dto.Map(&rates, ratesResponse)
	if err != nil {
		gate.logger.WithRequestID(ctx).Trace(err)
		return domain.Ratings{}, err
	}
	return rates, nil
}
