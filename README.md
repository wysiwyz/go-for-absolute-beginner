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
