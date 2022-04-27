package project2_wordSearchProgram

import (
	"bufio"
	"fmt"
	"os"            // 프로그램 실행 인자로 주어지는 값들 확인 (os.Args)위함.
	"path/filepath" // 경로명(들)을 출력하기 위함.
	"strings"
)

// 찾은 라인의 정보
type LineInfo struct {
	lineNo int
	line   string
}

// 파일 내 라인 정보
type FindInfo struct {
	fileName string
	lines    []LineInfo
}

func Proj2() {
	if len(os.Args) < 3 {
		fmt.Println("최소 3개 이상의 인자를 전달해야 합니다. (명령, 찾을이름, 디렉토리패턴(들))")
		return
	}

	word := os.Args[1]
	directoryPatterns := os.Args[2:]
	// PrintAllFiles(directoryPatterns, word)
	findInfos := []FindInfo{}

	// 인자로 주어진 각 패턴의 경로에 대해 word 찾기.
	for _, pathPattern := range directoryPatterns {
		findInfos = append(findInfos, FindWordInAllFiles(word, pathPattern)...)
	}
	for _, findInfo := range findInfos {
		fmt.Println(findInfo.fileName)
		fmt.Println("==================================")
		for _, lineInfo := range findInfo.lines {
			fmt.Println("\t", lineInfo.lineNo, "\t", lineInfo.line)
		}
		fmt.Println("==================================")
		fmt.Println()
	}
}

func FindWordInAllFiles(word, pathPattern string) []FindInfo {
	findInfos := []FindInfo{}
	fileList, err := GetFileList(pathPattern)
	if err != nil {
		fmt.Println("파일을 찾을 수 없습니다. err: ", err)
		return findInfos
	}

	for _, fileName := range fileList {
		findInfos = append(findInfos, FindWordInFile(word, fileName))

	}
	return findInfos
}

func FindWordInFile(word, fileName string) FindInfo {
	findInfo := FindInfo{fileName, []LineInfo{}}
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("파일을 찾을 수 없습니다. err: ", err)
		return findInfo
	}
	defer file.Close()

	lineNo := 1
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, word) {
			findInfo.lines = append(findInfo.lines, LineInfo{lineNo, line})
		}
		lineNo += 1
	}
	return findInfo
}

// directory 패턴을 받아서 해당하는 파일들의 이름을 반환한다.
func GetFileList(directoryPattern string) ([]string, error) {
	return filepath.Glob(directoryPattern)
}

// word가 들어간 파일들의 이름을 모두 출력한다.
func PrintAllFiles(directoryPatterns []string, word string) {
	for idx, directoryPattern := range directoryPatterns {
		fileList, err := GetFileList(directoryPattern)
		fmt.Println(idx+1, "번째 패턴 ", directoryPattern, "탐색 시작..")
		if err != nil {
			fmt.Println("패턴에 대응되는 파일들을 찾을 수 없습니다.")
		}
		// 찾으려는 파일 이름 리스트 출력
		fmt.Println(word, " 키워드가 포함되는 파일(디렉토리)들 이름 출력.")
		for _, fileName := range fileList {
			fmt.Println(fileName)
		}
		fmt.Println("=================================")
	}
}

// 현재 버전은 word를 갖고 어떤 처리를 안한다.
// ./TuckerGolang word ch* proj* 를 하면
// word 상관없이 ch로 시작하거나 proj로 시작하는 파일(경로)이름을 출력한다.
//
// 이제 파일목록들은 출력하니, 그 파일을 실제로 읽어서 한 줄씩 읽어보자.

func PrintFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("파일을 열 수 없습니다.")
		return
	}
	defer file.Close() // 함수 종료 전 반드시 파일 닫아주기.

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text()) // 한 줄씩 읽기.
	}
}
