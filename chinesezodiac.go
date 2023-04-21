package gozodiac

import "time"

// ChineseZodiacSign represents the chinese zodiac signs
type ChineseZodiacSign int

func (s ChineseZodiacSign) String() string {
	switch s {
	case Monkey:
		return "Monkey"
	case Rooster:
		return "Rooster"
	case Dog:
		return "Dog"
	case Pig:
		return "Pig"
	case Rat:
		return "Rat"
	case Ox:
		return "Ox"
	case Tiger:
		return "Tiger"
	case Rabbit:
		return "Rabbit"
	case Dragon:
		return "Dragon"
	case Snake:
		return "Snake"
	case Horse:
		return "Horse"
	case Goat:
		return "Goat"
	case Error:
		return "Error"
	}
	return ""
}

const (
	// Monkey chinese zodiac sign
	Monkey ChineseZodiacSign = iota
	// Rooster chinese zodiac sign
	Rooster
	// Dog chinese zodiac sign
	Dog
	// Pig chinese zodiac sign
	Pig
	// Rat chinese zodiac sign
	Rat
	// Ox chinese zodiac sign
	Ox
	// Tiger chinese zodiac sign
	Tiger
	// Rabbit chinese zodiac sign
	Rabbit
	// Dragon chinese zodiac sign
	Dragon
	// Snake chinese zodiac sign
	Snake
	// Horse chinese zodiac sign
	Horse
	// Goat chinese zodiac sign
	Goat
	// Error is an error-case chinese zodiac sign
	Error
)

// GetChineseZodiacSign returns the matching chinese zodiac sign
func GetChineseZodiacSign(date time.Time) ChineseZodiacSign {
	cZodiac := int(date.Year() - ((date.Year() / 12) * 12))

	switch cZodiac {
	case 0:
		return Monkey

	case 1:
		return Rooster

	case 2:
		return Dog

	case 3:
		return Pig

	case 4:
		return Rat

	case 5:
		return Ox

	case 6:
		return Tiger

	case 7:
		return Rabbit

	case 8:
		return Dragon

	case 9:
		return Snake

	case 10:
		return Horse

	case 11:
		return Goat
	}

	return Error // never reach here..
}
func GetChineseZodiacSignLocale(date time.Time, iso639_1 string) string {
	if iso639_1 == "en" {
		return GetChineseZodiacSign(date).String()
	}
	type ss []string
	icon := ss{
		"🐒",
		"🐓",
		"🐕",
		"🐖",
		"🐀",
		"🐂",
		"🐅",
		"🐇",
		"🐉",
		"🐍",
		"🐎",
		"🐐",
	}
	ru := ss{
		"Обезьяна",
		"Петух",
		"Собака",
		"Свинья",
		"Крыса",
		"Бык",
		"Тигр",
		"Кролик",
		"Дракон",
		"Змея",
		"Лошадь",
		"Коза",
	}
	msSs := map[string]ss{
		"ru": ru,
	}
	names, ok := msSs[iso639_1]
	if !ok {
		names = icon
	}
	return names[int(GetChineseZodiacSign(date))]
}
