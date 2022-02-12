package ch16_package

import (
	"fmt"
)

func PackageExample() {
	fmt.Println()
	// main패키지 : 프로그램의 시작점을 포함한 패키지. -> main()함수가 시작점이며, 이를 포함하는 패키지가 main 패키지이다.
	// 그 외 패키지: 오픈소스나 기타 구현한 패키지들을 말한다.

	// 패키지 임포트 시, 이름이 동일한 패키지는 별칭을 주어 구별할 수 있다.
	// "text/template"
	// "html/template"가 있을 경우,

	// "text/template"
	// htemplate "html/template" 처럼 htemplate라는 별칭을 패키지에 부여하여 서로를 구별하도록 할 수 있다.

	// 패키지를 import하고 사용하지 않으면 오류가 난다!
	// 사용하지 않는다면 aliasing하는 위치에 _ 만 해주면 된다.
	// e.g., _ "html/template"

	// 16. Package 장 정리
	// 1. 대소문자로 패키지 외부로 공개여부를 결정할 수 있다.
	// 2. 패키지는 전역변수, 전역상수, 함수, 구조체, 별칭 등을 공개할 수 있다.
	// 3. 새 패키지를 만들기 전에 필요한 기능을 제공하는 패키지가 있는지 검색해보고, 있다면 활용하자.
	// 4. Go모듈을 통해 자신만의 프로젝트 폴더를 만들 수 있다.
}
