package fizzbuzz

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// *****************************
// Tests Expecting to Complete *
// *****************************
func TestBasicInput(t *testing.T) {
	solution := []string{"1", "Fizz", "Buzz", "Fizz", "5", "FizzBuzz", "7", "Fizz", "Buzz", "Fizz", "11", "FizzBuzz", "13", "Fizz", "Buzz", "Fizz", "17", "FizzBuzz", "19", "Fizz"}
	result, err := FizzBuzz(20,2,3)
	assert.Nil(t, err)
	assert.Equal(t, solution, result)
}

func TestFizzAtBuzzAtOutsideTotal(t *testing.T) {
	solution := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	result, err := FizzBuzz(10, 11, 12)
	assert.Nil(t, err)
	assert.Equal(t, solution, result)
}

// **************************
// Tests Expecting to Error *
// **************************

func TestTotalLessThanOne(t *testing.T) {
	result, err := FizzBuzz(0,2,3)
	assert.EqualError(t, err,"ERROR: Total cannot be less than 1" )
	assert.Nil(t, result)
}

func TestFizzAtLessThanOne(t *testing.T) {
	result, err := FizzBuzz(20,0,3)
	assert.EqualError(t, err,"ERROR: fizzAt cannot be less than 1" )
	assert.Nil(t, result)
}

func TestBuzzAtLessThanOne(t *testing.T) {
	result, err := FizzBuzz(20,2,0)
	assert.EqualError(t, err,"ERROR: buzzAt cannot be less than 1" )
	assert.Nil(t, result)
}


