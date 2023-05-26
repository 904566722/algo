package graph

import (
	"fmt"
	"log"
)

type vertex struct {
	val int
}

/* 基于邻接表实现的无向图 */
type graphAdjMp struct {
	adjMp map[vertex]map[vertex]struct{}
}

func (g *graphAdjMp) addVertex(v vertex)  {
	if _, ok := g.adjMp[v]; ok {
		// 该顶点已经存在，无需重复加入
		return
	}
	// 新加一个节点，以及初始化一个链表
	g.adjMp[v] = map[vertex]struct{}{}
}

func (g *graphAdjMp) deleteVertex(v vertex)  {
	if _, ok := g.adjMp[v]; !ok {
		// 节点是不存在的
		return
	}
	/* 1.删除顶点 */
	delete(g.adjMp, v)
	/* 2.删除其他顶点中与之存在的关联 */
	for _, rt := range g.adjMp {
		delete(rt, v)
	}
}

func (g *graphAdjMp) addEdge(v1, v2 vertex)  {
	_, ok1 := g.adjMp[v1]
	_, ok2 := g.adjMp[v2]
	if !ok1 || !ok2 || v1 == v2 {
		log.Fatal("add edge error")
	}
	g.adjMp[v1][v2] = struct{}{}
	g.adjMp[v2][v1] = struct{}{}
}

func (g *graphAdjMp) deleteEdge(v1, v2 vertex) {
	_, ok1 := g.adjMp[v1]
	_, ok2 := g.adjMp[v2]
	if !ok1 || !ok2 || v1 == v2 {
		log.Fatal("delete edge error")
	}
	delete(g.adjMp[v1], v2)
	delete(g.adjMp[v2], v1)
}


// newGraphAdjMp 使用边来初始化一个图
//	edges 的形式应该是 [(v1,v2),(v1,v3)...]
func newGraphAdjMp(edges [][]vertex) *graphAdjMp {
	g := &graphAdjMp{
		adjMp: make(map[vertex]map[vertex]struct{}),
	}
	for _, e := range edges {
		g.addVertex(e[0])
		g.addVertex(e[1])
		g.addEdge(e[0], e[1])
	}
	return g
}

func (g *graphAdjMp) print()  {
	if g.adjMp == nil {
		fmt.Println("graph is nil")
	}
	for vt, toMp := range g.adjMp {
		fmt.Print("vertex#", vt, "have edge: ")
		for toVt := range toMp {
			fmt.Print("(", vt.val, toVt.val, ")")
		}
		fmt.Println()
	}
}

func (g *graphAdjMp) bfs(startVt vertex) []vertex {
	if g.adjMp == nil {
		fmt.Println("graph is nil")
	}

	var queue []vertex
	visited := make(map[vertex]struct{})
	var visitRst []vertex

	// visitRst 用来存储访问结果
	visitRst = append(visitRst, startVt)
	// 标记已经访问
	visited[startVt] = struct{}{}
	queue = append(queue, startVt)

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		// 遍历相邻节点，如果还没访问，就访问，并加入队列
		for toVt, _ := range g.adjMp[cur] {
			if _, ok := visited[toVt]; ok {
				// 节点已经访问过，跳过
				continue
			}

			visitRst = append(visitRst, toVt)
			visited[toVt] = struct{}{}
			queue = append(queue, toVt)
		}
	}
	return visitRst
}

func (g *graphAdjMp) dfs(startVt vertex) []vertex {
	visited := make(map[vertex]struct{})
	var visitRst []vertex
	g.dfsHelper(startVt, visited, &visitRst)
	return visitRst
}

func (g *graphAdjMp) dfsHelper(startVt vertex, visited map[vertex]struct{}, visitRst *[]vertex)  {
	*visitRst = append(*visitRst, startVt)
	visited[startVt] = struct{}{}
	for toVt := range g.adjMp[startVt] {
		if _, ok := visited[toVt]; ok {
			continue
		}

		g.dfsHelper(toVt, visited, visitRst)
	}
}

var edges = [][]vertex {
	{vertex{val: 1},vertex{val: 2}},
	{vertex{val: 1},vertex{val: 3}},
	{vertex{val: 2},vertex{val: 1}},
	{vertex{val: 2},vertex{val: 4}},
	{vertex{val: 2},vertex{val: 5}},
	{vertex{val: 3},vertex{val: 1}},
	{vertex{val: 3},vertex{val: 4}},
	{vertex{val: 4},vertex{val: 3}},
	{vertex{val: 4},vertex{val: 2}},
	{vertex{val: 4},vertex{val: 5}},
}


func TestGraphAdjMp() {
	g := newGraphAdjMp(edges)
	g.print()
	g.deleteEdge(vertex{val: 1}, vertex{val: 3})
	fmt.Println("delete edge(1,3)")
	g.print()
	g.addEdge(vertex{val: 3}, vertex{val: 5})
	fmt.Println("delete edge(5,3)")
	g.print()
	g.addVertex(vertex{val: 6})
	fmt.Println("add vertex 6")
	g.print()
	g.deleteVertex(vertex{val: 4})
	fmt.Println("delete vertex 4")
	g.print()
}

func TestGraphBFS()  {
	g := newGraphAdjMp(edges)
	if g.adjMp == nil {
		return
	}
	// 从某个节点开始遍历
	for vt := range g.adjMp {
		// 标记为已访问
		visitRst := g.bfs(vt)
		for _, vt := range visitRst {
			fmt.Println("visit rst:", vt)
		}
		break
	}
}

func TestGraphDFS()  {
	g := newGraphAdjMp(edges)
	if g.adjMp == nil {
		return
	}
	// 从某个节点开始遍历
	for vt := range g.adjMp {
		// 标记为已访问
		visitRst := g.dfs(vt)
		for _, vt := range visitRst {
			fmt.Println("visit rst:", vt)
		}
		break
	}
}