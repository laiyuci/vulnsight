# Vulnsight 使用說明書

**Vulnsight** 是一款快速、基於模板的漏洞掃描工具，支援 HTTP、DNS、TCP、SSL 等多種協定。

---

## 安裝方式

### 方法一：從 GitHub Releases 下載（推薦）

1. 前往 https://github.com/laiyuci/vulnsight/releases
2. 依照你的作業系統下載對應執行檔：
   - Windows：`vulnsight_版本號_windows_amd64.zip`
   - Linux：`vulnsight_版本號_linux_amd64.zip`
   - macOS：`vulnsight_版本號_macOS_amd64.zip`
3. 解壓縮後，將 `vulnsight`（或 `vulnsight.exe`）放到任意路徑

### 方法二：使用 Docker

```bash
# 拉取最新版本
docker pull laiyuci7/vulnsight:latest

# 執行掃描
docker run laiyuci7/vulnsight:latest -target example.com
```

### 方法三：從原始碼編譯

```bash
git clone https://github.com/laiyuci/vulnsight.git
cd vulnsight
go build -o vulnsight cmd/vulnsight/main.go
```

---

## 加入系統路徑（讓任何地方都能使用）

### Windows
將 `vulnsight.exe` 複製到 `C:\Windows\System32\`
或將執行檔所在資料夾加入「系統環境變數 → PATH」

### Linux / macOS
```bash
sudo mv vulnsight /usr/local/bin/
sudo chmod +x /usr/local/bin/vulnsight
```

---

## 基本指令

| 功能 | 指令 |
|------|------|
| 掃描單一目標 | `vulnsight -target example.com` |
| 掃描多個目標（清單檔） | `vulnsight -list targets.txt` |
| 指定模板目錄 | `vulnsight -target example.com -t http/cves/` |
| 指定嚴重性 | `vulnsight -target example.com -s critical,high` |
| 輸出 JSON | `vulnsight -target example.com -json-export output.json` |
| 輸出 Markdown | `vulnsight -target example.com -markdown-export report/` |
| 查看版本 | `vulnsight -version` |
| 更新模板 | `vulnsight -update-templates` |
| 顯示所有說明 | `vulnsight -h` |

---

## Docker 使用方式

```bash
# 掃描單一目標
docker run laiyuci7/vulnsight:latest -target example.com

# 掃描並儲存結果（掛載本機目錄）
docker run -v $(pwd)/output:/output laiyuci7/vulnsight:latest \
  -target example.com \
  -json-export /output/result.json

# 使用本機模板目錄
docker run -v $(pwd)/my-templates:/templates laiyuci7/vulnsight:latest \
  -target example.com \
  -t /templates
```

---

## 常用範例

```bash
# 掃描並顯示 critical/high 漏洞
vulnsight -target example.com -s critical,high

# 同時掃描多個目標
vulnsight -target example.com,test.com -t http/cves/

# 批次掃描清單中的所有目標，輸出結果
vulnsight -list hosts.txt -o scan_result.txt

# 使用 AI 提示產生模板並掃描
vulnsight -target example.com -ai "找出 SQL injection 漏洞"

# 掃描並上傳結果至雲端儀表板
vulnsight -target example.com -dashboard
```

---

## 注意事項

- 只能對**你有授權的目標**進行掃描
- 建議先用 `-validate` 驗證自訂模板
- 第一次執行會自動下載 nuclei-templates 模板庫

---

## 更多資訊

- GitHub：https://github.com/laiyuci/vulnsight
- Docker Hub：https://hub.docker.com/r/laiyuci7/vulnsight
- 官方模板文件：https://docs.nuclei.sh/getting-started/running
