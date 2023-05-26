package simple_and_midium

func FindCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	visited := map[int]struct{}{}

	cnt := 0
	for i := 0; i < n; i++ {
		if _, ok := visited[i]; ok {
			continue
		}
		graphDFS(isConnected, i, visited)
		cnt++
	}
	return cnt
}

func graphDFS(graph [][]int, start int, visited map[int]struct{}) {
	visited[start] = struct{}{}

	for adj, iscon := range graph[start] {
		if _, ok :=  visited[adj]; ok || adj == start || iscon == 0 {
			continue
		}
		graphDFS(graph, adj, visited)
	}
}
