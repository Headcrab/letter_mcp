package letters

import (
	"testing"
)

func TestCountLetters(t *testing.T) {
	counter := NewCounter()

	testCases := []struct {
		name     string
		word     string
		letters  string
		expected map[rune]LetterStat
	}{
		{
			name:    "Простой пример с разным регистром",
			word:    "ПриВЕТ",
			letters: "пр",
			expected: map[rune]LetterStat{
				'п': {Letter: 'п', LowerCount: 0, UpperCount: 1, TotalCount: 1},
				'р': {Letter: 'р', LowerCount: 1, UpperCount: 0, TotalCount: 1},
			},
		},
		{
			name:    "Пример с повторяющимися буквами",
			word:    "ПрограммироВАНИЕ",
			letters: "рае",
			expected: map[rune]LetterStat{
				'р': {Letter: 'р', LowerCount: 3, UpperCount: 0, TotalCount: 3},
				'а': {Letter: 'а', LowerCount: 1, UpperCount: 1, TotalCount: 2},
				'е': {Letter: 'е', LowerCount: 0, UpperCount: 1, TotalCount: 1},
			},
		},
		{
			name:    "Пустое слово",
			word:    "",
			letters: "абв",
			expected: map[rune]LetterStat{
				'а': {Letter: 'а', LowerCount: 0, UpperCount: 0, TotalCount: 0},
				'б': {Letter: 'б', LowerCount: 0, UpperCount: 0, TotalCount: 0},
				'в': {Letter: 'в', LowerCount: 0, UpperCount: 0, TotalCount: 0},
			},
		},
		{
			name:     "Пустой список букв",
			word:     "Пример",
			letters:  "",
			expected: map[rune]LetterStat{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := counter.CountLetters(tc.word, tc.letters)

			// Проверяем совпадение слова
			if result.Word != tc.word {
				t.Errorf("Ожидалось слово '%s', получено '%s'", tc.word, result.Word)
			}

			// Проверяем количество букв
			if len(result.LetterStats) != len(tc.expected) {
				t.Errorf("Ожидалось %d статистик, получено %d", len(tc.expected), len(result.LetterStats))
			}

			// Проверяем статистику для каждой буквы
			for letter, expectedStat := range tc.expected {
				actualStat, ok := result.LetterStats[letter]
				if !ok {
					t.Errorf("Не найдена статистика для буквы '%c'", letter)
					continue
				}

				if actualStat.LowerCount != expectedStat.LowerCount {
					t.Errorf("Для буквы '%c': ожидалось %d строчных, получено %d",
						letter, expectedStat.LowerCount, actualStat.LowerCount)
				}

				if actualStat.UpperCount != expectedStat.UpperCount {
					t.Errorf("Для буквы '%c': ожидалось %d заглавных, получено %d",
						letter, expectedStat.UpperCount, actualStat.UpperCount)
				}

				if actualStat.TotalCount != expectedStat.TotalCount {
					t.Errorf("Для буквы '%c': ожидалось %d всего, получено %d",
						letter, expectedStat.TotalCount, actualStat.TotalCount)
				}
			}
		})
	}
}
