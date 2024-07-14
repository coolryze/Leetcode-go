package reconstructitinerary

import (
	"container/heap"
)

type PriorityQueue []string

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i] < pq[j] }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(string))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}

func findItinerary(tickets [][]string) []string {
	graph := make(map[string]*PriorityQueue)
	for _, ticket := range tickets {
		src, dest := ticket[0], ticket[1]
		if graph[src] == nil {
			graph[src] = &PriorityQueue{}
		}
		heap.Push(graph[src], dest)
	}

	var result []string
	var dfs func(string)
	dfs = func(node string) {
		for graph[node] != nil && len(*graph[node]) > 0 {
			next := heap.Pop(graph[node]).(string)
			dfs(next)
		}
		result = append([]string{node}, result...)
	}

	dfs("JFK")
	return result
}
