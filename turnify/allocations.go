// SPDX-License-Identifier: Apache-2.0

package turnify

func AllocateWorkers(workers []*Worker, shifts []WorkShift) {
	for _, shift := range shifts {
		for _, worker := range workers {
			if worker.WorkShifts == nil {
				worker.WorkShifts = []WorkShift{}
			}
			worker.WorkShifts = append(worker.WorkShifts, shift)
		}
	}	
}
