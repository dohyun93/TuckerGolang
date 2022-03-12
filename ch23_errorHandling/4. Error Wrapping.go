package ch23_errorHandling

import(
	"fmt"
	"errors" // New()
	"bufio"
	"strings"
	"strconv"
)

// 에러를 감싸서 새로운 에러를 만들어야 할 경우도 있다.
// 예) 파일에서 텍스트를 읽어서 특정 타입의 데이터로 변환하는 경우, 1) 파일 읽기 에러 필요하고 2) 타입 변환 에러가 몇 줄의 몇 칸에서 발생한지
// 알면 더 유용할 것이다. 이럴 때 파일 읽기에서 발생한 에러를 감싸고 그 바깥에 줄과 칸 정보를 넣으면 된다.

func MultipleFromString