package ru_test

import (
	"testing"
	"time"

	"github.com/olebedev/when"
	"github.com/olebedev/when/rules"
	"github.com/olebedev/when/rules/ru"
)

func TestWeekday(t *testing.T) {
	// current is Friday
	fixt := []Fixture{
		// past/last
		{"это нужно было сделать в прошлый Понедельник", 45, "прошлый Понедельник", -(2 * 24 * time.Hour)},
		{"прошлая суббота", 0, "прошлая суббота", -(4 * 24 * time.Hour)},
		{"прошлая пятница", 0, "прошлая пятница", -(5 * 24 * time.Hour)},
		{"в последнюю среду", 3, "последнюю среду", -(7 * 24 * time.Hour)},
		{"в прошлый вторник", 3, "прошлый вторник", -(24 * time.Hour)},

		// next
		{"в следующий вторник", 3, "следующий вторник", 6 * 24 * time.Hour},
		{"напиши мне в следующую среду, договоримся", 23, "следующую среду", 7 * 24 * time.Hour},
		{"следующая суббота", 0, "следующая суббота", 3 * 24 * time.Hour},
		{"в следующую суббота", 3, "следующую суббота", 3 * 24 * time.Hour},

		// this
		{"в этот вторник", 3, "этот вторник", -(24 * time.Hour)},
		{"напиши мне в эту среду, договоримся", 23, "эту среду", 0},
		{"эта суббота", 0, "эта суббота", 3 * 24 * time.Hour},
	}

	w := when.New(nil)

	w.Add(ru.Weekday(rules.OverWrite))

	ApplyFixtures(t, "ru.Weekday", w, fixt)
}