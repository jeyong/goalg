# 알고리즘 분석

## 알고리즘 분석(analysis of algorithms)이란?
 * Donald Knuth가 만든 개념
 * 알고리즘으로 문제를 해결하는데 필요한 자원을 이론적으로 평가하는 것
 * 합리적으로 효과적인 알고리즘을 찾는 방법

 * 계산 복잡도(computational complexity)를 결정
    * 시간
    * 저장공간
    * 기타 자원
 * 일반
   * 알고리즘에 입력의 길이에 따라 시간 복잡도나 공간 복잡도를 결정
   * 성능은 실제 성능에서 upper bound로 worst case 입력으로 결정


 * 점근(asymptotic) vs. 정확(exact)
   * 점근적 평가 (asymptotic estimation)
     * 동일한 알고리즘에 대해서 구현이 달라지면 효율도 달라지니까
     * 이런 차이를 무시하기 위해서
   * 정확한 측정 (exact measure)

## 요약
 * 컴퓨터의 자원
   * CPU
   * 메모리
 * 분석
   * CPU 사용 - 수행시간
   * 메모리 - 용량
 * 성능 측정
   * 수행시간 
     * 동일 환경 (하드웨어, 운영체제 등) 
     * 동일 언어
     * 언어 idiom
   * 복잡도
     * 실제 수행하지 않고 성능 및 효율성 분석
     * 환경, 언어 등의 영향을 배제

 * 분석
   * time complexity (우선)
   * space complexity

 * 시간 복잡도
   * 관념적 수행 시간 
   * 복잡도 함수 : 입력 n에 따라 복잡도 함수

* 표기법
  * n이 커지면 차수가 가장 큰항에 따라  전체 값이 좌우되므로 가장 큰 차수 항만 표시
  * Big O, Big Ω, Big Θ
  * Big-O
    * O(1), O(n), O(log n), O(n log n), O(n2), O(n3), O(2n), O(n!)
    * O(n) vs. O(log n)
      * search vs. binary search
  * [Big-O Cheat Sheet](http://bigocheatsheet.com/)
  
## Big-O
O(1) : 복잡도는 입력에 따라 변하지 않는 경우
constant time complexity를 가진다!

입력의 갯수가 2배, 3배, 4배가 되든지 걸리는 시간은 동일.

"N" : 입력 갯수
O(N)
입력에 따라서 선형으로 증가하는 경우 O(N)의 복잡도를 갖는다고 말한다.

O(N2) :
알고리즘의 복잡도는 N의 증가에 따라 제곱으로 증가

N이 작은 경우 : 
O(N)이나 O(N2)이나 크게 차이가 나지 않는다. 
최근에 빅데이터와 같이 대용량의 데이터를 처리하는 경우가 많다.
만약 N이 10000000 이상이라면?

O(N2 + 1000) = O(N2)
N이 커지면 상수 1000은 무시할 수 있다.

O(N2+N) =  O(N2)과 동일 
N도 무시할 수 있다.


