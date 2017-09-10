package main 

import (
	"fmt"
)

const (
	INF = 1000
)

type Dijkstra struct {
	src, des, n int
	graph [][]int
	distance []int 
	unvisited []bool
}

var graph  = [][]int {
	{0, 2, 3, -1},
	{-1, 0, 4, 2},
	{-1, -1, 0, 3},
	{-1, -1 -1, 0},
}

func (d *Dijkstra) Init(src, des, n int, graph [][]int) {
	d.src, d.des, d.n = src, des, n 
	d.graph = graph
	for i:=0; i<n; i++ {
		d.distance = append(d.distance, INF)
		d.unvisited = append(d.unvisited, true)
	}
}

func (d *Dijkstra) minVertex() int {
	minVer := -1
	min := INF 
	for i:=0; i<4; i++ {
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
		v := d.minVertex()
		d.unvisited[v] = false
	
		//3. update distance
		d.updateDistance(v)
	}
	fmt.Println("src: ", d.src, " des: ", d.des, " distance: ", d.distance[d.des])
}

func main() {
	//1. init
	var d Dijkstra
	d.Init(0, 3, 4, graph)
	d.findShortestPath()

	/*
	fmt.Println("d.distance: ", d.distance)
	fmt.Println("d.unvisited: ", d.unvisited)
	fmt.Println("d.graph: ", d.graph)

	d.distance[0] = 0

	for i:=0; i<4-1; i++ {
		//2. choose vertex
		v := minVertex(d)
		fmt.Println("choose vertex : ", v)
		d.unvisited[v] = false

		//3. update distance
		d.distance = updateDistance(v, d)
	}
	*/
	//fmt.Println(d.distance)
}