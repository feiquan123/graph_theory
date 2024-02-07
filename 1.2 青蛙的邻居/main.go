package main

import (
	"fmt"
	"sort"
)

const (
	MaxLakeCnt = 10
)

var (
	groupCnt = 0
)

type Lake struct {
	Idx    int
	Degree int
}

type Lakes []*Lake

func (a Lakes) Len() int           { return len(a) }
func (a Lakes) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Lakes) Less(i, j int) bool { return a[i].Degree > a[j].Degree }

func main() {
	fmt.Println("请输入组数:")
	fmt.Scanln(&groupCnt)

GROUP:
	for g := 0; g < groupCnt; g++ {
		lakeCnt := 0
		fmt.Printf("[组%d]请输入湖的数量：\n", g)
		fmt.Scanln(&lakeCnt)

		lakes := make(Lakes, lakeCnt)
		arr := [MaxLakeCnt][MaxLakeCnt]int{}
		fmt.Printf("[组%d]请输入邻居数目：\n", g)
		for i := 0; i < int(lakeCnt); i++ {
			lakes[i] = &Lake{
				Idx: i,
			}
			fmt.Scanf("%d", &lakes[i].Degree)
		}

		for i := 0; i < lakeCnt; i++ {
			sort.Sort(lakes[i:])
			lake := lakes[i]

			if lake.Degree > lakeCnt-i-1 {
				fmt.Println("NO")
				continue GROUP
			}

			for offset := 1; offset <= lake.Degree; offset++ {
				nextLake := lakes[i+offset]
				nextLake.Degree--
				if nextLake.Degree < 0 {
					fmt.Println("NO")
					continue GROUP
				}
				arr[lake.Idx][nextLake.Idx] = 1
				arr[nextLake.Idx][lake.Idx] = 1
			}
		}

		fmt.Println("YES")
		for i := 0; i < len(lakes); i++ {
			for j := 0; j < len(lakes); j++ {
				fmt.Print(arr[i][j])
				if j != len(lakes)-1 {
					fmt.Print(" ")
				}
			}
			fmt.Print("\n")
		}
	}
}
