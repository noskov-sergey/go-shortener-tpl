package shortener

import (
	"context"
	"errors"
	"fmt"
	"sync"
)

type request struct {
	shortURL string
	username string
	err      *error
}

func (s *service) Delete(ctx context.Context, user string, urls []string) error {
	defer ctx.Done()
	var err error

	inputCh := generator(ctx, urls, user)

	channels := s.fanOut(ctx, inputCh)

	checkCh := fanIn(ctx, channels...)

	var goodData []string
	for res := range checkCh {
		if res.err != nil {
			err = errors.Join(err, *res.err)
			continue
		}
		goodData = append(goodData, res.shortURL)
	}

	errMark := s.repo.MarkDelete(goodData)
	if err != nil {
		return errors.Join(err, errMark)
	}

	return err
}

func generator(ctx context.Context, input []string, user string) chan request {
	inputCh := make(chan request)

	go func() {
		defer close(inputCh)

		for _, data := range input {
			select {
			case <-ctx.Done():
				return
			case inputCh <- request{data, user, nil}:
			}
		}
	}()

	return inputCh
}

func (s *service) check(ctx context.Context, inputCh chan request) chan request {
	checkReq := make(chan request)

	go func() {
		defer close(checkReq)

		for data := range inputCh {
			shortener, err := s.repo.GetByID(data.shortURL)
			if err != nil {
				continue
			}

			if shortener.Username != data.username {
				errOwner := errors.New(fmt.Sprintf("%s your are not owner of url %s", data.username, data.shortURL))
				data.err = &errOwner
			}

			select {
			case <-ctx.Done():
				return
			case checkReq <- data:
			}
		}
	}()
	return checkReq
}

func (s *service) fanOut(ctx context.Context, inputCh chan request) []chan request {
	numWorkers := 10
	channels := make([]chan request, numWorkers)

	for i := 0; i < numWorkers; i++ {
		addResultCh := s.check(ctx, inputCh)
		channels[i] = addResultCh
	}

	return channels
}

func fanIn(ctx context.Context, resultChs ...chan request) chan request {
	finalCh := make(chan request)

	var wg sync.WaitGroup

	for _, ch := range resultChs {
		chClosure := ch
		wg.Add(1)

		go func() {
			defer wg.Done()

			for data := range chClosure {
				select {
				case <-ctx.Done():
					return
				case finalCh <- data:
				}
			}
		}()
	}

	go func() {
		wg.Wait()
		close(finalCh)
	}()

	return finalCh
}
