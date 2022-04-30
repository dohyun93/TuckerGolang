package project2_wordSearchProgram

import (
	"bufio"
	"fmt"
	"os"            // 프로그램 실행 인자로 주어지는 값들 확인 (os.Args)위함.
	"path/filepath" // 경로명(들)을 출력하기 위함.
	"strings"
)

// 모든 파일들을 읽음. -> 2) append 필요
// 	한 파일 내에서 찾기 -> 1) append 필요

// 한 파일 내에서 찾은 라인 수, 라인정보. -> 한 파일 내에서 append해 나간 다음 FilesInfo.LineInfos에 append.
type aFileLineInfo struct {
	lineNo int
	line   string
}

type findInfo struct {
	fileName  string
	LineInfos []aFileLineInfo
}

func Proj2() {
	if len(os.Args) < 3 {
		fmt.Println("최소 3개 이상의 인자를 전달해야 합니다. (명령, 찾을이름, 디렉토리패턴(들))")
		return
	}

	word := os.Args[1]
	directoryPatterns := os.Args[2:]
	// PrintAllFiles(directoryPatterns, word)

	// 인자로 주어진 각 패턴의 경로에 대해 word 찾기.
	// 1. 파일 하나에서 정보 찾기 - v
	// 2. 파일 리스트에서 위 1번 반복해서 정보 append하기.
	foundInfos := []findInfo{}
	// 각 파일 순회하면서 정보 출력
	for _, dirPattern := range directoryPatterns {
		foundInfos = append(foundInfos, getFoundInfoFromPatterns(word, dirPattern)...)
	}
	for _, foundInfo := range foundInfos {
		fmt.Println(foundInfo.fileName, "에서 찾은 ", word, "정보 출력!")
		fmt.Println("==========================================")
		for _, Line := range foundInfo.LineInfos {
			fmt.Println("\t", Line.lineNo, ": \t", Line.line)
		}
		fmt.Println("==========================================")
		fmt.Println()
	}
}

// dirpattern에 해당하는 파일들의 findInfo 슬라이스를 만들고, 최종적으로 41라인에서 ...을 통해 내부 요소들만 갖다가 append시켜준다.
func getFoundInfoFromPatterns(word, dirPattern string) []findInfo {
	findInfosFromPattern := []findInfo{}
	fileNameLists, err := GetFileList(dirPattern)
	if err != nil {
		fmt.Println("패턴에 대응되는 파일들을 가져오는데 실패했습니다.")
		return findInfosFromPattern
	}
	for _, fileName := range fileNameLists {
		findInfosFromPattern = append(findInfosFromPattern, findWordInFile(word, fileName))
	}
	return findInfosFromPattern
}

func findWordInFile(word, filename string) findInfo {
	FoundInfo := findInfo{filename, []aFileLineInfo{}}
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("파일을 열 수 없습니다.")
		return FoundInfo
	}
	defer file.Close()

	lineNo := 1
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, word) {
			FoundInfo.LineInfos = append(FoundInfo.LineInfos, aFileLineInfo{lineNo, line})
		}
		lineNo += 1
	}
	return FoundInfo
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
