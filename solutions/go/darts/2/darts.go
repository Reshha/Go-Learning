package darts

func Score(x, y float64) int {

    formula := x * x + y * y

    if formula > 100 {
        return 0
    }
    
    if formula > 25 {
        return 1
    } 

	if formula > 1 {
        return 5
    } 

    return 10
}
