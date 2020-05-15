package trie

import (
	"errors"
	"github.com/OnlyAtJ/trie/enum/ReplaceType"
)

type Node struct {
	Children NodeChildren
	End      bool // 是否是单词的结束，如打人受伤  打(true) 人(true) 伤(true)，可以有三个词汇（打/打人/打人受伤）
}

func NewTrieNode() *Node {
	return &Node{
		Children: make(NodeChildren),
		End:      false,
	}
}

type NodeChildren map[rune]*Node

func NewTrie() *Trie {
	return &Trie{
		Root: NewTrieNode(),
	}
}

type Trie struct {
	Root *Node
}

// 添加词库内容
func (t *Trie) Add(keyword string) error {
	if keyword == "" {
		return ErrKeywordCanNotBeEmpty
	}

	runeChars := []rune(keyword)

	node := t.Root
	for i := 0; i < len(runeChars); i++ {
		runeChar := runeChars[i]

		if _, ok := node.Children[runeChar]; !ok {
			// 不存在，初始化子节点的map
			node.Children[runeChar] = NewTrieNode()
		}
		node = node.Children[runeChar] // 迭代
	}
	node.End = true // 叶子节点

	return nil
}

// 递归删除字符串
func (t *Trie) remove(pNode *Node, runeChars []rune, index int) {
	charsLen := len(runeChars)
	if index < charsLen {
		char := runeChars[index]
		if node, ok := pNode.Children[char]; ok {
			if index == charsLen-1 { // 达到词汇的最后一个开始进行删除
				// 判断是否是节点的最后一个
				if len(node.Children) > 0 {
					// 非最后一个的，删除该词汇的闭环
					node.End = false
				} else {
					// 如果非闭环，则为全词删除
					delete(pNode.Children, runeChars[index])
				}
			} else {
				t.remove(node, runeChars, index+1)
				if !node.End && len(node.Children) == 0 {
					// 不是叶子节点并且没有子节点了，删除
					delete(pNode.Children, runeChars[index])
				}
			}
		}
		index++
	}
}

// 删除词库内容
func (t *Trie) Remove(keyword string) error {
	if keyword == "" {
		return errors.New("keyword can not be empty")
	}
	runeChars := []rune(keyword)
	t.remove(t.Root, runeChars, 0)

	return nil
}

func (t *Trie) Replace(text string, opts ...ReplaceOption) string {
	defaultSetting := ReplaceSetting{
		ReplaceType: ReplaceType.Holder,
		PlaceHolder: '*',
	}
	for _, opt := range opts {
		opt.Apply(&defaultSetting)
	}

	runes := []rune(text)
	length := len(runes)
	res := make([]rune, 0)
	i := 0
	for i < length {
		node := t.Root // 每次都节点跟开始查询
		if _, ok := node.Children[runes[i]]; !ok {
			// 在词库里没有，正常字符串要加入
			res = append(res, runes[i])
			i++
			continue
		}

		var ok bool

		j := i
	Loop:
		for (!node.End || len(node.Children) > 0) && j < length {
			// 从i开始接着进行node的查询，直到字符串结束或者node结束
			node, ok = node.Children[runes[j]]
			if !ok {
				res = append(res, runes[j])
				j++
				break Loop
			}

			if defaultSetting.ReplaceType == ReplaceType.Holder {
				res = append(res, defaultSetting.PlaceHolder)
			}
			
			// delete则自动不添加

			j++
		}

		if j == i {
			i++
		} else {
			i = j
		}
	}

	return string(res)
}
