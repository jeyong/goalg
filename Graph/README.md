## Graph
### 기본
 * 용어 정의
  * vertex
   * data
   * index
   * == 연산
  * edge
   * from
   * to
   * == 연산
  * edgelist
   * vertex, edges
  * graph

### 개념
 * vertex와 node로 구성
 * 방향성을 가지는 경우와 가지지 않는 경우
 * 속성 : cost, capacity, length, 
 * 활용 : 

### 관련 연산
 * adjacent : a, b 2개 vertext 사이에 연결(edge)가 있는지
 * neighbors : a vertex와 연결된(edge)를 가지는 vertices의 list
 * add_vertex / delete_vertex : 특정 vertex 삭제
 * add_edge : a, b 사이에 연결(edge) 생성
 * remove_edge : a, b 사이에 연결 삭제

### 자료 표현 방법
 * 2차 행렬로 graph 표현
 * list로 표현


### Depth first search
 * Graph의 node를 방문하는 방법
 * 현재 node에서 다음 node를 방문할 때 왼쪽 node를 먼저 방문
 * 스택(stack)이나 재귀(recursion)를 사용 
 * <그림>

 * 순서
   * root vertex 부터 시작
   * 인접 vertex && unvisited vertex 조건의 vertex 선택
   * 이전 조건의 vertex를 더이상 찾지 못할떄까지 
   * recursive

 * pseudo - 재귀
```
 * root 시작
 * unvisited && neightbor 조건
 * 선택 vertex를 visited 
 * neighbor 중에 선택(전략)하여 재귀 호출
```

 * pseudo - stack
```
 * root 시작
 * unvisited && neightbor 조건
 * 선택 vertex를 visited
 * neighbor를 stack에 넣기 (push)
 * stack이 empty될때까지 stack에서 꺼내서 2번부터 반복(pop)
```

### Breadth first search
 * 동일 level의 node를 먼저 방문 (인접 node 먼저 방문)
 * ![](../img/graph/BFSanimation.gif)

 * peudo code
```
  *  Init
    * set (visited)
    * queue (root를 queue에 넣기)
  * loop
    * queue에서 꺼내기
    * neighbor를 queue에 넣기 (visited가 아니면)
```

### Spanning tree
 * Graph G의 spanning tree T
 * spanning tree는 Graph G의 모든 vertex를 포함하고 있어야 함
 * Graph G의 경우 모든 vertex는 최소 1개의 연결을 가지고 있어야 함
 * 따라서 Graph G는 최소 1개 이상의 Spanning tree를 가진다.
 * 먄약 Graph G가 tree라면 가질 수 있는 Spanning tree는 오직 1개 뿐.

#### 최소 Spanning tree (MST)
 * 각 edge가 가지는 추가 정보
  * cost, weight, length
 * 시작 vertex에서 각 vertex까지 가는데 걸리는 최소 비용
 * 

### Dijkstra algorithm
![](../img/graph/dijkstraanimation.gif)

 * 2개 node간 최단 경로를 찾는 알고리즘
 * 구현
   * 1단계
     * 현재 node를 첫번째 node로 표시
     * 모든 node를 unvisited set에 넣기
     * 임시 distance로 모든 노드를 초기화(무한, 최대, 특정 값)
   * 2단계
     * 현재 node에 대한 각 unvisited node에 대한 distance 계산
     * 현재 node로부터의 임시 distance = Sum = 현재 node distance + edge weight
     * 해당 node의 기존 distance보다 작은 경우에 update (계산한 Sum과 기존 값과 비교해서 작은 값)
   * 3단계
     * unvisited node set에서 현재 node를 삭제
   * 4단계
     * 타겟 node가 visited로 되면 알고리즘이 끝남
   * 5단계
     * unvisted node를 가장 작은 임시 distance값으로 설정. 
     * 2단계로 가서 다시 시작
----
 * 구현
   * 1단계
     * 모든 node에 distance 값을 할당 (초기 node에 대해서 0, 다른 node는 Infinity)
   * 2단계
     * 초기 node를 cur node로 설정. 다른 node는 unvisited set에 넣기.
   * 3단계
     * cur node 기준 neighbor node에 대한 distance 계산. 새로 계산한 값과 기존 값 중에 작은 값으로 update. (cur node의 distance + edge distance)
   * 4단계
     * cur node 기준 모든 neighbor node를 처리할 때 cur node를 visited로 처리해야하므로 unvisited set에서 삭제. visited node는 다시 체크 안함.
   * 5단계
     * 목적지 node가 visited로 되거나 unvisited set에 있는 node들 중에서 가장 작은 distance가 infinity면 종료. 알고리즘 종료!
   * 6단계
     * 가장 작은 distance를 가지는 unvisited node를 선택해서 cur node로 지정하고 단계 3으로 가서 계속 진행

----
 * 순서
```
   * 1단계 - Init
     * Graph
     * distance (INF로 설정)
     * unvisited (true)
   * 2단계 - 대상 Vertex 찾기
     * 가장 작은 distance && unvisited
   * 3단계 - neighbor distance update하기
     * visited
     * (대상 vertex distance + 대상 vertex -> neighbor vertex 거리) 와 neighbor vertex의 현재 distance 비교하여 작은 것을 선택
```

