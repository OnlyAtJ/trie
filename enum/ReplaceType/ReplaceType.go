package ReplaceType

type Enum uint8

const (
	Holder = iota + 1 // 使用占位符替换
	Del               // 删除敏感词
)
