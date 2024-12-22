// SPDX-License-Identifier: Apache-2.0

package turnify

// IDQueue represents a queue of IDs.
type IDQueue struct {
	cursor int
	ids []int
}

// NewIDQueue creates a new ID queue.
func NewIDQueue(size int) *IDQueue {
	return &IDQueue{
		cursor: -1,
		ids: GenerateRandomNumbers(size),
	}
}

// Next returns the next ID in the queue.
func (q *IDQueue) Next() int {
	q.cursor++
	if q.cursor >= len(q.ids) {
		q.cursor = 0
		q.ids = GenerateRandomNumbers(len(q.ids))
	}
	return q.ids[q.cursor]
}