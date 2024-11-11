package services

type ISearchValueService interface {
	LoadValues() error
	SearchIndex(value int) *SearchResultDto
}
