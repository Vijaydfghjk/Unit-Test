package calculater

import "testing"

func TestAdd(t *testing.T) {

	tests := []struct {
		a, b int

		expacted int
	}{

		{1, 2, 3},
		{6, 4, 10},
		{-1, -2, -3},
		{0, 0, 0},
		{500, 700, 1200},
	}

	for _, tt := range tests {

		result := Add(tt.a, tt.b)

		if result != tt.expacted {

			t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expacted)
		}
	}

}

/*

func Testadd(t *testing.T) {
	result := Add(2, 3)
	expected := 5

	if result != expected {

		t.Errorf("Add(2, 3) = %d; want %d", result, expected)
	}
}


*/
