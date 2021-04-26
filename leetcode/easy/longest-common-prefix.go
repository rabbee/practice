package easy

import "strings"

type HashTrieTreeNode struct {
	accessTime int
	v          uint8
	son        map[uint8]*HashTrieTreeNode
}

func NewHashTrieTreeNode(v uint8) *HashTrieTreeNode {
	return &HashTrieTreeNode{1, v, make(map[uint8]*HashTrieTreeNode)}
}

func (node *HashTrieTreeNode) Put(str string) {
	if len(str) == 0 {
		return
	}
	b := str[0]
	if son, existed := node.son[b]; existed {
		son.accessTime++
		son.Put(str[1:])
		//fmt.Println(string(node.v), node.accessTime, string(node.son[b].v), node.son[b].accessTime)
		if son.accessTime > 1 {
			return
		}
		return
	}
	newNode := NewHashTrieTreeNode(b)
	newNode.Put(str[1:])
	node.son[b] = newNode
	//fmt.Println(string(node.v), node.accessTime, string(node.son[b].v), node.son[b].accessTime)
	return
}

func (node *HashTrieTreeNode) MaxPrefix() string {
	var maxV uint8
	maxSonPrefix := ""
	for u, son := range node.son {
		if son.accessTime > 1 && (node.v == 0 || node.accessTime == son.accessTime) {
			sonPrefix := son.MaxPrefix()
			//fmt.Println(string(node.v), node.accessTime, string(son.v), son.accessTime, sonPrefix)
			if len(sonPrefix) < len(maxSonPrefix) {
				continue
			}
			maxSonPrefix = sonPrefix
			maxV = u
		}
	}
	if maxV == 0 {
		return maxSonPrefix
	}
	return string(maxV) + maxSonPrefix
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}
	tree := NewHashTrieTreeNode(0)
	for _, str := range strs {
		tree.Put(str)
		//fmt.Println()
	}
	return tree.MaxPrefix()
}

func longestCommonPrefixV2(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}
	ret := make([]uint8, 0, len(strs[0]))

	if len(strs[0]) == 0 {
		return ""
	}

	target := strs[0]

	idx := 0

	for {
		for _, str := range strs[1:] {

			if len(str) == 0 {
				return ""
			}

			if (idx >= len(target) || idx >= len(str)) || target[idx] != str[idx] {
				if len(ret) < 1 {
					return ""
				}
				return i2string(ret[:idx])

			}
		}
		ret = append(ret, target[idx])
		idx++
	}
}

func i2string(in []uint8) string {
	sb := strings.Builder{}
	for _, u := range in {
		sb.WriteString(string(u))
	}
	return sb.String()
}
