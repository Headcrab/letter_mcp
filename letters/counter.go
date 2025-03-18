package letters

import (
	"strings"
)

// Counter определяет интерфейс для подсчета букв в слове
type Counter interface {
	// CountLetters подсчитывает указанные буквы в слове и возвращает результаты
	CountLetters(word string, letters string) LetterCounts
}

// LetterCounts содержит результаты подсчета букв
type LetterCounts struct {
	Word        string
	LetterStats map[rune]LetterStat
}

// LetterStat содержит статистику для конкретной буквы
type LetterStat struct {
	Letter     rune
	LowerCount int
	UpperCount int
	TotalCount int
}

// DefaultCounter - стандартная реализация счетчика букв
type DefaultCounter struct{}

// NewCounter создает новый экземпляр счетчика букв
func NewCounter() Counter {
	return &DefaultCounter{}
}

// CountLetters подсчитывает указанные буквы в слове
func (c *DefaultCounter) CountLetters(word string, letters string) LetterCounts {
	results := LetterCounts{
		Word:        word,
		LetterStats: make(map[rune]LetterStat),
	}

	for _, letter := range letters {
		// Получаем строчную и прописную версии буквы
		lowerLetter := strings.ToLower(string(letter))
		upperLetter := strings.ToUpper(string(letter))

		// Инициализируем счетчики
		lowerCount := 0
		upperCount := 0

		// Проходим по каждому символу в слове
		for _, ch := range word {
			chStr := string(ch)
			// Проверяем, совпадает ли символ со строчной версией буквы
			if chStr == lowerLetter {
				lowerCount++
			}
			// Проверяем, совпадает ли символ с прописной версией буквы
			if chStr == upperLetter {
				upperCount++
			}
		}

		// Общее количество
		totalCount := lowerCount + upperCount

		results.LetterStats[letter] = LetterStat{
			Letter:     letter,
			LowerCount: lowerCount,
			UpperCount: upperCount,
			TotalCount: totalCount,
		}
	}

	return results
}
