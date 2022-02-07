package ch10_switch

import "fmt"

type ColorType int

const (
	RED ColorType = iota
	BLUE
	GREEN
	YELLOW
)

func getMyAge() (myAge int) {
	myAge = 30
	return
}

func SwitchExample() {
	// switch문에서 default는 만족하는 case가 없을 때 실행된다.
	day := "Sunday"
	switch day {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("It is still weekday ;\n")
	case "Saturday", "Sunday":
		fmt.Println("It is weekend!! :) \n")
	default:
		fmt.Println("Any other date???")
	}

	// 2. switch문을 조건문처럼 활용할 수도 있다.
	myAge := 15
	switch { // !!! 비교값을 적지 않을경우 true가 디폴트이다.!!!
	case myAge < 10, myAge > 30:
		fmt.Println("10살 미만이거나 30살 초과")
	case myAge >= 10, myAge <= 20:
		fmt.Println("10살 이상이거나 20살 이하")
		fallthrough // 이 case문 실행하더라도 switch 문 종료하지 않고 아래로 내려감.
	case myAge >= 15, myAge <= 20: // myAge가 15일때 실행안됨. 그러나 바로 위 fallthrough 때문에 실행됨.
		fmt.Println("15살 이상이거나 20살 이하")
		fallthrough
	case myAge >= 15, myAge <= 20: // myAge가 15일때 실행안됨. 그러나 바로 위 fallthrough 때문에 실행됨.
		fmt.Println("15살 이상이거나 20살 이하2")
	default:
		fmt.Println("20살 초과 30살 이하인 사람임.")
	}

	// 3. switch문의 초기문; 비교값
	// switch문도 if문의 초기문;조건문 처럼 초기문;비교값 을 사용할 수 있다.
	switch myAge2 := getMyAge(); true { // 초기문; 비교값(true) -> 비교값인 true와 case의 값이 일치하는 경우를 찾는게 switch문!!!
	case myAge2 <= 20:
		fmt.Println("아직 미성년자네")
	case myAge2 <= 30:
		fmt.Println("20대가 다 갔네")
	case myAge2 < 40:
		fmt.Println("30대가 다 갔네")
	}

	// 4. const 열거값과 switch..
	// const 열거값 (iota)에 따라 수행되는 로직을 변경할 때 switch문을 주로 사용한다.
	// 색깔을 나타내는 열거값을 문자열로 바꾸는 함수를 switch문을 사용해서 만들어보자.
	myFavoriteColor := RED
	fmt.Println("내 최애 색깔은: ", colorToString(myFavoriteColor), "이다.")

	// 5. break와 fallthrough
	// 타 언어는 break을 안해주면 다음 case로 코드가 이어서 실행되지만, Go는 케이스 실행 후 자동으로 switch문을 빠져나간다.
	// fallthrough는 케이스 실행 후 다음 케이스도 실행한다.
	a := 3
	switch a {
	case 1:
		fmt.Println("1입니다.")
	case 2:
		fmt.Println("2입니다.")
	case 3:
		fmt.Println("3입니다.")
		// 실행 후 fallthrough 없으면 switch문 종료
		// fallthrough // 다음 케이스도 실행. -> 가독성을 저해시키므로 되도록 사용하지 않는다.
	case 4:
		fmt.Println("4입니다.")
	case 5:
		fmt.Println("5입니다.")
	default:
		fmt.Println("Undefined.")
	}
}

func colorToString(color ColorType) string {
	switch color {
	case RED:
		return "Red"
	case GREEN:
		return "Green"
	case BLUE:
		return "Blue"
	case YELLOW:
		return "Yellow"
	default:
		return "Undefined"
	}
}
