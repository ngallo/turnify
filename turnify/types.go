// SPDX-License-Identifier: Apache-2.0

package turnify

import "time"

type WorkType string

const (
	Regular         WorkType = "FERIALE-NOTTE"
	PreHoliday      WorkType = "PRE-FESTIVO-MATTINA"
	SuperPreHoliday WorkType = "PRE-FESTIVO-SERA"
	HolidayDay      WorkType = "FESTIVO-MATTINA"
	HolidayNight    WorkType = "FESTIVO-SERA"
	Special         WorkType = "FESTA-SPECIALE"
	SuperSpecial    WorkType = "FESTA-SPECIALE-SUPER"
)

// WorkShift represents a work shift.
type WorkShift struct {
	Date             time.Time
	Weekday          string
	WorkType         WorkType
	NightShift       bool
	DayBeforeHoliday bool
	TeamSize		 int
	Description      string
}

// Worker represents a worker.
type Worker struct {
	Name string
	WorkShifts []WorkShift
}