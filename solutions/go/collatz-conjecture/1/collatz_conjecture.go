package collatzconjecture

import (
    "fmt"
)
func CollatzConjecture(n int) (int, error) {
	step := 0

	if n <= 0 {
            return 0, fmt.Errorf("digit is 0 or negative number")
        }
    
    for n != 1 {
        step += 1
        
        if n % 2 == 0 {
            n = n / 2
        } else {
            n = n * 3 + 1
        }
    }
    return step, nil
}
