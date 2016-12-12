package multipleadder

func AddMultiplesOfUntil(numberOne int, numberTwo int, upperBound int) int {
    sum := 0

    for i := 0; i < upperBound; i++ {
        if i % numberOne == 0 || i % numberTwo == 0 {
            sum += i
        }
    }

    return sum
}
