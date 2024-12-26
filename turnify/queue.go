// SPDX-License-Identifier: Apache-2.0

package turnify

import (
	"math/rand"
	"sort"
	"time"
)

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

func getSortedIDsByWorkShifts(workers []*Worker, workType WorkType) []int {
	counts := make(map[int]int)
	for id, worker := range workers {
		for _, shift := range worker.WorkShifts {
			if shift.WorkType == workType {
				counts[id]++
			}
		}
	}
	ids := make([]int, 0, len(counts))
	for id := range counts {
		ids = append(ids, id)
	}

	rand.Seed(time.Now().UnixNano())

	sort.Slice(ids, func(i, j int) bool {
		if counts[ids[i]] != counts[ids[j]] {
			return counts[ids[i]] > counts[ids[j]]
		}
		return rand.Intn(2) == 0
	})

	return ids
}

// Next returns the next ID in the queue.
func (q *IDQueue) Next(wType WorkType, workers []*Worker) *Worker {
	q.cursor++
	if q.cursor >= len(q.ids) {
		q.cursor = 0
		ids := getSortedIDsByWorkShifts(workers, wType)
		//q.ids = GenerateRandomNumbers(len(q.ids))
		q.ids = ids
	}
	return workers[q.cursor]
}