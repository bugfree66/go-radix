package goradix

import (
    "fmt"
    "net"
    "testing"
)

func TestIPRadix(t *testing.T) {
    var node *IPRadixNode
    var ipnet *net.IPNet
    var ip net.IP

    fmt.Println("-----test TestIPRadix--------")

    tree := NewIPRadixTree()

    _, ipnet, _ = net.ParseCIDR("1.1.1.0/16")
    tree.Insert(ipnet, "a")
    _, ipnet, _ = net.ParseCIDR("1.1.1.0/24")
    tree.Insert(ipnet, "aa")
    tree.Print()

    ip, _, _ = net.ParseCIDR("1.1.1.2/32")
    node = tree.Find(ip)
    if node == nil || node.Value != "aa" {
        t.Errorf("testing key=%s, val=%v\n fail", ip.String(), node)
    }
    fmt.Println()
}
