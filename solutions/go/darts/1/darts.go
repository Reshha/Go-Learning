package darts

func Score(x, y float64) int {

    formula := x * x + y * y

    if formula <= 100 && formula > 25 {
        return 1
    } 

	if formula <= 25 && formula > 1 {
        return 5
    } 

    if formula <= 1 {
        return 10
    }

    return 0
}
