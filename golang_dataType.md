1. Go 정의된 데이터 타입과 함수
상수	true, false, iota, nil
타입	int, int8, int16, int32, int64
uint, uint8, uint16, uint 32, uint64, uintptr
float32, float64, complex64, complex128
bool, byte, rune, string, error
함수	make, len, cap, new, append, copy, close, delete
complex, real, imag
panic, recover
 

2. 데이터 타입
int	음수가 있는 int로 32비트 환경에서는 int32 의 범위를 갖고 64비트 환경에서는 int64범위를 갖는다.
int8	-128 ~ 127
int16	-32,768 ~ 32,767
int32	-2,147,483,648 ~ 2,147,483,647
int64	-9,223,372,036,854,775,808 ~ 9,223,372,036,854,775,807
uint	음수가 없는 int로 32비트 환경에서는 uint32 의 범위를 갖고 64비트 환경에서는 uint64범위를 갖는다.
uint8	0 ~ 255
uint16	0 ~ 65,535
uint32	0 ~ 4,294,967,295
uint64	0 ~ 18,446,744,073,709,551,615
float32	IEEE-754 32비트 부동소수점 (https://en.wikipedia.org/wiki/Single-precision_floating-point_format)
float64	IEEE-754 64비트 부동소수점 (https://en.wikipedia.org/wiki/Double-precision_floating-point_format)
complex64	float32 크기의 복소수
complex128	float64 크기의 복소수
rune	유니코드를 표현할 때 사용
 

3. 함수
new	동적할당으로 포인터 반환 (zero값으로 nill 초기화)
make	slice, map, chan을 사용할 때 사용(사용시 초기화 가능)
len	길이 반환(배열, 슬라이드 등)
cap	capacity 반환 (배열, 슬라이스 등)
real	실수부 추출
imag	허수부 추출
complex	실수부와 허수부를 복소수로 생성
panic	현재 함수를 즉시 멈추고 defer함수를 모두 실행 후 즉시 리턴
recover	defer안에 recover를 넣으므로서 panic발생 시 즉시 리턴하지 않고 panic상태를 제거하고 정상적으로 실행하도록 돕는다.
