## 슬라이스

슬라이스는 모든 원소가 **같은** 가변 길이 시퀀스이다. 슬라이스는 "슬라이스의 내부 배열" 이라고 알려져 있는 배열의 원소들(또는 일부)에 접근할 수 있는 경량 자료구조이다. 슬라이스에는 세 가지 구성 요소인 **포인터**, **길이**, **용량**이 있다.

- 포인터는 슬라이스로 접근할 수 있는 배열의 첫 번째 원소를 가리키고, 반드시 배열의 첫 번째 원소일 필요는 없다.

- 길이는 슬라이스 원소의 개수이다. 길이는 용량을 초과할 수 없다. 내장 함수의 `len`은 이 값을 반환한다.

- 용량은 슬라이스 내부 배열의 시작과 끝 사이에 있는 원소의 개수를 의미한다. 내장함수 `cap`은 이 값들을 반환한다.

---

슬라이스는 배열과 달리 비교할 수 없다. `==`로 두 슬라이스가 동일한 원소를 가지고 있는지 테스트할 수 없다. 표준 라이브러리 중 최적화된 `bytes.Equal` 함수로 두 바이트 슬라이스를 비교할 수 있지만, 다른 타입의 슬라이스는 직접 비교 연산을 수행해야 한다.

---

### append 함수

내장된 `append` 함수는 슬라이스에 항목을 추가한다.

```go
var runes = []rune
for _, r := range "Hello" {
    runes = append(runes, r);
}

fmt.Printf("%q"\n, runes) // "['H', 'e', 'l', 'l', 'o']"
```

append가 어떻게 작동하는지는 `append.go` 파일을 살펴보면 알 수 있다.
내장된 append는 하나 이상의 새 원소를 추가하거나 전체 슬라이스도 추가할 수 있다.

```go
var x []int
x = append(x, 1)
x = append(x, 2, 3)
x = append(x, x...) // 슬라이스 x를 추가함
```

`append.go`도 가변형으로 인자를 받기 위해서는 아래와 같이 수정할 수 있다.

```go
func appendInt(x []int, y ...int) []int {
    var z = []int
    zlen := len(x) + len(y)
    copy(z[len(x):], y)
    return z
}
```

---

### 슬라이스 직접 변경

`rotate`나 `reverse`처럼 슬라이스 원소를 직접 변경하지 않는 함수의 예제를 확인해보자.

```go
func remove1(slice []int, i int) []int {
    copy(slice[i:], slice[i+1:])
    return slice[:len(slice) - 1]
}

func remove2(slice []int, i int) []int {
    slice[i] = slice[len(slice) - 1] // 순서를 유지할 필요 없을 때
    return slice[:len(slice) - 1]
 }

func main () {
    s := []int{5, 6, 7, 8, 9}
    fmt.Println(remove(s, 2)) // {5, 6, 8, 9}
}
```
