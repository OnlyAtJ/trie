# GoTrie
# 声明
本代码仅供用于学习探讨，使用者使用于商业任何损失与本人无关
# 示例
## 创建一课数
```go
tr := NewTrie()
```
## 添加词汇
```go
tr.Add("测试敏感词")
```
## 删除词汇
```go
tr.Remove("测试敏感词")
```
## 替换词汇
```go
tr.Replace("我是测试的语句，用来测试敏感词，我爱毛主席，习大大威武", '*')
```
