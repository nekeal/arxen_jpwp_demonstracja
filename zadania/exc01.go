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
	// TODO if function is implemented delete line below
	panic("implement me")

	// validate the range
	if start >= limit {
		return 0, errors.New("start must be less than the limit")
	}

	// TODO declare variable storing number of primes in range

	numOfWorkers := 0

	// sync group is optional
	wg := sync.WaitGroup{}

	errs := make(chan error, numOfWorkers)
	// TODO create new channel used by PrimesCounterWorker to return results (similar to errs defined above)

	// used for defining (not directly) how many goroutines will be generated
	workersResolution := uint64(500)

	// create numOfWorkers goroutines counting primes
	for st := start; st < limit; st += workersResolution {
		// TODO change value of numOfWorkers

		// end of range for particular goroutine
		var end uint64
		if st+workersResolution > limit {
			end = limit
		} else {
			end = st + workersResolution
		}
		wg.Add(1)
		// TODO run PrimesCounterWorker as goroutine
		println(end)
	}

	// listen for every new element in each channel and handle it
	// if any error occurs return error
	// if every worker is done (number of results == numberOfWorkers) return answer
	for {
		select {
		case err := <-errs:
			return 0, err
			// TODO add next case, that will handle results as follows:
			// for each result add its value to some sum variable
			// if it number of results exceeds number of goroutines return sum
		}
	}
}
