package domain

type SearchContent struct {
	Total   uint32
	Content []Content
}

type SearchPerson struct {
	Total   uint32
	Persons []Person
}

type SearchResult struct {
	SearchContent SearchContent
	SearchPerson  SearchPerson
}

type SearchQuery struct {
	Query      string
	TargetSlug string
	Limit      uint32
	Offset     uint32
}
