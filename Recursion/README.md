# Recursion
 * 헷갈리지만 이해하면 쉬운 부분
 * 예제들로 익히기
 * 종이로 실제 function call을 흐름 이해하기
## Concept
 * 함수 stack
 * 실제 함수를 호출하는 구조
 * breakdown!
## Method
 * 가장 작은 set부터 시작
 * 새로운 element 추가
 * 원래 set에서 모든 element를 다 사용할때까지 반복
 * Base Case 찾기
    * 공집합인 경우 
* 새로운 subset을 생성하기 위해서 해당 elemnet를 기존 subset에 추가
## Example
 * Factory
    * breakdown
 * Hanoi Tower
    * breakdown
 * Reverse String
    * breakdown
 * Binary Search
    * breakdown
 * Subset count
    * breakdown
        * Base
            * 원래 집합이 공집합인 경우, 공집합의 부분집합은 자기 자신
        * Recursion
            * 새로운 부분집함들을 생성하기 위해서 해당 원소를 기존 부분집합들에 추가
    * 가장 작은 subset 생성
    * 모든 subset에 대해서 새로운 subset을 생성하기 위해서 새로운 element를 추가
    * 원래 set로부터 모든 element를 사용할때까지 반복
 * 미로찾기
    * breakdown
        * Base
            * end cell
        * Recursion
            * 방문한 cell 유지
            * 이웃 cell이 방문한 cell인 경우 방문하지 않기
* 철자 순서 바꾸기
    * breakdown
        * Base
            * 1개 단어로 구성된 단어는 자기 자신
        * Recursion
            * 원래 단어에서 글자를 제거하면서 더 작은 단어의 철자 바꾸기 수행
            * 모든 가능한 위치에 글자를 추가하기
## Excersie
 * 