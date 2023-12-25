package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"golang.org/x/exp/maps"
)

func main() {
	var nCompany, BobCompany int
	fmt.Scan(&nCompany, &BobCompany)

	ch := make(chan int, nCompany)
	chs := make(chan string, nCompany)

	mBobCompany := CreateMCompany(BobCompany)
	go getData(nCompany, ch, chs)

	_, root := buildTreeNew(nCompany, ch, chs)
	var res int
	CountBurbl(root, mBobCompany, &res)
	fmt.Println(res)
}

type TreeNode struct {
	Parent *TreeNode
	Key    int
	Val    string
	Price  int
	Child  []*TreeNode
}

func CreateMCompany(n int) map[string]int {
	m := make(map[string]int, n)
	for i := 0; i < n; i++ {
		var s string
		fmt.Scan(&s)
		m[s] = 0
	}
	return m
}

func getData(n int, out chan int, outs chan string) {

	var buf []byte
	buf, _ = ioutil.ReadAll(os.Stdin)
	num := int(0)
	var s []byte
	for _, v := range buf {
		if '0' <= v && v <= '9' && true {
			num = 10*num + int(v-'0')
		} else if 'A' <= v && v <= 'Z' {
			s = append(s, v)
		} else if v == 32 {
			out <- num
			num = 0
		} else if v == 10 {
			outs <- string(s)
			s = nil
		}
	}
	close(out)
	close(outs)
}

func buildTreeNew(n int, chi chan int, chs chan string) (map[int]*TreeNode, *TreeNode) {
	Tree := make(map[int]*TreeNode, n)
	for i := 0; i <= n; i++ {
		Tree[i] = &TreeNode{Key: i, Val: "", Price: 0, Child: []*TreeNode{}}
	}
	i := 1
	var root *TreeNode
	for v := range chs {
		i1, i2 := <-chi, <-chi
		fmt.Println(i1)
		s := Tree[i1].Child
		if i1 == 0 {
			root = Tree[i]
		}
		Tree[i].Parent = Tree[i1]
		Tree[i].Val = v
		Tree[i].Price = i2
		s = append(s, Tree[i])
		Tree[i1].Child = s
		i++
	}
	return Tree, root
}

func CountBurbl(v *TreeNode, m map[string]int, res *int) {
	var sum int
	flag := true
	var sn []map[string]int
	if v.Child != nil {
		for _, j := range v.Child {
			n := maps.Clone(m)
			CountBurbl(j, n, res)
			sn = append(sn, n)
		}
		for _, s := range sn {
			for k, v := range s {
				m[k] = v
			}
		}
	}
	m[v.Val] += v.Price
	for _, v := range m {
		if v == 0 {
			flag = false
			break
		}
		sum += v
	}
	if flag && (*res > sum || *res == 0) {
		*res = sum
	}
}
