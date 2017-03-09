package interfaces

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

// Example showing usage inf interfaces

// Dice is a struct representing set of D'n'D dices like 2k12 means "tow dices
// with twelve sides each". Dice is an example of 'jsonification' so it should
// implement json.Marshaler and json.Unmarshaler interfaces
type Dice struct {
	Sides int
	Count int
}

var _ json.Marshaler = Dice{}
var _ json.Unmarshaler = &Dice{}

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
		return fmt.Errorf("Invalid Dice format: should be '<count>k<sides>' but is %s", s)
	}

	sCount, sSides := parts[0], parts[1]

	count, err := strconv.Atoi(sCount)
	if err != nil {
		return fmt.Errorf("Invalid Dice format: count should be int but is '%s'", sCount)
	}

	sides, err := strconv.Atoi(sSides)
	if err != nil {
		return fmt.Errorf("Invalid Dice format: sides should be int but is '%s'", sSides)
	}

	d.Sides = sides
	d.Count = count
	return nil
}
