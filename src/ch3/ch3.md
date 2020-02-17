# 기본 데이터 타입

Go의 타입은 기본타입, 결합 타입, 참조 타입, 인터페이스 타입 네 가지 범주로 분류할 수 있다.

기본 타입은 숫자, 문자열, 불리언이 있고, 결합 타입은 여러 단순한 타입의 값을 결합해 더 복잡한 데이터 타입을 구성한다. 참조 타입은 포인터, 슬라이스, 맵, 함수, 채널 등이 있다.

## 기본 타입

숫자에는 정수, 부동소수점 수, 복소수가 있다.

### 정수

Go는 부호 있는 정수 연산과 부호 없는 정수 연산 모두를 제공한다. 부호화된 정수에는 네 가지 크기가 있고, 각 `int8`, `int16`, `int32`, `int64` 타입으로 표시한다. 부호가 없는 버전으로는 앞에 `u`를 붙인다. (`uint8` ... ) 일반적으로 사용되는 그냥 `int`, `uint` 타입도 있다.

### 부동소수점 수

Go에는 두 가지 크기의 부동소수점 수인 `float32`와 `float64`가 있다. 부동소수점의 한계는 `math` 패키지에서 찾을 수 있다.

```go
math.MaxFloat32 // 가장 큰 float32
math.MaxFloat64 // 가장 큰 float64
```

대부분의 경우 `float64`를 사용해야 한다.

cf) NaN에 대해서
math 패키지에는 수학적으로 숫자가 아님을 정의한 NaN 값이 있다. 이 값은 비교 대상이 아니다. (모든 비교 연산자에서 false를 반환한다.)

```go
var z flaot64
fmt.Println(z, -z, 1/z, -1/z, z/z) // 0, -0, +Inf -Inf NaN

nan := math.NaN()
fmt.Println(nan == nan, nan > nan, nan < nan) // 모두 false
```

따라서 부동소수점 결과를 반환하는 함수에서 실패가 예상된다며 , 별도로 실패를 보고 하는 것이 좋다.

```go
func compute() (value float64, ok bool) {
    ...
    if failed {
        return 0, false
    }

    return result, true
}
```

### 불리언

불리언 타입은 값으로 `true`, `false` 두 가지만 허용한다. 일반적인 언어와는 다르게 숫자 값 0 또는 1로 변환되지 않고 반대 역시 falsy truthy 개념이 없다.

### 문자열

문자열은 불변의 바이트 시퀀스이다. 텍스트 문자열은 유니코드 코드 포인트를 UTF-8로 인코딩한 시퀀스로 해석한다.

```go
s := "hello, world"
fmt.Println(len(s)) // 12
fmt.Println(s[0], s[7]) // 104 119 ("h", "w")
```

인덱스 슬라이싱을 사용해서 문자열을 나눌 수도 있다.

```go
fmt.Println(s[0:5]) // "hello"
```

- 연산자를 사용하면 두 문자열을 연결해 새로운 문자열을 만든다.

문자열 값은 불변이므로, 문자열 값 내의 바이트 시퀀스는 변경할 수 없지만, 문자열 변수에 새 값을 할당할 수는 있다.

```go
s := "left"
t := s
s += ", right"

fmt.Println(s)  // left, right
fmt.Println(t)  // left

s[0] = "L" // error
```

### 문자열과 바이트 슬라이스

문자열 조작에는 네 가지 표준 패키지 `bytes`, `strings`, `strconv`, `unicode`가 특히 중요하다.

- `strings` 패키지에는 문자열 검색, 교체, 비교, 잘라내기, 쪼개기, 합치기 등을 수행하는 다양한 함수가 있다.
- `bytes` 패키지는 []byte 타입의 바이트 슬라이스를 조작하며, `strings`의 일부 속성을 공유하는 비슷한 함수들로 이루어져있다.
- `strconv` 패키지에는 불리언, 정수, 부동소수점 수 같은 값을 문자열 값으로 변환 및 역변환하는 함수와 문자열을 참조하거나 역참조 하는 함수가 있다.
- `unicode` 패키지에는 룬을 분류하기 위한 `IsDigit`, `IsLetter`, `IsUpper`, `IsLower` 등의 함수가 있다.

---

바이트 배열을 포함한 문자열은 한 번 생성되면 변경할 수 없다. 반면 바이트 슬라이스의 원소는 자유롭게 변경할 수 있다.

문자열은 바이트 슬라이스로 변환하고 되돌릴 수 있다.

```go
s := "abc"
b := []byte(s)
s2 := string(B)
```

### 문자열과 숫자 사이 변환

숫자와 문자열 표현 사이의 변환이 필요한 경우가 자주 있는데, 이러한 변환은 `strconv`를 사용하는 경우가 많다. 정수를 아스키로 변환하는 함수인 Itoa 함수가 있다.
변환 하기 위한 다른 방법은 `fmt.Sprintf` 함수를 사용할 수도 있다.

```go
x := 123
y := fmt.Sprintf("%d", x);
fmt.Println(y, strconv.Itoa(x)) // 123 123
```

`FormatInt`와 `FormatUint`는 다른 기수를 가진 숫자를 포맷팅할 때 사용한다.

```go
fmt.Println(strconv.FormatInt(int64(x), 2)) // 1111011
```

정수를 나타내는 문자열을 파싱할 때는 `strconv`의 함수 중 `Atoi`나 `ParseInt`를 사용하고, 부호 없는 정수에는 `ParseUint`를 사용한다.

```go
x, err := strconv.Atoi("123")
y, err := strconv.ParseInt("123", 10, 64) // 10진수, 최대 64비트
```

### 상수

싱수는 컴파일러가 값을 알고 있으며, 컴파일 시 평가되는 표현식이다. 상수를 선언할 때 `const`를 사용해서 선언하지만, 실수로 값이 변경되는 것을 방지한다. 상수에는 타입도 지정할 수 있지만, 명시적인 타입이 없으면, 오른쪽 표현식에서 추정하게 되어 있다.

그룹으로 상수가 선언되면 첫 번째 그릅 외에 오른쪽 표현식을 생략할 수 있고 생략한 경우, 그 전 표현식과 타입이 동일하게 적용된다. 이런 특징은 상수라는 점에서 별로 유용하지는 않다.

```go
const (
    a = 1
    b
    c = 2
    d
)

fmt.Println(a, b, c, d) // "1, 1, 2, 2"
```

### 상수 생성기 iota

`iota`를 사용하면 연관된 값을 하나하나 명시 하지 않고 생성할 수 있다. `const` 값에서 `iota` 값은 0에서 시작하고, 각 항목마다 하나씩 증가한다.

```go
type Weekday int

const (
    Sunday Weekday = iota
    Monday
    Tuesday
    Wednesday
    Thursday
    Friday
    Saturday
)
```

좀 더 복잡한 표현식으로도 사용할 수 있다. 아래는
