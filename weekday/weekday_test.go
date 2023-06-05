package main

import (
	"testing"
)

func TestWeekday(t *testing.T) {
	testCases := []struct {
		day  int
		want string
	}{
		{0, "Некорректный день недели"},
		{1, "Понедельник"},
		{2, "Вторник"},
		{3, "Среда"},
		{4, "Четверг"},
		{5, "Пятница-Развратница"},
		{6, "Суббота"},
		{7, "Воскресенье"},
	}

	for _, tc := range testCases {
		got := weekday(tc.day)
		if got != tc.want {
			t.Errorf("Weekday(%d) = %s; want %s", tc.day, got, tc.want)
		}
	}
}
