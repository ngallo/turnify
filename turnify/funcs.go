// SPDX-License-Identifier: Apache-2.0

package turnify

import (
	"fmt"
	"math/rand"
	"time"
)

// GenerateRandomNumbers generates a slice of random numbers.
func GenerateRandomNumbers(x int) []int {
	if x <= 0 {
		return nil
	}
	numbers := make([]int, x)
	for i := 1; i <= x; i++ {
		numbers[i-1] = i
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(numbers), func(i, j int) {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	})
	return numbers
}

// ConvertDayToDayType converts a day to a slice of WorkShift.
func ConvertDayToDayType(day time.Time, specials []WorkShift) []WorkShift {
	var result []WorkShift

	// Specials
	for _, special := range specials {
		if special.Date.Equal(day) {
			result = append(result, special)
		}
	}
	if len(result) > 0 {
		return result
	}

	// Day before specials
	for _, special := range specials {
		if !special.DayBeforeHoliday {
			continue
		}
		dayBefore := special.Date.AddDate(0, 0, -1)
		if day.Equal(dayBefore) {
			result = []WorkShift{
				{Date: day, WorkType: PreHoliday, NightShift: false, Description: fmt.Sprintf("Prefestivo %s", special.Description), DayBeforeHoliday: true},
				{Date: day, WorkType: SuperPreHoliday, NightShift: true, Description: fmt.Sprintf("Prefestivo %s", special.Description), DayBeforeHoliday: true},
			}
			break	
		}
	}
	if len(result) > 0 {
		return result
	}

	// Any other day
	weekDay := day.Weekday()
	switch weekDay {
	case 0:
		return []WorkShift{
			{Date: day, WorkType: HolidayDay, NightShift: false, Description: "Domenica Mattina"},
			{Date: day, WorkType: HolidayNight, NightShift: true, Description: "Domenica Sera"},
		}
	case 6:
		return []WorkShift{
			{Date: day, WorkType: PreHoliday, NightShift: false, Description: "Sabato Mattina"},
			{Date: day, WorkType: SuperPreHoliday, NightShift: true, Description: "Sabato Sera"},
		}
	case 1, 2, 3, 4, 5:
		return []WorkShift{
			{Date: day, WorkType: Regular, NightShift: true, Description: "Settimana Sera"},
		}
	}
	return []WorkShift{}
}

// BuildWeekDays builds a slice of WorkShift.
func BuildWeekDays(firstDay time.Time, lastDay time.Time, specialDays []WorkShift) []WorkShift {
	weekDays := []WorkShift{}
	for d := firstDay; !d.After(lastDay); d = d.AddDate(0, 0, 1) {
		wdays := ConvertDayToDayType(d, specialDays)
		for _, wday := range wdays {
			italianDays := []string{
				"Domenica",  // Sunday
				"Lunedì",    // Monday
				"Martedì",   // Tuesday
				"Mercoledì", // Wednesday
				"Giovedì",   // Thursday
				"Venerdì",   // Friday
				"Sabato",    // Saturday
			}
			wdayItalian := italianDays[d.Weekday()]
			wday.Weekday = wdayItalian
			switch wday.WorkType {
			case Regular:
				wday.TeamSize = 2
			default:
				wday.TeamSize = 3
			}
			weekDays = append(weekDays, wday)
		}
	}
	weekMap := map[WorkType][]WorkShift{}
	for _, wday := range weekDays {
		if _, ok := weekMap[WorkType(wday.WorkType)]; !ok {
			weekMap[wday.WorkType] = []WorkShift{}
		}
		weekMap[wday.WorkType] = append(weekMap[wday.WorkType], wday)
	}
	return weekDays
}

// BuildDoctors builds a slice of doctors.
func BuildDoctors() []*Worker {
	workers := make([]*Worker, 9)
	for i := 0; i < 9; i++ {
		workers[i] = &Worker{
			Name:       fmt.Sprintf("MD%d", i+1),
			WorkShifts: []WorkShift{},
		}
	}
	return workers
}
