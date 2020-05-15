package trie

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTrie_AddAndRemove(t *testing.T) {
	ast := assert.New(t)
	tr := NewTrie()

	err := tr.Add("")
	ast.Error(err, ErrKeywordCanNotBeEmpty.Error(), "空字符串报错")
	err = tr.Add("我要钱多多")
	ast.NoError(err)
	jsonB, err := json.Marshal(tr)
	ast.NoError(err)
	expJson := `{"Root":{"Children":{"25105":{"Children":{"35201":{"Children":{"38065":{"Children":{"22810":{"Children":{"22810":{"Children":{},"End":true}},"End":false}},"End":false}},"End":false}},"End":false}},"End":false}}`
	ast.EqualValues(expJson, string(jsonB))

	err = tr.Remove("")
	ast.Error(err, ErrKeywordCanNotBeEmpty.Error(), "空字符串报错")

	err = tr.Remove("我要钱多多")
	jsonB, err = json.Marshal(tr)
	ast.NoError(err)
	expJson = `{"Root":{"Children":{},"End":false}}`
	ast.EqualValues(expJson, string(jsonB))

	err = tr.Add("我要钱多多")
	ast.NoError(err)
	err = tr.Add("我要")
	ast.NoError(err)
	err = tr.Add("我要钱")
	ast.NoError(err)
	jsonB, err = json.Marshal(tr)
	expJson = `{"Root":{"Children":{"25105":{"Children":{"35201":{"Children":{"38065":{"Children":{"22810":{"Children":{"22810":{"Children":{},"End":true}},"End":false}},"End":true}},"End":true}},"End":false}},"End":false}}`
	ast.EqualValues(expJson, string(jsonB))

	err = tr.Remove("我要钱多多")
	ast.NoError(err)
	jsonB, err = json.Marshal(tr)
	expJson = `{"Root":{"Children":{"25105":{"Children":{"35201":{"Children":{"38065":{"Children":{},"End":true}},"End":true}},"End":false}},"End":false}}`
	ast.EqualValues(expJson, string(jsonB))

	err = tr.Remove("我要")
	ast.NoError(err)
	jsonB, err = json.Marshal(tr)
	expJson = `{"Root":{"Children":{"25105":{"Children":{"35201":{"Children":{"38065":{"Children":{},"End":true}},"End":false}},"End":false}},"End":false}}`
	ast.EqualValues(expJson, string(jsonB))
}

func TestTrie_FindForbiddenString(t *testing.T) {
	ast := assert.New(t)
	tr := NewTrie()
	_ = tr.Add("测试敏感词")
	_ = tr.Add("敏感词")
	_ = tr.Add("测试")
	_ = tr.Add("毛主席")
	_ = tr.Add("习大大")

	text := "我是测试的语句，用来测试敏感词，我爱毛主席，习大大威武"
	act := tr.Replace(text, '*')
	ast.EqualValues("我是**的语句，用来*****，我爱***，***威武", act)

	text = "这段文字不包含"
	act = tr.Replace(text, '*')
	ast.EqualValues("这段文字不包含", act)
}
