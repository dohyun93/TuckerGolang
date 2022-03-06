package ch23_errorHandling

import (
	"bufio"
	"fmt"
	"os"
)

func readFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}

	defer file.Close()

	rd := bufio.NewReader(file)
	line, _ := rd.ReadString('\n')
	return line, nil
}

func writeFile(filename, line string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = fmt.Fprintf(file, line)
	return err
}

const fileName = "myFile.txt"

func ErrorHandling() {
	fmt.Println("에러 핸들링은 프로그램이 실행중 예기치 못한 상황으로 종료되지 않고 적절히 상황대처를 하여 계속 동작하도록 하는 방법.")
	// 예시로 파일 내용을 한 줄 읽고, 출력할 때 만약 파일 읽기 실패시 새로 생성한 후 시도하는 예제를 구현해보자......!
	line, err := readFile(fileName)
	if err != nil {
		err = writeFile(fileName, "This is new line of new file.")
		if err != nil {
			fmt.Println("파일 생성에 실패했다.", err)
			return
		}
		line, err = readFile(fileName)
		if err != nil {
			fmt.Println("파일 읽기에 실패했다.", err)
			return
		}
	}
	fmt.Println("파일 내용: ", line)
}
