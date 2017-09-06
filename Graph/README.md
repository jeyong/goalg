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

### Depth first search
 * Graph의 node를 방문하는 방법
 * 현재 node에서 다음 node를 방문할 때 왼쪽 node를 먼저 방문
 * <그림>

### Breadth first search
 * 동일 level의 node를 먼저 방문 (인접 node 먼저 방문)
 * 

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
