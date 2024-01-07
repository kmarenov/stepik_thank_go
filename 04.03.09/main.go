package main

import (
	"errors"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

// начало решения

// Task описывает задачу, выполненную в определенный день
type Task struct {
	Date  time.Time
	Dur   time.Duration
	Title string
}

// ParsePage разбирает страницу журнала
// и возвращает задачи, выполненные за день
func ParsePage(src string) ([]Task, error) {
	lines := strings.Split(src, "\n")
	date, err := parseDate(lines[0])
	if err != nil {
		return nil, err
	}
	tasks, err := parseTasks(date, lines[1:])
	if err != nil {
		return nil, err
	}
	sortTasks(tasks)
	return tasks, nil
}

// parseDate разбирает дату в формате дд.мм.гггг
func parseDate(src string) (time.Time, error) {
	t, err := time.Parse("02.01.2006", src)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}

// parseTasks разбирает задачи из записей журнала
func parseTasks(date time.Time, lines []string) ([]Task, error) {
	re := regexp.MustCompile(`(\d+:\d+) - (\d+:\d+) (.+)`)
	tasksMap := make(map[string]time.Duration, len(lines))
	for _, line := range lines {
		groups := re.FindStringSubmatch(line)
		if len(groups) != 4 {
			return nil, errors.New("wrong format")
		}
		h1, m1, err := parseTime(groups[1])
		if err != nil {
			return nil, err
		}
		start := time.Date(date.Year(), date.Month(), date.Day(), h1, m1, 0, 0, date.Location())
		h2, m2, err := parseTime(groups[2])
		if err != nil {
			return nil, err
		}
		end := time.Date(date.Year(), date.Month(), date.Day(), h2, m2, 0, 0, date.Location())
		title := groups[3]
		dur := end.Sub(start)
		if dur <= 0 {
			return nil, errors.New("wrong duration")
		}

		_, ok := tasksMap[title]
		if !ok {
			tasksMap[title] = 0
		}
		tasksMap[title] += dur
	}
	tasks := make([]Task, 0, len(tasksMap))
	for title, d := range tasksMap {
		tasks = append(tasks, Task{
			Date:  date,
			Dur:   d,
			Title: title,
		})
	}
	return tasks, nil
}

func parseTime(str string) (hour int, minute int, err error) {
	times := strings.Split(str, ":")
	if len(times) != 2 {
		return 0, 0, errors.New("wrong format")
	}
	h, err := strconv.Atoi(times[0])
	if err != nil {
		return 0, 0, err
	}
	if h < 0 || h > 23 {
		return 0, 0, errors.New("wrong hour")
	}
	m, err := strconv.Atoi(times[1])
	if err != nil {
		return 0, 0, err
	}
	if m < 0 || m > 59 {
		return 0, 0, errors.New("wrong minute")
	}
	return h, m, nil
}

// sortTasks упорядочивает задачи по убыванию длительности
func sortTasks(tasks []Task) {
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].Dur > tasks[j].Dur
	})
}

// конец решения
// ::footer

func main() {
	page := `15.04.2022
8:00 - 8:30 Завтрак
8:30 - 9:30 Оглаживание кота
9:30 - 10:00 Интернеты
10:00 - 14:00 Напряженная работа
14:00 - 14:45 Обед
14:45 - 15:00 Оглаживание кота
15:00 - 19:00 Напряженная работа
19:00 - 19:30 Интернеты
19:30 - 22:30 Безудержное веселье
22:30 - 23:00 Оглаживание кота`

	entries, err := ParsePage(page)
	if err != nil {
		panic(err)
	}
	fmt.Println("Мои достижения за", entries[0].Date.Format("2006-01-02"))
	for _, entry := range entries {
		fmt.Printf("- %v: %v\n", entry.Title, entry.Dur)
	}

	// ожидаемый результат
	/*
		Мои достижения за 2022-04-15
		- Напряженная работа: 8h0m0s
		- Безудержное веселье: 3h0m0s
		- Оглаживание кота: 1h45m0s
		- Интернеты: 1h0m0s
		- Обед: 45m0s
		- Завтрак: 30m0s
	*/
}
