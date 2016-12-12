package multipleadder

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func Test_AddMultiplesOfUntil_AddsCorrectly(t *testing.T) {
    Assert_AddsCorrectly(t, 3, 5, 10, 23)
    Assert_AddsCorrectly(t, 3, 5, 1000, 233168)
}

func Assert_AddsCorrectly(t *testing.T, numberOne int, numberTwo int, upperBound int, expectedResult int) {
    result := AddMultiplesOfUntil(numberOne, numberTwo, upperBound)

    assert.Equal(t, result, expectedResult, "Should be " + string(expectedResult))
}
