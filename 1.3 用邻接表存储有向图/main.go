package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type ArcNode struct { // 边节点
	AdjVex  int      // 邻接点序号
	NextArc *ArcNode // 下一个邻接边
}

type VNode struct { // 顶点
	Data   int
	ArcOut *ArcNode // 出度边
	ArcIn  *ArcNode // 入度边
}

type LGraph struct { // 图的邻接表存储结构
	ArcNum int      // 边数
	VNodes []*VNode // 顶点数
}

func main() {
	file, err := os.Open("./testdata/input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var vNum, aNum int
		fmt.Sscan(scanner.Text(), &vNum, &aNum)
		if vNum == 0 && aNum == 0 {
			return
		}

		g := LGraph{
			ArcNum: aNum,
			VNodes: make([]*VNode, vNum),
		}
		for i := 0; i < aNum && scanner.Scan(); i++ {
			var s, e int
			fmt.Sscan(scanner.Text(), &s, &e)
			s--
			e--

			if g.VNodes[s] == nil {
				g.VNodes[s] = &VNode{Data: s}
			}

			if g.VNodes[e] == nil {
				g.VNodes[e] = &VNode{Data: e}
			}

			sNode, eNode := &ArcNode{AdjVex: s}, &ArcNode{AdjVex: e}
			sNode.NextArc = g.VNodes[s].ArcOut
			g.VNodes[s].ArcOut = sNode

			eNode.NextArc = g.VNodes[e].ArcIn
			g.VNodes[e].ArcIn = eNode
		}

		inCounts, outCounts := make([]int, vNum), make([]int, vNum)
		for i := range g.VNodes {
			vNode := g.VNodes[i]
			p := vNode.ArcIn
			for p != nil {
				inCounts[i]++
				p = p.NextArc
			}
			p = vNode.ArcOut
			for p != nil {
				outCounts[i]++
				p = p.NextArc
			}
		}
		for _, v := range outCounts {
			fmt.Print(v)
		}
		fmt.Println()
		for _, v := range inCounts {
			fmt.Print(v)
		}
		fmt.Println()
	}
}
