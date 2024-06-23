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
  - 可以用在 microservices 微服務
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

print.Scan(&userName) 要加`&`否則 terminal 不會顯示輸入 user input 的要求，
這個 pointer 就是變數指向 memory address

當我們初始化一個變數名稱而且賦值給它，就像是拿一個箱子，在箱子寫上名稱(variable)，然後把物品 (value) 放進去

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
- slice 是 array 的抽象化
- slices is more flexible with variable length or getting an sub-array of its own
- 另外 slice 也是基於 index 具有 slice length/size of elements，但有需要時可以重新調整長度
  ```go
  nameOfSlice = append(nameOfSlice, "the string added to the end of a slice")
  ```
- 宣告時也可以用 `:=` 這語法糖
  ```go
  bookings := []string{}
  ```

## Loops in GO
- 迴圈：重複執行一段程式碼很多次
- GO語言裡面只有`for loop`，沒有 while, do while 或 for-each 迴圈
- 底線佔位符 blank identifier
  - 用來忽略不會用到的變數
  - 如果變數沒使用到，需要 make it explicit

## If else in GO
- conditionals
  - `=` for assigning values; `==` for comparing values
- 如果下單了比庫存還多的數量，使用 if + break(or continue) 判斷
  ```bash
  Enter number of tickets: 
  43
  ...
  Enter number of tickets: 
  10
  Thank you pleee her for booking 10 tickets.
  You will receive a confirmation email at y@o
  77g18446744073709551613 tickets remaianing for Go Conference
  The first names of bookings are: [yiwen pleee]
  ```

## Else if statement
- between if and else block

## Conditionalas in for loop
```go
for true {
    // do something infinitely until break
}

for {
    // do something until it's exit using break word
    // more common
}
```

## User input validation 輸入驗證
- 使用`len()`
  - 內建功能，回傳變數的長度，依型別不同
    - arrays & slices: list的大小（含有多少元素）
    - strings: 字符的數量
- logical operator `&&`
    - 稱為邏輯 AND 運算元
    - 兩邊的條件都必須為真，整個expression才會為真

## Switch statement
- 使單一一個變數能同時跟許多個值做比較/驗證
```go
city := "London"

switch city {
    case "New York":
        // 執行預定 New York 會議的票券
    case "Singapore", "Hong Kong":
        // consolidate several values 
    case "London", "Berlin":
        // execute the same logic
    case "Mexico city":
        // 執行預定 Mexico city 會議的票券
    default:
        fmt.Print("No valid city selected")
}
```

## Encapsulate logics with functions
- 將程式碼封裝進各自的container (即 function)
- 每個function有自己的名稱
- 只有主程式呼叫到 function 名稱，才會被執行
- 可以用來減少重複的程式碼

### Functions with return values
- 可以接收傳入參數，再傳回return value
- input/output的參數/回傳值都要明確定義型別
- Return multiple values 一個函式可以回傳多個值
- 如果有多個回傳值，所有回傳值的型別，要在函式簽名中用小括號括起來

### Make code cleaner with package level variables
- 定義多個 functions 都能共用的變數，這樣就不用傳來傳去
- package level variables:
  - 定義在所有 functions 之外
  - 相同 package 底下的所有 files, 所有 functions 都可以使用
  - 無法使用 `:=` 語法糖宣告

### More use cases of functions
- 將類似的邏輯分在同一組
- 重複使用邏輯，減少重複程式碼


## Packages in Go
- 用 package 組織 GO programs，一個 package 就是 `<filename>.go` 檔案的集合
  ```go
  package main
  
  // some functions moved from main.go into helper.go
  ```
  ```bash
  go run main.go helper.go
  // or below if the files are in the same directory
  go run .
  ```

### Multiple packages in your application
- 以 booking-app 為例的話，可以依目的地/國家區分不同 package
  - singapore package
  - vienna package
  - common package
- 步驟：
  1. 將 helper.go 第一行 package main 改成 `package helper`
  2. 在專案目錄下建立新 folder `helper` 並將 helper.go 移動至此目錄下
  3. 調用到 helper 其中函式的 main.go 需要 import helper package
     - 這個模組中，所有 package 的路徑都定義在 `go.mod` 檔案裡面
       ```GO
       import (
           "fmt"
           "strings"
           "go-for-absolute-beginner/helper"
        )
       ```
     - 被調用的函式要加上 package名稱. 的前綴，例 `helper.validateUserInput()`
  4. 被調用的函式自己也要 explictedly exported by its package
     - 把函式名稱改大駝峰，句首大寫 `func ValidateUserInput() {...}
  5. 另外如果要 export package level variable 也可以使用句首大寫的方式

### Scope rules (變數的作用域)
變數的作用域分成三類
1. Local (function level variable)
   - 在函式 function 裡面宣告
   - 在 block 裡面宣告 (例如 for 迴圈裡面, if-else 某條件下)
   - 只能在該 function 或 block 裡面使用
2. Package
   - 在 function 以外宣告
   - 相同 package 不同檔案都能存取使用
3. Global
   - 在 function 以外宣告
   - 且名稱第一個字母大寫
   - 可以跨不同 packages 使用

## Maps
就目前進度而言，每當一個user完成預訂，只會存姓名，現在要使用 maps 建立多個 key-value 對
```
firstName: "Yoshimite"
lastName: "Yeji"
email: "how@how"
ticket: 3
```
- Maps unique keys to values
- 使用 key 存取 value
- 用內建函式 `make(mapp[Tk]Tv)` 建立一個空白的 map