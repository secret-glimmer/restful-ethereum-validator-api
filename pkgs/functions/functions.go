package functions

import (
	"math/big"
	"strconv"
	t "time"
)

func IsSlotInFuture(slot int) bool {
	firstDate := t.Date(2020, 12, 1, 12, 0, 0, 0, t.UTC)
	currentSlot := int(t.Since(firstDate)/t.Second) / 12
	return slot > currentSlot
}

func TimeFromSlot(slot int) t.Time {
	firstTime := t.Date(2020, 12, 1, 12, 0, 23, 0, t.UTC)
	date := firstTime.Add(t.Second * 12 * t.Duration(slot))
	return date
}

func SlotFromTime(time t.Time) int {
	firstTime := t.Date(2020, 12, 1, 12, 0, 23, 0, t.UTC)
	slot := time.Sub(firstTime) / t.Second / 12

	return int(slot)
}

func FloatFromHex(hex string) *big.Float {
	if len(hex) < 2 {
		return nil
	}
	intValue := new(big.Int)
	intValue.SetString(hex[2:], 16)

	floatValue := new(big.Float).SetInt(intValue)
	floatValue.Quo(floatValue, big.NewFloat(1e9))
	return floatValue
}

func IntFromHex(hex string) int64 {
	if len(hex) < 2 {
		return 0
	}
	value, err := strconv.ParseInt(hex[2:], 16, 64)
	if err != nil {
		return 0
	}
	return value
}
