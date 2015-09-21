package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "path/filepath"
    "strings"
)

func merge(rootPath string) {
    outFileName := "/dev/out.txt"
    outFile, openErr := os.OpenFile(outFileName, os.O_CREATE|os.O_WRONLY, 0600)
    if openErr != nil {
        fmt.Printf("Can not open file %s\r\n", outFileName)
        return
    }
    bWriter := bufio.NewWriter(outFile)
    filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
        fmt.Println("Processing:", path)
        //这里是文件过滤器，表示我仅仅处理txt文件
        if strings.HasSuffix(path, ".txt") {
            fp, fpOpenErr := os.Open(path)
            if fpOpenErr != nil {
                fmt.Printf("Can not open file %v", fpOpenErr)
                return fpOpenErr
            }
            //这里我们采取按字节读取的方式，因为如果采用按行读取的方式的话，那么如果文件
            //内容的行后面没有换行符，那么这行就会被丢失，所以不能采用按行读取的方式
            bReader := bufio.NewReader(fp)
            for {
                buffer := make([]byte, 1024)
                readCount, readErr := bReader.Read(buffer)
                if readErr == io.EOF {
                    break
                } else {
                    bWriter.Write(buffer[:readCount])
                }
            }
        }
        return err
    })
    bWriter.Flush()
}

func main() {
    var rootPath = "/dev/gocodes/src/mabetle/"
    merge(rootPath)
}


