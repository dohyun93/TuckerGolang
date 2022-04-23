package project2_wordSearchProgram

import (
	"fmt"
	"os"
	"path/filepath"
)

func Proj2() {
	if len(os.Args) < 3 {
		fmt.Println("명령, 단어, 파일이름 세 개의 인수를 입력하세요. e.g., find apple filepath")
		return
	}

	word := os.Args[1]
	filesRegExpressions := os.Args[2:]

	fmt.Println("찾으려는 단어: ", word)
	PrintAllFiles(filesRegExpressions)
}

func GetFileList(path string) ([]string, error) {
	return filepath.Glob(path) // path/filepath 라이브러리로 인자 path에 포함되는 파일이름 목록을 가져올 수 있다.
}

func PrintAllFiles(filesRegExpressions []string) {
	for _, fileReg := range filesRegExpressions {
		fileList, err := GetFileList(fileReg)
		if err != nil {
			fmt.Println("파일을 찾을 수 없다. err: ", err)
			return
		}
		fmt.Println("찾으려는 파일 리스트")
		for idx, name := range fileList {
			fmt.Println(idx+1, ". ", name)
		}
	}
}

//Family@DESKTOP-D4RAK17 MINGW64 /c/Go/src/TuckerGolang (main)
//$ ./TuckerGolang.exe main go*
//찾으려는 단어:  main
//찾으려는 파일 리스트
//1 .  go.mod
// // go env GOOS=linux GOARCH=amd64 go build -v main.go
//
//
//
//
