package main

import (
	"fmt"
	"time"

	"github.com/ngallo/turnify.git/turnify"
)


func main() {
	startMonth := 2
	startYear := 2025
	endMonth := 1
	endYear := 2026

	specialDays := []turnify.WorkShift{
		{Date: time.Date(2025, 4, 20, 0, 0, 0, 0, time.UTC), WorkType: turnify.Special, NightShift: false, DayBeforeHoliday: true, Description: "Pasqua"},
		{Date: time.Date(2025, 4, 20, 0, 0, 0, 0, time.UTC), WorkType: turnify.Special, NightShift: true, DayBeforeHoliday: true, Description: "Pasqua"},

		{Date: time.Date(2025, 4, 21, 0, 0, 0, 0, time.UTC), WorkType: turnify.Special, NightShift: false, DayBeforeHoliday: true, Description: "Pasquetta"},
		{Date: time.Date(2025, 4, 21, 0, 0, 0, 0, time.UTC), WorkType: turnify.Special, NightShift: true, DayBeforeHoliday: true, Description: "Pasquetta"},

		{Date: time.Date(2025, 4, 25, 0, 0, 0, 0, time.UTC), WorkType: turnify.Special, NightShift: false, DayBeforeHoliday: true, Description: "Festa della Liberazione"},
		{Date: time.Date(2025, 4, 25, 0, 0, 0, 0, time.UTC), WorkType: turnify.Special, NightShift: true, DayBeforeHoliday: true, Description: "Festa della Liberazione"},
		
		{Date: time.Date(2025, 5, 8, 0, 0, 0, 0, time.UTC), WorkType: turnify.Special, NightShift: false, DayBeforeHoliday: false, Description: "Festa Bari"},
		{Date: time.Date(2025, 5, 8, 0, 0, 0, 0, time.UTC), WorkType: turnify.Special, NightShift: true, DayBeforeHoliday: false, Description: "Festa Bari"},

		{Date: time.Date(2025, 5, 1, 0, 0, 0, 0, time.UTC), WorkType: turnify.Special, NightShift: false, DayBeforeHoliday: true, Description: "Festa dei lavoratori"},
		{Date: time.Date(2025, 5, 1, 0, 0, 0, 0, time.UTC), WorkType: turnify.Special, NightShift: true, DayBeforeHoliday: true, Description: "Festa dei lavoratori"},

		{Date: time.Date(2025, 6, 2, 0, 0, 0, 0, time.UTC), WorkType: turnify.Special, NightShift: false, DayBeforeHoliday: true, Description: "Festa della Repubblica"},
		{Date: time.Date(2025, 6, 2, 0, 0, 0, 0, time.UTC), WorkType: turnify.Special, NightShift: true, DayBeforeHoliday: true, Description: "Festa della Repubblica"},

		{Date: time.Date(2025, 8, 15, 0, 0, 0, 0, time.UTC), WorkType: turnify.Special, NightShift: false, DayBeforeHoliday: true, Description: "Ferragosto"},
		{Date: time.Date(2025, 8, 15, 0, 0, 0, 0, time.UTC), WorkType: turnify.Special, NightShift: true, DayBeforeHoliday: true, Description: "Ferragosto"},

		{Date: time.Date(2025, 11, 1, 0, 0, 0, 0, time.UTC), WorkType: turnify.Special, NightShift: false, DayBeforeHoliday: true, Description: "Tutti i Santi"},
		{Date: time.Date(2025, 11, 1, 0, 0, 0, 0, time.UTC), WorkType: turnify.Special, NightShift: true, DayBeforeHoliday: true, Description: "Tutti i Santi"},

		{Date: time.Date(2025, 12, 8, 0, 0, 0, 0, time.UTC), WorkType: turnify.SuperSpecial, NightShift: false, DayBeforeHoliday: true, Description: "Immacolata Concezione"},
		{Date: time.Date(2025, 12, 8, 0, 0, 0, 0, time.UTC), WorkType: turnify.Special, NightShift: true, DayBeforeHoliday: true, Description: "Immacolata Concezione"},

		{Date: time.Date(2025, 12, 25, 0, 0, 0, 0, time.UTC), WorkType: turnify.SuperSpecial, NightShift: false, DayBeforeHoliday: true, Description: "Natale"},
		{Date: time.Date(2025, 12, 25, 0, 0, 0, 0, time.UTC), WorkType: turnify.Special, NightShift: true, DayBeforeHoliday: true, Description: "Natale"},

		{Date: time.Date(2025, 12, 26, 0, 0, 0, 0, time.UTC), WorkType: turnify.SuperSpecial, NightShift: false, DayBeforeHoliday: true, Description: "Santo Stefano"},
		{Date: time.Date(2025, 12, 26, 0, 0, 0, 0, time.UTC), WorkType: turnify.Special, NightShift: true, DayBeforeHoliday: true, Description: "Santo Stefano"},
		
		{Date: time.Date(2025, 12, 31, 0, 0, 0, 0, time.UTC), WorkType: turnify.Special, NightShift: false, DayBeforeHoliday: true, Description: "San Silvestro"},
		{Date: time.Date(2025, 12, 31, 0, 0, 0, 0, time.UTC), WorkType: turnify.SuperSpecial, NightShift: true, DayBeforeHoliday: true, Description: "San Silvestro"},
	}

	firstDay := time.Date(startYear, time.Month(startMonth), 1, 0, 0, 0, 0, time.UTC)
	lastDay := time.Date(endYear, time.Month(endMonth), 31, 0, 0, 0, 0, time.UTC)

	weekDays := turnify.BuildWeekDays(firstDay, lastDay, specialDays)
	for _, weekDay := range weekDays {
		fmt.Printf("%s,%s,%s,%s\n", weekDay.Date.Format("02/01/2006"), weekDay.Weekday, weekDay.WorkType, weekDay.Description)
	}

	workers := turnify.BuildDoctors()
	for _, woorker := range workers {
		fmt.Printf("%s\n", woorker.Name)
	}
}
