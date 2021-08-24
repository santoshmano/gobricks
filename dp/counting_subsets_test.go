package dp

import "testing"

func TestNumSubsetsSizekInt(t *testing.T) {

	if num_subsets_size_k_int(5, 2) != 10 {
		t.Errorf(`Error`)
	} else {
		t.Log(`test success`)
	}

	if num_subsets_size_k_rec(5, 2) != 10 {
		t.Errorf(`Error`)
	} else {
		t.Log(`test success`)
	}
}
