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
  * Big-O
    * O(1), O(n), O(log n), O(n log n), O(n2), O(n3), O(2n), O(n!)
    * O(n) vs. O(log n)
      * search vs. binary search
  * [Big-O Cheat Sheet](http://bigocheatsheet.com/)
