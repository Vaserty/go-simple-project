package infrastructure

import (
	"bufio"
	"fmt"
	"go-simple-project/internal/common/dependencies"
	"os"
	"strconv"
)

type FileNumberRepository struct {
	deps *dependencies.Dependency
}

func (f *FileNumberRepository) lineContentGenerator() (<-chan string, <-chan error) {

	lines := make(chan string)
	errs := make(chan error, 1)

	go func() {
		defer close(lines)
		defer close(errs)

		file, err := os.Open(string(f.deps.Config.NumbersFile.SourcePath))

		if err != nil {
			errs <- err
			return
		}

		defer file.Close()
		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			lines <- scanner.Text()
		}

		if err := scanner.Err(); err != nil {
			errs <- err
		}
	}()

	return lines, errs
}

func (f *FileNumberRepository) Get() ([]int, error) {
	lines, errs := f.lineContentGenerator()
	var numbers []int

	for line := range lines {
		number, conv_err := strconv.Atoi(line)

		if conv_err != nil {
			return numbers, fmt.Errorf(
				"error occured during convert line '%s' value to int: %v",
				line,
				conv_err,
			)
		}
		numbers = append(numbers, number)
	}

	if err := <-errs; err != nil {
		return numbers, fmt.Errorf(
			"error occured during read file with path '%s'. %v",
			f.deps.Config.NumbersFile.SourcePath,
			err,
		)
	}

	return numbers, nil
}
