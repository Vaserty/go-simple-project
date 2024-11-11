package services

import (
	"fmt"
	"go-simple-project/internal/common/dependencies"
	"go-simple-project/internal/infrastructure"
	"math"
	"sort"

	logger "github.com/sirupsen/logrus"
)

type SearchValueService struct {
	deps       *dependencies.Dependency
	repository infrastructure.IFileNumberRepository
	numbers    []int
}

func (s *SearchValueService) setNumbers(numbers []int) error {
	if !sort.IntsAreSorted(numbers) {
		return fmt.Errorf(
			"you cannot load these numbers because they are unsorted",
		)
	}
	s.numbers = numbers
	return nil
}

func (s *SearchValueService) getTolerance(value int) float64 {
	toleranceAsNumber := s.deps.Config.Searcher.MaxToleranceDiffPercent / 100
	tolerance := toleranceAsNumber * float64(value)
	logger.Debugf("Tolerance for value '%v': %v", value, tolerance)
	return tolerance
}

func (s *SearchValueService) findClosestIndex(index int, value int) int {
	closestIndex := -1
	tolerance := s.getTolerance(value)
	minDifference := math.MaxFloat64

	for _, idx := range []int{index - 1, index} {
		if idx >= 0 && idx < len(s.numbers) {
			difference := math.Abs(float64(s.numbers[idx] - value))
			if difference <= tolerance && difference < minDifference {
				closestIndex = idx
				minDifference = difference
			}
		}
	}
	return closestIndex
}

func (s *SearchValueService) SearchIndex(value int) *SearchResultDto {
	index := sort.Search(len(s.numbers), func(idx int) bool {
		return s.numbers[idx] >= value
	})

	if index < len(s.numbers) && s.numbers[index] == value {
		logger.Debugf("Value '%v' exactly found. Index: %v", value, index)
		return newFoundSearchResult(index, value)
	}

	if closestIndex := s.findClosestIndex(index, value); closestIndex != -1 {
		searchedValue := s.numbers[closestIndex]
		logger.Debugf(
			"Closest value found: %v for value %v. Index: %v",
			searchedValue,
			value,
			closestIndex,
		)
		return newFoundSearchResult(closestIndex, searchedValue)
	}

	return newNotFoundSearchResult()
}

func (s *SearchValueService) LoadValues() error {
	numbers, err := s.repository.Get()
	if err != nil {
		return fmt.Errorf("error during download number from file: %v", err)
	}
	s.setNumbers(numbers)
	return nil
}
