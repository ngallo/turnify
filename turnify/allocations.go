// SPDX-License-Identifier: Apache-2.0

package turnify

import "fmt"

// WorkerIsAvailable checks if a worker is available.
func WorkerIsAvailable(worker *Worker, shift WorkShift) bool {
	for _, workShift := range worker.WorkShifts {
		// Check if the worker already has a shift on the same day
		if workShift.Date.Equal(shift.Date) {
			return false
		}

		// Calculate the previous and next days
		prevDay := shift.Date.AddDate(0, 0, -1)
		nextDay := shift.Date.AddDate(0, 0, 1)

		// Check if the worker has a night shift on the previous day and the requested shift is a day shift
		if workShift.Date.Equal(prevDay) && workShift.NightShift && !shift.NightShift {
			return false
		}

		// Check if the worker has a day shift on the next day and the requested shift is a night shift
		if workShift.Date.Equal(nextDay) && !workShift.NightShift && shift.NightShift {
			return false
		}
	}

	// If none of the conditions are met, the worker is available
	return true
}


// AllocateWorkers allocates workers to shifts.
func AllocateWorkers(workers []*Worker, shifts []WorkShift) error {
	notAllocated := []WorkShift{}
	for _, shift := range shifts {
		randomNumber := GenerateRandomNumbers(len(workers))
		for _, i := range randomNumber {
			worker := workers[i]
			if WorkerIsAvailable(worker, shift) {
				worker.WorkShifts = append(worker.WorkShifts, shift)
				break
			}
			if i == len(workers) - 1 {
				notAllocated = append(notAllocated, shift)
			}
		}
	}	
	if len(notAllocated) > 0 {
		return fmt.Errorf("Not all shifts were allocated")
	}
	return nil
}
