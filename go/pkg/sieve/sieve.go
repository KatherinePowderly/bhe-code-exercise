package sieve

import (
	"fmt"
	"log"
)

const minimumPrime int64 = 2

type Sieve interface {
	GetNthPrime(n int64) (int64, error)
}

// Prime struct implements the Sieve interface.
type Prime struct{}

func NewSieve() Sieve {
	return &Prime{}
}

// GetNthPrime returns the nth prime number i.e. 2 as the 0th prime number.
func (prime *Prime) GetNthPrime(n int64) (int64, error) {
	scope := n
	primeList := []int64{}

	if n < 0 {
		return 0, fmt.Errorf("invalid number: %d, negative numbers not allowed", n)
	} else if n < minimumPrime {
		scope = n + 1

		log.Printf("%d is less than minimum prime: %d, using %d as scope", n, minimumPrime, scope)
	}

	// Scope is used to control the range of prime numbers to generate.
	// If the list is not long enough, triple the scope to generate more prime numbers
	// until the list is large enough to retrieve the nth prime number.
	for int64(len(primeList)) <= n {
		scope *= 3
		primeList = prime.generateList(scope)
	}

	return primeList[n], nil
}

// generateList generates a list of prime numbers up to the limit.
func (prime *Prime) generateList(limit int64) []int64 {
	// Initialize slice of booleans to indicate numbers as prime.
	isPrime := make([]bool, limit+1)
	for i := range isPrime {
		isPrime[i] = true
	}

	// Sieve of Eratosthenes
	// Begin at the minimum prime number and mark all multiples of prime 
	// numbers as not prime.
	for i := minimumPrime; i <= limit; i++ {
		if isPrime[i] {
			for j := i * i; j <= limit; j += i {
				isPrime[j] = false
			}
		}
	}

	// Assign prime numbers to list.
	primeList := []int64{}
	for i := minimumPrime; i <= limit; i++ {
		if isPrime[i] {
			primeList = append(primeList, i)
		}
	}

	return primeList
}
