## Testing and Profile

### Testing

테스팅은 go tool과 표준 라이브러리로 제공하고 있습니다. 테스팅은 프로젝트 개발 프로젝트의 시간을 엄청나게 줄여주기 때문에 개발 프로세스에서 중요한 부분입니다. 테스팅 기능과 함께 벤치마킹은 매우 강력한 도구입니다. 작성한 코드의 일부의 성능을 리뷰할 수 있습니다. 프로파일링은 함수 사이에 상호작용을 살펴보고 어떤 함수가 가장 많이 사용되는지 볼 수 있습니다.

 * Go에서 tool로 테스팅과 벤치마킹을 지원합니다.
 * tool이 매우 유연해서 다양한 옵션을 줄 수 있습니다.
 * 개발과 나란히 테스트를 작성합니다.
 * 예제 코드는 테스트와 문서의 역할을 합니다.
 * 개발, QA, 릴리즈 단계에서 벤치마킹을 할 수 있습니다.
 * 성능 문제를 발견하면, 코드를 프로파일링해서 어떤 함수에 집중되는지를 볼 수 있습니다.
 * tool은 서로 영향을 줄 수 있습니다. 정밀한 메모리 프로파일링은 CPU 프로파일, goroutine blocking 프로파일링이 스케쥴러 trace에 영향을 미친다. 각 필요한 프로파일링 모드에 대해서 테스트를 재실행합니다.

 * 테스팅 방식
   * 기본 Unit Test
   * 테이블 Unit Test

### Test Coverage
 * 테스트로 여러분의 코드의 얼마를 cover하는 알아보는 도구
 * 결과를 비쥬얼로 표시 가능

```sh
go test -coverprofile cover.out
go tool cover -html=cover.out
```

### Benchmark
벤치마크는 여러분이 작성한 코드의 성능을 테스트할 수 있는 기능입니다. 
go에서는 built-in 도구로 벤치마크를 수행할 수 있습니다.


 * 기본 벤치마크
```go
package benchmark

import(
    "fmt"
    "testing"
)

func BenchmarkSprint(b *testing.B) {
    var s string

    for i:=0; i<b.N; i++ {
        s = fmt.Sprint("hello")
    }
    gs = s
}

func BenchmarkSprintf(b *testing.B) {
    var s string 

    for i:=0; i<b.N; i++ {
        s = fmt.Sprintf("hello")
    }

    gs = s
}
```

 * 부분 벤치마크
```go
package benchmark

import (
    "fmt"
    "testing"
)

var gs string

func BenchmarkSprint(b *testing.B) {
    b.Run("none", benchSprint)
    b.Run("format", benchSprintf)
}

func benchSprint(b *testing.B) {
    var s string
    for i:=0; i<b.N; i++ {
        s = fmt.Sprint("hello")
    }
    gs = s
}

func benchSprintf(b *testing.B) {
    var s string

    for i:=0; i<b.N; i++ {
        s = fmt.Sprintf("hello")
    }
    gs = s
}

```