package content

import (
	"context"

	"github.com/dranikpg/dto-mapper"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/pkg/domain"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/pkg/proto/content"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/pkg/proto/genre"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/pkg/proto/person"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/pkg/proto/search"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/pkg/proto/selection"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/grpc/ping"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"
	"google.golang.org/grpc"
)

type Grpc struct {
	personService    person.PersonServiceClient
	selectionService selection.SelectionServiceClient
	searchService    search.SearchServiceClient
	genreService     genre.GenreServiceClient
	contentService   content.ContentServiceClient

	logger logging.Logger
}

func NewGateway(cfg ServiceContentConfig, logger logging.Logger) (*Grpc, error) {
	grpcConn, err := grpc.Dial(cfg.Addr, grpc.WithInsecure())
	if err != nil {
		return &Grpc{}, err
	}

	result := Grpc{}
	result.personService = person.NewPersonServiceClient(grpcConn)
	result.selectionService = selection.NewSelectionServiceClient(grpcConn)
	result.searchService = search.NewSearchServiceClient(grpcConn)
	result.genreService = genre.NewGenreServiceClient(grpcConn)
	result.contentService = content.NewContentServiceClient(grpcConn)

	result.logger = logger

	err = ping.Ping(grpcConn)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (gateway *Grpc) GetAllSelections(ctx context.Context, limit, offset uint32) ([]domain.Selection, error) {
	selections, err := gateway.selectionService.GetAll(ctx, &selection.GetAllCfg{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return nil, err
	}
	var result []domain.Selection

	err = dto.Map(&result, selections.GetSelections())
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (gateway *Grpc) GetSelectionByID(ctx context.Context, id uint64) (domain.Selection, error) {
	selectionDTO, err := gateway.selectionService.GetByID(ctx, &selection.ID{ID: id})
	if err != nil {
		return domain.Selection{}, err
	}

	var result domain.Selection
	err = dto.Map(&result, selectionDTO)
	if err != nil {
		return domain.Selection{}, err
	}

	return result, nil
}

func (gateway *Grpc) GetFilmByContentID(ctx context.Context, ContentID uint64) (domain.Film, error) {
	filmDTO, err := gateway.contentService.GetFilmByContentID(ctx, &content.ContentID{ID: ContentID})
	if err != nil {
		return domain.Film{}, err
	}

	var result domain.Film
	err = dto.Map(&result, filmDTO)
	if err != nil {
		return domain.Film{}, err
	}
	return result, nil
}

func (gateway *Grpc) GetSeriesByContentID(ctx context.Context, ContentID uint64) (domain.Series, error) {
	seriesDTO, err := gateway.contentService.GetSeriesByContentID(ctx, &content.ContentID{ID: ContentID})
	if err != nil {
		return domain.Series{}, err
	}

	var result domain.Series
	err = dto.Map(&result, seriesDTO)
	if err != nil {
		return domain.Series{}, err
	}
	return result, nil
}

func (gateway *Grpc) GetPersonByID(ctx context.Context, id uint64) (domain.Person, error) {
	personDTO, err := gateway.personService.GetByID(ctx, &person.ID{ID: id})
	if err != nil {
		return domain.Person{}, err
	}

	var result domain.Person
	err = dto.Map(&result, personDTO)
	if err != nil {
		return domain.Person{}, err
	}
	return result, nil
}

func (gateway *Grpc) Search(ctx context.Context, query string) (domain.Search, error) {
	searchDTO, err := gateway.searchService.Search(ctx, &search.SearchParams{Query: query})
	if err != nil {
		return domain.Search{}, err
	}

	var result domain.Search
	err = dto.Map(&result, searchDTO)
	if err != nil {
		return domain.Search{}, err
	}
	return result, nil
}

func (gateway *Grpc) GetContentByOptions(ctx context.Context, options domain.ContentFilter) ([]domain.Content, error) {
	contentDTO, err := gateway.genreService.GetContentByOptions(ctx, &genre.Options{
		ID:     options.ID,
		Limit:  options.Limit,
		Offset: options.Offset,
	})
	if err != nil {
		return nil, err
	}
	var result []domain.Content
	err = dto.Map(&result, contentDTO.Content)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (gateway *Grpc) GetContentByContentIDs(ctx context.Context, ContentIDs []uint64) ([]domain.Content, error) {
	contentIDs := &content.ContentIDs{}
	for _, id := range ContentIDs {
		contentIDs.ContentIDs = append(contentIDs.ContentIDs, &content.ContentID{ID: id})
	}
	contentSeq, err := gateway.contentService.GetContentByContentIDs(ctx, contentIDs)
	if err != nil {
		return []domain.Content{}, err
	}
	var contentSlice []domain.Content
	err = dto.Map(&contentSlice, contentSeq.Content)
	if err != nil {
		return []domain.Content{}, err
	}
	return contentSlice, nil
}

func (gateway *Grpc) GetAllGenres(ctx context.Context) ([]domain.Genre, error) {
	genresDTO, err := gateway.genreService.GetAllGenres(ctx, &genre.Nothing{})
	if err != nil {
		return nil, err
	}

	var result []domain.Genre
	err = dto.Map(&result, genresDTO.Genres)
	if err != nil {
		return nil, err
	}
	return result, nil
}
