# GoRoutine

## Thread and Context Swtiching
* Context Switching : 쓰레드 전환에 비용(성능)이 발생하는 데 이를 컨텍스트 스위칭이라 함
* Multi thread vs Multi process
*   * process : Process는 실행 주체 (프로그램을 실행시키는 주체)
*   * multi thread : 하나의 프로세스 안 에서도 실행 흐름이 여러개가 존재(프로그램 안에서)

## Go Routine
* Go에서 만든 OS Thread를 이용하는 경량 Thread
* 모든 Go 프로그램은 한 개의 Go Routine(main go routine)이 있음.
* Go Routine은 아무리 많은 고루틴을 생성해도 OS 쓰레드가 고정되어 있기 때문에 컨텍스트 스위칭 비용이 발생하지 않는 장점이 있음!!
* go [함수 이름]

## sync.WaitGroup
```go
var wg sync.WaitGroup

wg.Add(3)   // 작업 개수 설정
wg.Done()   // 작업이 완료될 때마다 호출
wg.Wait()   // 모든 작업이 완료될 때까지 대기
```

## 동시성 프로그래밍(Concurrent Programming)
* 

