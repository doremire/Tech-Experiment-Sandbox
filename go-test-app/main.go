package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    // コマンドライン引数を取得
    args := os.Args

    // 引数が指定されていない場合、ヘルプメッセージを表示
    if len(args) != 2 {
        fmt.Println("Usage: file-reader <file-path>")
        return
    }

    // ファイルをオープン
    file, err := os.Open(args[1])
    if err != nil {
        fmt.Println(err)
        return
    }
    defer file.Close()

    // ファイルを読み取り、各行を表示
    scanner := bufio.NewScanner(file)
    lineNum := 1
    for scanner.Scan() {
        line := scanner.Text()
        fmt.Printf("%d │ %s\n", lineNum, line)
        lineNum++
    }

    if err := scanner.Err(); err != nil {
        fmt.Println(err)
    }
}
