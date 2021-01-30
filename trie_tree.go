/*
	author: yuansudong
	date: 2019-08-12
*/
package sensitive

// _TrieTree 短语组成的Trie树.
type _TrieTree struct {
	Root *_Node
}

// _Node Trie树上的一个节点.
type _Node struct {
	isRootNode bool
	isPathEnd  bool
	Character  rune
	Children   map[rune]*_Node
}

// _NewTrie 新建一棵Trie
func _NewTrie() *_TrieTree {
	return &_TrieTree{
		Root: _NewRootNode(0),
	}
}

// Add 添加若干个词
func (tree *_TrieTree) Add(words ...string) {
	for _, word := range words {
		tree.add(word)
	}
}

func (tree *_TrieTree) add(word string) {
	var current = tree.Root
	var runes = []rune(word)
	for position := 0; position < len(runes); position++ {
		r := runes[position]
		if next, ok := current.Children[r]; ok {
			current = next
		} else {
			newNode := _NewNode(r)
			current.Children[r] = newNode
			current = newNode
		}
		if position == len(runes)-1 {
			current.isPathEnd = true
		}
	}
}

// Del 执行批量删除操作
func (tree *_TrieTree) Del(words ...string) {
	for _, word := range words {
		tree.del(word)
	}
}

// del 删除单个
func (tree *_TrieTree) del(word string) {
	var current = tree.Root
	var runes = []rune(word)
	for position := 0; position < len(runes); position++ {
		r := runes[position]
		if next, ok := current.Children[r]; !ok {
			return
		} else {
			current = next
		}
		if position == len(runes)-1 {
			current.SoftDel()
		}
	}
}

// Replace 词语替换
func (tree *_TrieTree) Replace(text string, character rune) string {
	var (
		parent  = tree.Root
		current *_Node
		runes   = []rune(text)
		length  = len(runes)
		left    = 0
		found   bool
	)
	for position := 0; position < len(runes); position++ {
		current, found = parent.Children[runes[position]]
		if !found || (!current.IsPathEnd() && position == length-1) {
			parent = tree.Root
			position = left
			left++
			continue
		}
		if current.IsPathEnd() && left <= position {
			for i := left; i <= position; i++ {
				runes[i] = character
			}
		}
		parent = current
	}
	return string(runes)
}

// Filter 直接过滤掉字符串中的敏感词
func (tree *_TrieTree) Filter(text string) string {
	var (
		parent      = tree.Root
		current     *_Node
		left        = 0
		found       bool
		runes       = []rune(text)
		length      = len(runes)
		resultRunes = make([]rune, 0, length)
	)
	for position := 0; position < length; position++ {
		current, found = parent.Children[runes[position]]

		if !found || (!current.IsPathEnd() && position == length-1) {
			resultRunes = append(resultRunes, runes[left])
			parent = tree.Root
			position = left
			left++
			continue
		}
		if current.IsPathEnd() {
			left = position + 1
			parent = tree.Root
		} else {
			parent = current
		}
	}
	resultRunes = append(resultRunes, runes[left:]...)
	return string(resultRunes)
}

// Validate 验证字符串是否合法，如不合法则返回false和检测到
// 的第一个敏感词
func (tree *_TrieTree) Validate(text string) (bool, string) {
	const (
		Empty = ""
	)
	var (
		parent  = tree.Root
		current *_Node
		runes   = []rune(text)
		length  = len(runes)
		left    = 0
		found   bool
	)
	for position := 0; position < len(runes); position++ {
		current, found = parent.Children[runes[position]]
		if !found || (!current.IsPathEnd() && position == length-1) {
			parent = tree.Root
			position = left
			left++
			continue
		}
		if current.IsPathEnd() && left <= position {
			return false, string(runes[left : position+1])
		}
		parent = current
	}
	return true, Empty
}

// FindIn 判断text中是否含有词库中的词
func (tree *_TrieTree) FindIn(text string) (bool, string) {
	validated, first := tree.Validate(text)
	return !validated, first
}

// FindAll 找有所有包含在词库中的词
func (tree *_TrieTree) FindAll(text string) []string {
	var matches []string
	var (
		parent  = tree.Root
		current *_Node
		runes   = []rune(text)
		length  = len(runes)
		left    = 0
		found   bool
	)
	for position := 0; position < length; position++ {
		current, found = parent.Children[runes[position]]
		if !found {
			parent = tree.Root
			position = left
			left++
			continue
		}
		if current.IsPathEnd() && left <= position {
			matches = append(matches, string(runes[left:position+1]))
		}
		if position == length-1 {
			parent = tree.Root
			position = left
			left++
			continue
		}
		parent = current
	}
	var i = 0
	if count := len(matches); count > 0 {
		set := make(map[string]struct{})
		for i < count {
			_, ok := set[matches[i]]
			if !ok {
				set[matches[i]] = struct{}{}
				i++
				continue
			}
			count--
			copy(matches[i:], matches[i+1:])
		}
		return matches[:count]
	}
	return nil
}

// NewNode 新建子节点
func _NewNode(character rune) *_Node {
	return &_Node{
		Character: character,
		Children:  make(map[rune]*_Node, 0),
	}
}

// NewRootNode 新建根节点
func _NewRootNode(character rune) *_Node {
	return &_Node{
		isRootNode: true,
		Character:  character,
		Children:   make(map[rune]*_Node, 0),
	}
}

// IsLeafNode 判断是否叶子节点
func (node *_Node) IsLeafNode() bool {
	return len(node.Children) == 0
}

// IsRootNode 判断是否为根节点
func (node *_Node) IsRootNode() bool {
	return node.isRootNode
}

// IsPathEnd 判断是否为某个路径的结束
func (node *_Node) IsPathEnd() bool {
	return node.isPathEnd
}

// SoftDel 置软删除状态
func (node *_Node) SoftDel() {
	node.isPathEnd = false
}
