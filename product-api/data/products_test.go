package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "skeaks",
		Price: 1337.00,
		SKU:   "nike-jords-one",
	}
	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
