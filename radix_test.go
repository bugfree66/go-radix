package goradix

import (
    "testing"
)

func TestRadix(t *testing.T) {
    var node *RadixNode
    var key string

    tree := NewRadixTree(false)
    tree.Insert("abc", "123")
    tree.Insert("acd", "123445")
    tree.Insert("cd", 98)
    tree.Print()

    key = "a"
    node = tree.Find(key)
    if node == nil || node.Value != nil {
        t.Errorf("testing key=%s, node:=%v\n fail", key, node)
    }

    key = "abc"
    node = tree.Find(key)
    if node == nil || node.Value.(string) != "123" {
        t.Errorf("testing key=%s, node:=%v\n fail", key, node)
    }

    key = "ab"
    node = tree.Find(key)
    if node != nil {
        t.Errorf("testing key=%s, node:=%v\n fail", key, node)
    }

    key = "acd"
    node = tree.Find(key)
    if node == nil || node.Value.(string) != "123445" {
        t.Errorf("testing key=%s, node:=%v\n fail", key, node)
    }

    key = "ac"
    node = tree.Find(key)
    if node != nil {
        t.Errorf("testing key=%s, node:=%v\n fail", key, node)
    }

    key = "cd"
    node = tree.Find(key)
    if node == nil || node.Value.(int) != 98 {
        t.Errorf("testing key=%s, node:=%v\n fail", key, node)
    }

    key = "c"
    node = tree.Find(key)
    if node != nil {
        t.Errorf("testing key=%s, node:=%v\n fail", key, node)
    }

    key = "bcd"
    node = tree.Find(key)
    if node != nil {
        t.Errorf("testing key=%s, node:=%v\n fail", key, node)
    }
}

func TestRadixIgnore(t *testing.T) {
    var node *RadixNode
    var key string

    tree := NewRadixTree(true)
    tree.Insert("abc", "123")
    tree.Insert("Acd", "123445")
    tree.Insert("cd", 98)
    tree.Print()

    key = "a"
    node = tree.Find(key)
    if node == nil || node.Value != nil {
        t.Errorf("testing key=%s, node:=%v\n fail", key, node)
    }

    key = "aBc"
    node = tree.Find(key)
    if node == nil || node.Value.(string) != "123" {
        t.Errorf("testing key=%s, node:=%v\n fail", key, node)
    }

    key = "ab"
    node = tree.Find(key)
    if node != nil {
        t.Errorf("testing key=%s, node:=%v\n fail", key, node)
    }

    key = "acd"
    node = tree.Find(key)
    if node == nil || node.Value.(string) != "123445" {
        t.Errorf("testing key=%s, node:=%v\n fail", key, node)
    }

    key = "ac"
    node = tree.Find(key)
    if node != nil {
        t.Errorf("testing key=%s, node:=%v\n fail", key, node)
    }

    key = "cd"
    node = tree.Find(key)
    if node == nil || node.Value.(int) != 98 {
        t.Errorf("testing key=%s, node:=%v\n fail", key, node)
    }

    key = "c"
    node = tree.Find(key)
    if node != nil {
        t.Errorf("testing key=%s, node:=%v\n fail", key, node)
    }

    key = "bcd"
    node = tree.Find(key)
    if node != nil {
        t.Errorf("testing key=%s, node:=%v\n fail", key, node)
    }
}
