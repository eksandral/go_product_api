package price

import (
	"bytes"
	"strconv"
)

type DiscountPercentage uint

type Price struct {
	Original           uint               `json:"original"`
	Final              uint               `json:"final"`
	DiscountPercentage DiscountPercentage `json:"discount_percentage"`
	Currency           string             `json:"currency"`
}

func (self DiscountPercentage) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	if self == 0 {
		buf.WriteString(`null`)
	} else {
		buf.WriteString(strconv.Itoa(int(self))) // add double quation mark as json format required
	}
	return buf.Bytes(), nil
}
