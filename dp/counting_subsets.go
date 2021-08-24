package dp

import (
	"math/big"
)

// Given n elements how many subsets of size k items
// == how many ways can you choose k items from n
// == n choose k == C(n,k) == nCk == (n|k)

// approach #1: combinations formula
// C(n,k) = n!/k!(n-k)!
// T(n) = O(n)
// S(n) = O(1) but this approach needs calculation of factorials
// 				and it would lead to integer/float overflows
//				that needs a Big Integer library, and it is most definitely
//				much lower in performance compared to regular int/floats
//
// C(n,0) == C(n,n) == 1, subset(n, 0) = {}, an empty subset
// C(n,1) == C(n,n-1) == n
// C(n,2) == C(n,n-2) == n(n-1)/2 ??

//          | 1, k==0 || k==n
//          | n, k==1 || k==n-1 , we dont need this as below condition will take care of it recursively
// C(n,k) = | C(n-1, k) + C(n-1, k-1)
//

func num_subsets_size_k_rec(n, k int) int {
	if (k == 0) || (k == n) {
		return 1
	}
	return num_subsets_size_k_rec(n-1, k) + num_subsets_size_k_rec(n-1, k-1)
}

func num_subsets_size_k_int(n, k int) int {

	var factBig func(n *big.Int) *big.Int

	factBig = func(n *big.Int) *big.Int {
		zero := big.NewInt(0)
		one := big.NewInt(1)
		n :=
		if n.Cmp(zero) == 0 {
			return one
		}
		return n.Mul(n, factBig(n.Sub(n, one)))
	}

	nBig := big.NewInt(int64(n))
	kBig := big.NewInt(int64(k))
	nMinuskBig := big.NewInt(int64(n-k))

	nFact := factBig(nBig)
	kFact := factBig(kBig)

	comb := nFact.Div(nFact, kFact.Mul(kFact, factBig(nMinuskBig)))

	if comb.IsInt64() {
		return int(comb.Int64())
	}

	return -1
}


func num_subsets_size_k_bigint(n, k int) int {

	// create a n+1 x k+1 array, n+1 rows, k+1 cols
	comb := make([][]int, n+1)
	for row := range(comb) {
		comb[row] = make([]int, k+1)
	}

	return 0
}
