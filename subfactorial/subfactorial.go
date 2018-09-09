package subfactorial

// SubFactorial counts subfactorial
var SubFactorial int

// Generate finds possible permutations using Heap algorithm.
func Generate(n int, A []int){
    if n == 1 {
		// fmt.Println(A)
		countSubFactorial(A)
	} else {
        for i := 0; i < n - 1; i ++ {
            Generate(n - 1, A)
            if n % 2 == 0 {
				// swap
				t := A[i]
				A[i] = A[n-1]
				A[n-1] = t
			} else {
				// swap
				t := A[0]
				A[0] = A[n-1]
				A[n-1] = t
			}
		}
        Generate(n - 1, A)
	}
}

// countSubFactorial counts subfactorial
func countSubFactorial(A []int){
	same := 0
		for index, elem := range A {
			// check if elem and its index + 1 is same
			// break if same since that won't count for Subfactorial
			if elem == index + 1 {
				same++
				break
			} 
		}
		// if no same index + 1 and elem, count SubFactorial
		if same == 0 {
			SubFactorial++
		}
}
