package functions

import (
	"time"
)

func IsSlotInFuture(slot int) bool {
	currentSlot := int(time.Now().Unix() / 12) // Assuming 12 seconds per slot
	return slot > currentSlot
}
