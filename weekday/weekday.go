package main

func weekday(day int) string {
	const invalidDay = "Некорректный день недели"
	calendar := map[int]string{
		1: "Понедельник",
		2: "Вторник",
		3: "Среда",
		4: "Четверг",
		5: "Пятница-Развратница",
		6: "Суббота",
		7: "Воскресенье",
	}
	dayweek, ok := calendar[day]
	if ok {
		return dayweek
	}
	return invalidDay
}

// 	switch day {
// 	case 1:
// 		return "Понедельник"
// 	case 2:
// 		return "Вторник"
// 	case 3:
// 		return "Среда"
// 	case 4:
// 		return "Четверг"
// 	case 5:
// 		return "Пятница"
// 	case 6:
// 		return "Суббота"
// 	case 7:
// 		return "Воскресенье"
// 	default:
// 		return invalidDay
// 	}
// }
