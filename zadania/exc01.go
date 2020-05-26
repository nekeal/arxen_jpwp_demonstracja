package main

import (
	"github.com/pkg/errors"
	"math"
	"sync"
)

func isPrime(x uint64) bool {
	for i := uint64(2); i <= uint64(math.Floor(math.Sqrt(float64(x)))); i++ {
		if x%i == 0 {
			return false
		}
	}
	return x > 1
}

// PrimesCounter returns number of primes in range form start to limit
func PrimesCounter(start uint64, limit uint64) (uint64, error) {

	// validate the range
	if start >= limit {
		return 0, errors.New("start must be less than the limit")
	}

	// number of primes in slice [start, limit]
	numberOfPrimes := uint64(0)

	// iterate until start is greater or equal limit
	for start < limit {
		if isPrime(start) {
			numberOfPrimes += 1
		}
		start += 1
	}

	return numberOfPrimes, nil
}

func PrimesCounterWorker(start uint64, limit uint64, result chan uint64, err chan error, wg *sync.WaitGroup) {
	defer wg.Done()

	// validate the range
	if start >= limit {
		err <- errors.New("start must be less than the limit")
		return
	}

	// number of primes in slice [start, limit]
	numberOfPrimes := uint64(0)

	// iterate until start is greater or equal limit
	for start < limit {
		if isPrime(start) {
			numberOfPrimes += 1
		}
		start += 1
	}

	result <- numberOfPrimes
}

func PrimesCounterHandler(start uint64, limit uint64) (uint64, error) {

	// validate the range
	if start >= limit {
		return 0, errors.New("start must be less than the limit")
	}

	// number of primes in slice [start, limit]
	numberOfPrimes := uint64(0)

	numOfWorkers := 0

	wg := sync.WaitGroup{}
	results := make(chan uint64, numOfWorkers)
	errs := make(chan error, numOfWorkers)

	workersResolution := uint64(500)

	//for i := 1; i <= numOfWorkers; i++ {
	//	st := uint64(uint64(i)*(limit-start)/uint64(numOfWorkers) + start)
	//	var end uint64
	//	if end = uint64(uint64(i+1)*(limit-start)/uint64(numOfWorkers) + limit); end > limit {
	//		limit = end
	//	}
	//	wg.Add(1)
	//	go PrimesCounterWorker(st, end, results, errs, &wg)
	//}

	for st := start; st < limit; st += workersResolution {
		numOfWorkers += 1
		var end uint64
		if st+workersResolution > limit {
			end = limit
		} else {
			end = st +workersResolution
		}
		wg.Add(1)
		go PrimesCounterWorker(st, end, results, errs, &wg)
	}

	wg.Wait()

	for {
		select {
		case err := <-errs:
			return 0, err
		case res := <-results:
			numberOfPrimes += res
			numOfWorkers--
			if numOfWorkers == 0 {
				return numberOfPrimes, nil
			}
		}
	}
}
