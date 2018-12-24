package median

import "testing"

func TestGetMedian(t *testing.T) {
	cases := []struct {
		in, want []int
	}{
		{[]int{90, 92, 93, 88, 95, 88, 97, 87, 98}, []int{93}},
		{[]int{12, 15, 26}, []int{15}},
		{[]int{12, 15, 20, 26, 32}, []int{20}},
	}
	for _, c := range cases {
		got := getMedian(c.in)
		if got == nil {
			t.Errorf("getMedian(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
