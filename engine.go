/*
	author: yuansudong
	date: 2019-08-12
*/

package sensitive

import (
	"regexp"
)

// Engine 敏感词过滤器
type Engine struct {
	_Trie  *_TrieTree
	_Noise *regexp.Regexp
}

// _New 返回一个敏感词过滤器
func _New() *Engine {
	return &Engine{
		_Trie:  _NewTrie(),
		_Noise: regexp.MustCompile(`[\|\s&%$@*]+`),
	}
}

// UpdateNoisePattern 更新去噪模式
func (E *Engine) UpdateNoisePattern(pattern string) {
	E._Noise = regexp.MustCompile(pattern)
}

// LoadFromList 用于
func (E *Engine) LoadFromList(
	sArr []string, //  数组
) *Engine {
	E.AddWord(sArr...)
	return E
}

// AddWord 添加敏感词
func (E *Engine) AddWord(words ...string) {
	E._Trie.Add(words...)
}

// DelWord 删除敏感词
func (E *Engine) DelWord(words ...string) {
	E._Trie.Del(words...)
}

// Filter 过滤敏感词
func (E *Engine) Filter(text string) string {
	return E._Trie.Filter(text)
}

// Replace 和谐敏感词
func (E *Engine) Replace(text string) string {
	return E._Trie.Replace(text, _Ignore)
}

// FindIn 检测敏感词
func (E *Engine) FindIn(text string) (bool, string) {
	text = E.RemoveNoise(text)
	return E._Trie.FindIn(text)
}

// FindAll 找到所有匹配词
func (E *Engine) FindAll(text string) []string {
	return E._Trie.FindAll(text)
}

// Validate 检测字符串是否合法
func (E *Engine) Validate(text string) (bool, string) {
	text = E.RemoveNoise(text)
	return E._Trie.Validate(text)
}

// RemoveNoise 去除空格等噪音
func (E *Engine) RemoveNoise(text string) string {
	return E._Noise.ReplaceAllString(text, "")
}
