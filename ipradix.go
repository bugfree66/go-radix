package goradix

import (
    "fmt"
    "net"
)

type IPRadixCallBack func(node *IPRadixNode)

type IPRadixNode struct {
    ipnet net.IPNet
    subnodes [2]*IPRadixNode
    Value interface{}
}

func NewIPRadixNode() *IPRadixNode {
    node := new(IPRadixNode)
    return node
}

type IPRadixTree struct {
    root *IPRadixNode
}

func NewIPRadixTree() *IPRadixTree {
    tree := new(IPRadixTree)
    return tree
}

func getBit(data []byte, n int) int {
    if len(data) * 8 < n {
        return -1;
    }
    b := data[n/8]
    if b & 1 << (n % 8) == 1 {
        return 1
    }
    return 0
}

func (tree *IPRadixTree) longestPrefix(ipnet1, ipnet2 *net.IPNet) int {
    i := 0
    n, _ := ipnet1.Mask.Size()
    if n1, _ := ipnet2.Mask.Size(); n1 < n {
        n = n1
    }
    for ; i < n; i++ {
        if getBit(ipnet1.IP, i) != getBit(ipnet2.IP, i) {
            break
        }
    }
    return i
}

func (tree *IPRadixTree) Insert(ipnet *net.IPNet, val interface{}) (*IPRadixNode, bool) {
    var node, parent *IPRadixNode
    var bit int

    insert := false

    if ipnet == nil {
        return nil, false
    }
    if tree.root == nil {
        node = NewIPRadixNode()
        node.ipnet = *ipnet
        node.Value = val
        tree.root = node
        return node, true
    }
    node = tree.root
    for {
        if node == nil {
            break
        }
        n := tree.longestPrefix(&node.ipnet, ipnet)
        maskSize1, _ := node.ipnet.Mask.Size()
        maskSize2, _ := ipnet.Mask.Size()
        if n == maskSize1 && n == maskSize2 {
            break
        } else {
            if n == maskSize1 {
                parent = node
                bit = getBit(ipnet.IP, n)
                if bit == -1 {
                    return nil, false
                }
                if node.subnodes[bit] == nil {
                    newNode := NewIPRadixNode()
                    newNode.ipnet = *ipnet
                    node.subnodes[bit] = newNode
                    node = newNode
                    insert = true
                } else {
                    node = node.subnodes[bit]
                }
            } else {
                if parent == nil {
                    parent = NewIPRadixNode()
                    if n == 0 {
                        _, tmp, _ := net.ParseCIDR("0.0.0.0/0")
                        parent.ipnet = *tmp
                    } else {
                        _, tmp, _ := net.ParseCIDR(fmt.Sprintf("%s/%d", ipnet.IP.String(), n))
                        parent.ipnet = *tmp
                    }
                    tree.root = parent
                }
                bit = getBit(node.ipnet.IP, n)
                if bit == -1 {
                    return nil, false
                }
                parent.subnodes[bit] = node
                if n < maskSize2 {
                    newNode := NewIPRadixNode()
                    newNode.ipnet = *ipnet
                    bit = getBit(ipnet.IP, n)
                    if bit == -1 {
                        return nil, false
                    }
                    parent.subnodes[bit] = newNode
                    node = newNode
                } else {
                    node = parent
                }
                insert = true
            }
            if insert {
                node.Value = val
                break
            }
        }
    }

    return node, insert
}

func (tree *IPRadixTree) Search(ip net.IP) *IPRadixNode {
    var match *IPRadixNode

    _, ipnet, _ := net.ParseCIDR(fmt.Sprintf("%s/%d", ip.String(), 32))

    node := tree.root
    for {
        if node == nil {
            break
        }
        n := tree.longestPrefix(ipnet, &node.ipnet)
        maskSize1,_ := node.ipnet.Mask.Size()
        if n != maskSize1 {
            break
        }
        match = node
        maskSize2,_ := ipnet.Mask.Size()
        if n == maskSize2 {
            break
        } else {
            bit := getBit(ipnet.IP, n)
            if bit == -1 {
                return nil
            }
            node = node.subnodes[bit]
        }
    }
    return match
}

func (tree *IPRadixTree) Find(ip net.IP) *IPRadixNode {
    return tree.Search(ip)
}

func (tree *IPRadixTree) TraverseNode(root *IPRadixNode, cb IPRadixCallBack) {
    if root == nil {
        return
    }
    if cb != nil {
        cb(root)
    }
    for _, v := range root.subnodes {
        tree.TraverseNode(v, cb)
    }
}

func printIPRadixNode(node *IPRadixNode) {
    fmt.Printf("%s=%v\n", node.ipnet.String(), node.Value)
}

func (tree *IPRadixTree) Print() {
    tree.TraverseNode(tree.root, printIPRadixNode)
}
