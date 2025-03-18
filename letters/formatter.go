package letters

import (
	"fmt"
	"strings"
)

// Formatter определяет интерфейс для форматирования результатов подсчета букв
type Formatter interface {
	// Format форматирует результаты подсчета букв в строку
	Format(counts LetterCounts) string
}

// TextFormatter - реализация форматирования результатов в текстовом виде
type TextFormatter struct{}

// NewTextFormatter создает новый форматтер результатов в текстовом виде
func NewTextFormatter() Formatter {
	return &TextFormatter{}
}

// Format форматирует результаты подсчета букв в строку
func (f *TextFormatter) Format(counts LetterCounts) string {
	result := fmt.Sprintf("Результаты подсчёта в слове '%s':\n", counts.Word)

	for _, letter := range getSortedLetters(counts) {
		stat := counts.LetterStats[letter]
		lowerLetter := strings.ToLower(string(letter))
		upperLetter := strings.ToUpper(string(letter))

		result += fmt.Sprintf("'%s' (строчная): %d\n", lowerLetter, stat.LowerCount)
		result += fmt.Sprintf("'%s' (заглавная): %d\n", upperLetter, stat.UpperCount)
		result += fmt.Sprintf("'%s' (всего): %d\n\n", string(letter), stat.TotalCount)
	}

	return result
}

// getSortedLetters возвращает буквы в отсортированном порядке для стабильного вывода
func getSortedLetters(counts LetterCounts) []rune {
	var letters []rune
	for letter := range counts.LetterStats {
		letters = append(letters, letter)
	}

	// Простая сортировка для стабильного вывода
	for i := 0; i < len(letters); i++ {
		for j := i + 1; j < len(letters); j++ {
			if letters[i] > letters[j] {
				letters[i], letters[j] = letters[j], letters[i]
			}
		}
	}

	return letters
}
