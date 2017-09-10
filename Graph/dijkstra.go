package main 

import (
	"fmt"
)

const (
	INF = 1000
	UNCONN = -1
)

type Dijkstra struct {
	src, des, n int
	graph [][]int
	distance []int 
	unvisited []bool
}

var graph  = [][]int {
	{0, 2, 3, UNCONN},
	{UNCONN, 0, 4, 2},
	{UNCONN, UNCONN, 0, 3},
	{UNCONN, UNCONN, UNCONN, 0},
}

func (d *Dijkstra) Init(src, des, n int, graph [][]int) {
	d.src, d.des, d.n = src, des, n 
	d.graph = graph
	for i:=0; i<n; i++ {
		d.distance = append(d.distance, INF)
		d.unvisited = append(d.unvisited, true)
	}
}

func (d *Dijkstra) pickNextVertex() int {
	minVer := -1
	min := INF 
	for i:=0; i<d.n-1; i++ {
		if d.distance[i] < min && d.unvisited[i] == true {
			min = d.distance[i]
			minVer = i 
		}
	}
	return minVer
}

func (d *Dijkstra) updateDistance(s int){
	for i:=0; i<4; i++ {
		if (d.graph[s][i] > -1) && (d.graph[s][i] + d.distance[s] < d.distance[i]) && (d.unvisited[i] == true) {
			d.distance[i] = d.graph[s][i] + d.distance[s]
		}
	}
}

func (d *Dijkstra) findShortestPath(){
	d.distance[d.src] = 0
	
	for i:=0; i<4-1; i++ {
		//2. choose vertex
		v := d.pickNextVertex()
		d.unvisited[v] = false
	
		//3. update distance
		d.updateDistance(v)
	}
	fmt.Println("src: ", d.src, " des: ", d.des, " distance: ", d.distance[d.des])
}

func main() {
	var d Dijkstra

	//1. init
	d.Init(0, 3, 4, graph)
	d.findShortestPath()

}