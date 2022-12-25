## 포인터
포인터는 메모리 주소를 값으로 갖는 타입
<br></br>

```go
var a int
var p *int

p = &a      // a의 메모리 주소를 포인터 변수 p에 대입

*p = 20     // p의 메모리 공간에 20을 복사
``` 
<br></br>
* 여러 포인터 변수가 하나의 변수를 가리킬 수 있다.(여러 포인터가 같은 메모리 공간을 가리킬 수 있다.)

```go
var a int
var p1 *int
var p2 *int
var p3 *int

p1 = &a
p2 = p1
p3 = p1     // p 모두 같은 값
```
[영상 체크포인트](https://youtu.be/gvWXcGDG0k4?t=1024)






