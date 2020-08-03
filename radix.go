package github.com/kitrap/go-radix

import (
    "fmt"
    "unicode"
    "strings"
)

type CallBack func(data interface{})

type RadixNode struct {
    Prefix string
    SubNodes map[byte]*RadixNode
    Value interface{}
}

type RadixTree struct {
    Root *RadixNode

    insensitive bool
}

func longestPrefix(s1, s2 string, insensitive bool) int {
    n := len(s1)
    if n > len(s2) {
        n = len(s2)
    }
    i := 0
    for ; i < n; i++ {
        if insensitive {
            if unicode.ToLower(rune(s1[i])) != unicode.ToLower(rune(s2[i])) {
                break
            }
        } else {
            if s1[i] != s2[i] {
                break
            }
        }
    }
    return i
}

func NewRadixNode(prefix string) *RadixNode {
    node := new(RadixNode)
    node.Prefix = prefix
    node.SubNodes = make(map[byte]*RadixNode)
    return node
}

func (node *RadixNode) DebugString(prefix string) string {
    return strings.Join([]string{ prefix+node.Prefix, fmt.Sprintf("%v", node.Value)}, "=")
}

func NewRadixTree(insensitive bool) *RadixTree {
    tree := new(RadixTree)
    tree.insensitive = insensitive
    return tree
}

func (tree *RadixTree) Insert(key string, val interface{}) (*RadixNode, bool) {
    var parent *RadixNode = nil
    var ok bool

    insert := false

    if tree.Root == nil {
        tree.Root = NewRadixNode(key)
        tree.Root.Value = val
        return tree.Root, true
    }

    node := tree.Root
    for {
        if node == nil {
            break
        }
        n := longestPrefix(key, node.Prefix, tree.insensitive)
        if n == len(key) {
            break
        } else {
            if n == len(node.Prefix) {
                parent = node
                if tree.insensitive {
                    node, ok = parent.SubNodes[byte(unicode.ToLower(rune(key[0])))]
                } else {
                    node, ok = parent.SubNodes[key[0]]
                }
                if !ok {
                    newNode := NewRadixNode(key)
                    if tree.insensitive {
                        parent.SubNodes[byte(unicode.ToLower(rune(key[0])))] = newNode
                    } else {
                        parent.SubNodes[key[0]] = newNode
                    }
                    node = newNode
                    insert = true
                }
            } else {
                if parent == nil {
                    if n == 0 {
                        parent = NewRadixNode("")
                    } else {
                        parent = NewRadixNode(key[0:n])
                    }
                    tree.Root = parent
                }
                newNode := NewRadixNode(key)
                if tree.insensitive {
                    parent.SubNodes[byte(unicode.ToLower(rune(node.Prefix[n])))] = node
                    parent.SubNodes[byte(unicode.ToLower(rune(key[n])))] = newNode
                } else {
                    parent.SubNodes[node.Prefix[n]] = node
                    parent.SubNodes[key[n]] = newNode
                }
                node = newNode
                insert = true
            }
        }
        if insert {
            node.Value = val
            break
        }
    }

    return node, insert
}

func (tree *RadixTree) Search(key string) (*RadixNode, bool) {
    var match *RadixNode = nil
    var ok bool
    exact := false

    node := tree.Root
    n := 0
    for {
        if node == nil {
            break
        }
        n = longestPrefix(key, node.Prefix, tree.insensitive)
        if n != len(node.Prefix) {
            break
        }
        match = node
        if n == len(key) {
            exact = true
            node = nil
        } else {
            if tree.insensitive {
                node, ok = node.SubNodes[byte(unicode.ToLower(rune(key[n])))]
            } else {
                node, ok = node.SubNodes[key[n]]
            }
            if !ok {
                break
            }
        }
    }

    return match, exact
}

func (tree *RadixTree) Find(key string) *RadixNode {
    node, exact := tree.Search(key)
    if exact {
        return node
    }
    return nil
}

func (tree *RadixTree) TraverseNode(root *RadixNode, cb CallBack) {
    if root == nil {
        return
    }
    if cb != nil {
        cb(root)
    }
    for _, v := range root.SubNodes {
        tree.TraverseNode(v, cb)
    }
}

func (tree *RadixTree) Remove(key string) {
}

func printNode(node interface{}) {
    if n, ok := node.(*RadixNode); ok {
        fmt.Printf("%s=%v\n", n.Prefix, n.Value)
    }
}

func (tree *RadixTree) Print() {
    tree.TraverseNode(tree.Root, printNode)
}
