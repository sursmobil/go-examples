package interfaces

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

// Example showing usage interfaces

// Dice is a struct representing set of D'n'D dices like 2k12 means "two dices
// with twelve sides each". Dice is an example of 'jsonification' so it should
// implement json.Marshaler and json.Unmarshaler interfaces
type Dice struct {
	Sides int
	Count int
}

var _ json.Marshaler = Dice{}
var _ json.Unmarshaler = &Dice{}

func invalidFormat(reason string, a ...interface{}) error {
	return fmt.Errorf("Invalid Dice format: "+reason, a)
}

// MarshalJSON implements json.Marshaler for Dice struct
func (d Dice) MarshalJSON() ([]byte, error) {
	s := fmt.Sprintf("\"%vk%v\"", d.Count, d.Sides)
	return []byte(s), nil
}

// UnmarshalJSON implements json.Unmarshaler for Dice struct
func (d *Dice) UnmarshalJSON(in []byte) error {
	s := ""
	err := json.Unmarshal(in, &s)
	if err != nil {
		return fmt.Errorf("Input is not valid json string: %v", string(in))
	}

	parts := strings.Split(s, "k")
	if len(parts) != 2 {
		return invalidFormat("should be '<count>k<sides>' but is %s", s)
	}

	sCount, sSides := parts[0], parts[1]

	count, err := strconv.Atoi(sCount)
	if err != nil {
		return invalidFormat("count should be int but is '%s'", sCount)
	}

	sides, err := strconv.Atoi(sSides)
	if err != nil {
		return invalidFormat("sides should be int but is '%s'", sSides)
	}

	d.Sides = sides
	d.Count = count
	return nil
}
