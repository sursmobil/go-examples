package interfaces

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDiceMarshaling(t *testing.T) {
	dice := &Dice{Sides: 12, Count: 2}
	data, err := json.Marshal(dice)
	require.Nil(t, err, "Dice cannot be marshaled")
	require.Equal(t, "\"2k12\"", string(data), "Marshaled Dice have improper format")
}

func TestDiceUnmarshaling(t *testing.T) {
	data := []byte("\"5k7\"")
	dice := &Dice{}
	err := json.Unmarshal(data, dice)
	require.Nil(t, err, "Dice cannot be unmarshaled")
	require.Equal(t, 5, dice.Count, "Unmarshaled Dice have wrong count")
	require.Equal(t, 7, dice.Sides, "Unmarshaled Dice have wrong sides")
}

func TestDiceUnmarashalInvalidJson(t *testing.T) {
	data := []byte("5k7")
	dice := &Dice{}
	err := json.Unmarshal(data, dice)
	require.NotNil(t, err, "Dice unmarshaling should fail: invalid json string")
}

func TestDiceUnmarashalInvalidFormat(t *testing.T) {
	data := []byte("\"aaa\"")
	dice := &Dice{}
	err := json.Unmarshal(data, dice)
	require.NotNil(t, err, "Dice unmarshaling should fail: invalid dice format")
}

func TestDiceUnmarashalCountNotNumber(t *testing.T) {
	data := []byte("\"ak5\"")
	dice := &Dice{}
	err := json.Unmarshal(data, dice)
	require.NotNil(t, err, "Dice unmarshaling should fail: count not number")
}

func TestDiceUnmarashalSidesNotNumber(t *testing.T) {
	data := []byte("\"5kb\"")
	dice := &Dice{}
	err := json.Unmarshal(data, dice)
	require.NotNil(t, err, "Dice unmarshaling should fail: sides not number")
}
