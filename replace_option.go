package trie

import (
	"github.com/OnlyAtJ/trie/enum/ReplaceType"
)

type ReplaceOption interface {
	Apply(opt *ReplaceSetting)
}

type ReplaceSetting struct {
	ReplaceType ReplaceType.Enum
	PlaceHolder rune
}

type ReplaceOptionFunc func(v *ReplaceSetting)

func (fn ReplaceOptionFunc) Apply(opt *ReplaceSetting) {
	fn(opt)
}

// 替换敏感词
func WithPlaceHolder(char rune) ReplaceOption {
	return ReplaceOptionFunc(func(c *ReplaceSetting) {
		c.ReplaceType = ReplaceType.Holder
		c.PlaceHolder = char
	})
}

// 删除敏感词
func WithDelete() ReplaceOption {
	return ReplaceOptionFunc(func(c *ReplaceSetting) {
		c.ReplaceType = ReplaceType.Del
	})
}
