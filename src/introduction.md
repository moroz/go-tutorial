# 肉肉大學網頁開發入門課程

## 前言：追求簡單易懂

我開始認真學習編程大概是在2011年的寒假，當時尚未成年的我前往我爸爸家過聖誕節與元旦之間的幾天。
由於我爸爸很久以前就當了網頁工程師，當時學習 Ruby on Rails 框架與 Git 的時候偶爾向爸爸請教。
當時，靜態網頁的概念還沒有非常流行，所以第一個 Rails 專案便是我的部落格。
2016年的冬天我參加了一個時間為兩週的 bootcamp，而這兩週內我進一步了解了更進階的概念，如 OAuth2（使用第三方帳號註冊與登入，像是 Log in with Facebook 的按鈕），並開始學習 JavaScript。
爾後的幾年內，我在日常工作中使用 Ruby on Rails。我很快就發現，雖然我會用 Rails 與 Ruby 開發相當複雜的功能，但其實一直以來，我完全不了解該框架的基本操作方法，甚至不明白應用程式內的資料流程。
Ruby on Rails 早期的賣點就是，寫程式的時候沒必要了解程式的細節，反而多加利用預設值與隱藏的邏輯。

## 課程內容

本課程的目標是希望帶領各位讀者從「略懂電腦」到「深入了解網頁開發」。如果這樣的目標聽起來很難達成，那就是因為它真的很難達成。開發軟體很複雜：首先，需要學會一門程式語言，培養分析任務的能力。學會程式語言之後，便要了解 HTTP 協議、瀏覽器、HTML、資料庫等技術，然後再將那幾種技術結合成一款應用程式。

課程中，我們將利用 Go 程式語言與 PostgreSQL 資料庫管理系統開發一款網頁應用程式。Go 語言的基本語法與編程的基礎不在本課程的範圍內，還沒有掌握的同學建議先完成一個 Go 課程。

## 目標群組

我寫這個課程主要是為了我的學生，而我的學生多數為想要轉職並換遠端工作的台灣上班族。

## 學習用電腦

想要修本課程的各位同學將會需要一臺電腦與網路連接。我們將會在那臺電腦上安裝各式各樣的軟體，所以也要確保硬碟上有充足的空間。或許你家裏放著一臺破舊的筆電，可以專門拿來學習，那就是理想狀況，這樣不論我們做什麼樣的實驗都不用擔心資料損失。

寫此篇的時候我有兩臺電腦，一臺是我私人的 Dell XPS，作業系統為 Linux Mint，另一臺是公司提供給我的14吋蘋果 MacBook Pro。不論使用 Linux 還是 macOS，開發經驗都差不多。至於 MS Windows，我們將在本課程所使用的工具：Go, PostgreSQL 與 Node.js 也都可以在 Windows 10 以上的作業系統上正常使用。然而，由於我已經很多年沒有使用 MS Windows 系統，我在學習的途中無法給讀者太多相關的幫助。如果你用的是一臺 PC，我會建議你考慮在上面安裝一個 Linux 發行版，Debian、Manjaro 與 Linux Mint 都是很好的選擇。

## 安裝所需軟體

為了參與這堂課，請各位同學預先安裝最新的 Go。寫本文的時候最新的版本為 1.21.4。
使用 MS&nbsp;Windows 作業系統的同學可以至 <a href="https://go.dev/dl/" target="_blank" rel="noopener noreferrer">Go 官方網站</a>下載安裝包。
Linux 或 macOS 作業系統的使用者可以用 <a href="https://github.com/jdx/rtx" target="_blank" rel="noopener noreferrer">rtx</a> 或 <a href="https://brew.sh/" target="_blank" rel="noopener noreferrer">Homebrew</a> 安裝。

使用 Windows 作業系統的同學請另外安裝 <a href="https://git-scm.com/download/win" target="_blank" rel="noopener noreferrer">Git 版本控制系統</a>。
如果你還沒有<a href="https://github.com/signup" target="_blank" rel="noopener noreferrer">註冊 GitHub</a>，請先註冊並<a href="https://docs.github.com/en/authentication/connecting-to-github-with-ssh/generating-a-new-ssh-key-and-adding-it-to-the-ssh-agent" target="_blank" rel="noopener noreferrer">設定 SSH 金鑰</a>。

你會需要一個程式碼編輯器。我本身偏好 <a href="https://neovim.io/" target="_blank" rel="noopener noreferrer">Neovim</a>，但對於初學者而言，<a href="https://code.visualstudio.com/" target="_blank" rel="noopener noreferrer">Visual Studio Code</a> （以下簡稱 VS&nbsp;Code）是還不錯的選擇。

如使用 VS&nbsp;Code，請安裝 <a href="https://marketplace.visualstudio.com/items?itemName=golang.Go" target="_blank" rel="noopener noreferrer">Rich Go language support for Visual Studio Code</a> 的擴充包。

最後，各位會需要一個瀏覽器。少為人知，目前市場上只剩下三種瀏覽器引擎：Chromium、Gecko 與 Webkit。
市場佔有率最高的便是大家如此熟悉的 Google Chrome 與它背後的引擎 Chromium。Chromium 的開發經驗最好，當我自己在開發前端應用程式的時候或需要測試比較多 HTTP 請求的時候，我一般都用 Google Chrome 測試。

Webkit 就是 macOS 預設瀏覽器 Safari 背後的引擎。其他作業系統也有基於 Safari 的瀏覽器，Linux 上有一個叫作 GNOME Web 的瀏覽器。由於它背後的引擎與 Safari 一樣，因此用 GNOME Web 就可以測試並解決 Safari 使用者才會遇到的錯誤。

最後就是火狐，Mozilla Firefox。它是最後一個背後沒有大企業的瀏覽器，因此珍惜自由的大家應該多加使用 Firefox 瀏覽器。

