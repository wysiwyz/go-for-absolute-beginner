# Golang

## Introduction
- Use cases
- Comparison
- History
- Environment set up
- CLI application

## Core concepts:
- Data types
  - Strings, Integers, Booleans
  - Arrays and Slices
  - Maps
  - Structs
- Variables and constants
- Formatted output
- User input
- Pointers
- Scope Rules
- Loops
- If-else & Switch
- Functions
- Package: how to organize your code in GO packages
- Goroutines

2007 由 Google 開發，2009 開源，為什麼需要這門新的程式語言？因為：
- 近幾年infrastructure變了很多
  - 擴展性更佳，分布式架構
  - dynamic
  - 產能更好 more capacity
  > 但是既有的程式語言無法充分利用這些infrastructure優勢。 
  > 另外 multithreading 可以同時進行很多操作，例如在 google drive 一邊切換上下頁，一邊上傳檔案，在 youtube 可以一邊按讚影片一邊寫留言一邊看影片。但並行處理多執行緒也面臨到兩邊同時編輯文件導致衝突，或者多個用戶同時購買票卷可能會重複預定同樣一張座位/同樣一間房的情況。
- 多核心處理器變得很常見
- 大型網路計算叢集

支援多核心並行 multi-core concurrency
- Go專門設計用來在多核心上面run，而且它支援並發 (concurrency: 系統同時處理多個任務，但任務並不一定在同一個時間點進行)
- Go的Concurrency不僅便宜而且容易達到

Go語言的主要使用情境
- 注重高效能的應用程式
- 在可擴展且分佈式的系統run的程式

Go的語言特性
- Go 目標在於結合以下兩種語言的特性：
  - Python: 語法簡單可讀性高，動態類型的語言
    > dynamic typed: 指的是變數的資料型別不用在編譯時期就決定好，可以在這變數的生命週期中變更資料型別）
  - C++: 靜態型別語言提供的效率跟安全性
- 屬於 Server-side 後端語言
  - 可以用在 microservices 為服務
  - web 網頁應用程式
  - 資料庫服務
- 許多雲端相關服務都是GO寫的：docker, hashiCorp vault, kubernetes, cockroachDB
- 除此之外GO提供了以下優勢：
  1. 語法簡單：容易上手，好維護，好擴充
  2. build, start up, run 速度快
  3. 需要的資源（CPU, ram）比較少
  4. compiled language, 編譯成 binary machine code
     - 因此速度比 Python 這樣的 interpreted language 還要快
     - 在不同作業系統上的表現一致
  5. Simplicity & Speed 
     - 基於這兩個優勢，使其在自動化應用程式跟指令列介面CLI的 DevOps/SRE 應用場景越來越常見

## Local Setup:
安裝 Go & Editor 
1. Go 安裝（附帶 Go CLI tool）
   - [go install](https://go.dev/doc/install)
2. IDE 安裝 (這裡使用 VSCODE)

## `booking-app`
- 先裝好 `GO` extension
- 初始化模組，會建立一個 `go.mod` 檔案
  ```bash
  go mod init <module path such go-for-absolute-beginner>
  ```
- 所有Go程式碼都必須屬於一個 package，也就是說 go file 的第一句話都是 `package ...`
- declaration 
  ```
  func main() {}
  ```
- import fmt package
- run command `go run main.go`

## Variable and constants
### variable
- 用來存值
- 可以想成是值的容器container
- 提供變數一個名稱，用來在app其它地方做引用
- 使用`,`逗號連接字串與變數時會自動加上空格
  ```go
  	var conferenceName = "Go Conference"
	fmt.Println("Welcome to", conferenceName, "booking application")
    // print Welcome to Go Conference booking application
  ```

### Integer Types
- Whole numbers
| Sr.no. | Types and Description                                                |
|--------|----------------------------------------------------------------------|
| uint8  | unsigned 8-bit integers (0 to 255)                                   |
| uint16 | unsigned 16-bit integers (0 to 65535)                                |
| uint32 | unsigned 32-bit integers (0 to 4294967295)                           |
| uint64 | unsigned 64-bit integers (0 to 18446744073709551615)                 |
| int8   | signed 8-bit integers (-128 to 127)                                  |
| int16  | signed 16-bit integers (-32768 to 32767)                             |
| int32  | signed 32-bit integers (-2147483648 to 2147483647)                   |
| int64  | signed 64-bit integers (-9223372036854775808 to 9223372036854775807) |

## Floating Types
- Numbers that contain a decimal component (real numbers)

## Getting user input
fmt package provides functions for formatted input and output (I/O)
- print messages
- collect user input
- write into a file
```go
fmt.Printf("formated content %T and %v", 1, "String")
fmt.Print("content without line break")
fmt.Println("content with line break at the end")
fmt.Scan()
```

## What is a pointer?

print.Scan(&userName) 要加`&`否則terminal不會顯示輸入 user input 的要求，
這個 pointer 就是變數指向 memory address

當我們初始化一個變數名稱而且賦值給它，就像是拿一個箱子，在箱子寫上名稱(variable)，然後把物品(value)放進去

pointer 是另外一個變數，這個變數會指向另外一個變數的記憶體位置，在GO語言裡又稱之為 special variable

C, C++ 有pointer的設計，但 Java, Javascript 沒有

```go
fmt.Scan(userName)
fmt.Scan("") // 這裡傳的是值 pass by value

fmt.Scan(&userName)
fmt.Scan(0x14000102050) // 這裡傳的是參照 pass by reference
```

## Array 陣列
- 固定 size
  - size: 指的是這個 array 可以放多少個元素
- 只能放相同型別的元素
- 向陣列增加元素，透過陣列的index(索引位置)增加或存取元素
  ```go
  nameOfArray[3] = "the string added to index 3 of the array"
  ```

## Slice 切片
- slice 是 array 的抽象話
- slices 更彈性化：variable length or get an sub-array of its own
- 另外 slice 也是基於 index 具有切片大小，但有需要時可以重新調整長度
  ```go
  nameOfSlice = append(nameOfSlice, "the string added to the end of a slice")
  ```
- 宣告時也可以用 `:=` 這語法糖
  ```go
  bookings := []string{}
  ```
- 