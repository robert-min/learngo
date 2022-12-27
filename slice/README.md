# Slice(슬라이스)
* Go에서 제공하는 동적 배열 타입 => Go에서 제공하는 배열을 가리키는 포인터 타입
* []int => 사이즈를 고정하지 않음
<br></br>

## cf) 정적 vs 동적
* 정적 -> static : compile 단계에서(실행도중에는 절대 바뀌지 않음) ex. [10]int
* 동적 -> dynamic : runtime 단계에서
<br></br>

## Slice 선언
```go
// slice 타입 선언
var var_name []int
var_name := []int{}

// 초기화
var var_name = make([]int, 3) // 3은 요소 개수(Len)
var var_name = make([]int, 3, 5) // Len, Cap

// [1 0 0 0 0 2]
var test1 = []int{1, 5:2}
```
<br></br>

## Slice 요소 추가 - append()
```go
var test1 = []int{1, 2, 3}

test2 := append(test1, 4)   // [1, 2, 3, 4]
```
<br></br>

## Slice 동작 원리
* 슬라이스는 실제 배열을 가르키고 있음
```go
// slice은 내부에 struct로 구성되어 있다.
type SliceHeader struct {
    Data uintptr    // 실제 배열을 가리키는 포인터
    Len int         // 요소 개수
    Cap int         // 실제 배열의 길이(Capacity)
}
```
<br></br>

[영상체크포인트](https://youtu.be/z-_6o7WYkiE?t=1644)



