package data

import "testing"

func TestCheckValidation(t *testing.T) {
	p := &Product{
		Name:  "coffee",
		Price: 1.00,
		SKU:   "asb-asb-asb",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
