package utils_test

import (
	"testing"

	"github.com/shinya-ac/TodoAPI/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestNormalizeString(t *testing.T) {
	t.Run("全角文字を半角に変換し、小文字にすることを確認", func(t *testing.T) {
		input := "ＡＢＣＤＥａｂｃｄｅ"
		expected := "abcdeabcde"
		result := utils.NormalizeString(input)
		assert.Equal(t, expected, result)
	})
}
