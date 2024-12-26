// SPDX-License-Identifier: Apache-2.0

package turnify

import "fmt"

// WorkerIsAvailable checks if a worker is available.
func WorkerIsAvailable(worker *Worker, shift WorkShift, workerMax int) bool {
	counterMax := 0
	for _, workShift := range worker.WorkShifts {
		// Check if the worker already has a shift on the same day
		if workShift.Date.Equal(shift.Date) {
			return false
		}

		// Calculate the previous and next days
		prevDay := shift.Date.AddDate(0, 0, -1)
		nextDay := shift.Date.AddDate(0, 0, 1)

		// Check if the worker has a shift on the previous and next days
		if workShift.Date.Equal(prevDay) && workShift.Date.Equal(nextDay) {
			return false
		}
 
		// Check if the worker has a night shift on the previous day and the requested shift is a day shift
		if workShift.Date.Equal(prevDay) && workShift.NightShift && !shift.NightShift {
			return false
		}

		// Check if the worker has a day shift on the next day and the requested shift is a night shift
		if workShift.Date.Equal(nextDay) && !workShift.NightShift && shift.NightShift {
			return false
		}

		if workShift.WorkType == shift.WorkType {
			counterMax++
			if counterMax > workerMax {
				return false
			}
		}
	}
	// If none of the conditions are met, the worker is available
	return true
}

// AllocateWorkers allocates workers to shifts.
func AllocateWorkers(wType WorkType, workers []*Worker, shifts []WorkShift) error {
	availableShifts := []WorkShift{}
	for _, shift := range shifts {
		for range make([]struct{}, shift.TeamSize) {
			availableShifts = append(availableShifts, shift)
		}
	}
	workersNum := len(workers)
	availableShiftsNum := len(availableShifts)
	maxForWorker := ((workersNum + availableShiftsNum - 1) / workersNum)+100
	queue := NewIDQueue(workersNum)
	for _, availableShift := range availableShifts {
		allocated := false
		for j := 0; j < workersNum; j++ {
			worker := queue.Next(wType, workers)
			if WorkerIsAvailable(worker, availableShift, maxForWorker) {
				worker.WorkShifts = append(worker.WorkShifts, availableShift)
				allocated = true
				break
			}
		}
		if !allocated {
			return fmt.Errorf("No available workers for shift %s", availableShift.Description)	
		}		
	}
	return nil
}
