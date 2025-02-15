package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// f, err := os.Open("./example.txt")
	// if err != nil {
	// 	//log the error
	// 	panic(err)

	// }
	// fileInfo, err := f.Stat()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("file name:", fileInfo.Name())
	// fmt.Println("file or folder:", fileInfo.IsDir())
	// fmt.Println("file size:", fileInfo.Size())
	// fmt.Println("file permission:", fileInfo.Mode())
	// fmt.Println("file modified at ", fileInfo.ModTime())

	// //read file

	// defer f.Close()

	// buf := make([]byte, fileInfo.Size())
	// d, err := f.Read(buf)
	// if err != nil {
	// 	panic(err)
	// }
	// for i := 0; i < len(buf); i++ {

	// 	println("data", d, buf)
	// }
	//read in one go
	// file, err := os.ReadFile("example.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(file))

	//read folder

	dir, err := os.Open(".")
	if err != nil {
		panic(err)
	}
	defer dir.Close()
	fileInfo, err := dir.ReadDir(-1)
	for _, fi := range fileInfo {
		fmt.Println(fi.Name(), fi.IsDir())
	}
	//create a file
	os.Create("example2.txt")
	f, err := os.Create("example2.txt")
	if err != nil {
		panic(err)

	}
	defer f.Close()

	f.WriteString("hi go ")

	bytes := []byte("hey go")
	f.Write(bytes)

	//read and write ot another file (streaming fashion )

	sourceFile, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}
	defer sourceFile.Close()
	destFile, err := os.Create("example.txt")
	if err != nil {
		panic(err)

	}

	defer destFile.Close()
	reader := bufio.NewReader(sourceFile)
	writer := bufio.NewWriter(destFile)
	for {
		B, ERR := reader.ReadByte()
		if ERR != nil {
			if ERR.Error() != "EOF" {
				panic(ERR)
			}
			break
		}
		err := writer.WriteByte(B)
		if err != nil {
			panic(err)
		}

	}
	writer.Flush()
	fmt.Println("written to new file successfully")

	//DELETE a FILE

	err = os.Remove("example.txt")
	if err != nil {
		panic(err)
	}

}
