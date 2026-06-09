# Vulnsight 完整使用說明書

> 快速、基於模板的漏洞掃描工具｜支援 HTTP / DNS / TCP / SSL 等多種協定

---

## 目錄

1. [安裝方式](#安裝方式)
2. [快速開始](#快速開始)
3. [指令參數總覽](#指令參數總覽)
4. [目標設定](#目標設定)
5. [模板管理](#模板管理)
6. [篩選條件](#篩選條件)
7. [輸出格式](#輸出格式)
8. [速度與效能調整](#速度與效能調整)
9. [Docker 使用方式](#docker-使用方式)
10. [常見使用情境](#常見使用情境)
11. [注意事項](#注意事項)

---

## 安裝方式

### 方法一：Docker（推薦，不需安裝任何環境）

```bash
docker pull laiyuci7/vulnsight:latest
```

### 方法二：從 GitHub Releases 下載執行檔

前往 https://github.com/laiyuci/vulnsight/releases 下載對應版本：

| 作業系統 | 檔案名稱 |
|----------|----------|
| Windows 64 位元 | `vulnsight_版本_windows_amd64.zip` |
| Linux 64 位元 | `vulnsight_版本_linux_amd64.zip` |
| macOS | `vulnsight_版本_macOS_amd64.zip` |

解壓縮後，將執行檔加入系統 PATH 即可全域使用。

**Windows — 加入 PATH：**
```
將 vulnsight.exe 複製到 C:\Windows\System32\
```

**Linux / macOS — 加入 PATH：**
```bash
sudo mv vulnsight /usr/local/bin/
sudo chmod +x /usr/local/bin/vulnsight
```

### 方法三：從原始碼編譯

```bash
git clone https://github.com/laiyuci/vulnsight.git
cd vulnsight
go build -o vulnsight cmd/vulnsight/main.go
```

---

## 快速開始

```bash
# 第一次使用：更新模板庫
vulnsight -update-templates

# 掃描單一目標（全模板，速度慢）
vulnsight -u https://example.com

# 掃描單一目標（只掃高危，速度快）
vulnsight -u https://example.com -s critical,high

# Docker 方式掃描
docker run --rm laiyuci7/vulnsight:latest -u https://example.com -s critical,high
```

---

## 指令參數總覽

```
vulnsight [選項]

選項說明：
  -u, -target       目標 URL 或 IP（可多個，逗號分隔）
  -l, -list         目標清單檔案（每行一個）
  -t, -templates    指定模板或模板目錄
  -s, -severity     依嚴重性篩選（critical/high/medium/low/info）
  -tags             依標籤篩選模板
  -o, -output       結果輸出到檔案
  -j, -jsonl        輸出 JSONL 格式
  -c, -concurrency  同時執行的模板數（預設 25）
  -rl, -rate-limit  每秒最多請求數（預設 150）
  -timeout          請求逾時秒數（預設 10）
  -retries          失敗重試次數（預設 1）
  -v, -verbose      顯示詳細輸出
  -silent           只顯示發現的漏洞
  -version          顯示版本資訊
  -h                顯示完整說明
```

---

## 目標設定

### 單一目標

```bash
vulnsight -u https://example.com
vulnsight -u example.com          # 自動加 https
vulnsight -u 192.168.1.1          # IP 位址
vulnsight -u 192.168.1.0/24       # CIDR 範圍
```

### 多個目標（逗號分隔）

```bash
vulnsight -u https://example.com,https://test.com,https://demo.com
```

### 從清單檔案批次掃描

```bash
# targets.txt 每行一個目標
vulnsight -l targets.txt
```

targets.txt 範例：
```
https://example.com
https://test.example.com
192.168.1.1
192.168.1.0/24
```

---

## 模板管理

### 更新模板庫

```bash
# 更新到最新版本（建議每次掃描前執行）
vulnsight -update-templates

# 查看目前模板版本
vulnsight -tv
```

### 使用指定模板目錄

```bash
# 只跑 HTTP 類模板
vulnsight -u https://example.com -t http/

# 只跑 CVE 模板
vulnsight -u https://example.com -t http/cves/

# 只跑 SSL 模板
vulnsight -u https://example.com -t ssl/

# 只跑 DNS 模板
vulnsight -u https://example.com -t dns/

# 指定多個目錄
vulnsight -u https://example.com -t http/cves/ -t http/exposures/
```

### 使用指定模板檔案

```bash
vulnsight -u https://example.com -t /path/to/my-template.yaml
```

### 常用模板目錄說明

| 目錄 | 說明 |
|------|------|
| `http/cves/` | 已知 CVE 漏洞檢測 |
| `http/exposures/` | 敏感資訊洩漏 |
| `http/misconfiguration/` | 設定錯誤檢測 |
| `http/takeovers/` | 子網域接管檢測 |
| `http/technologies/` | 技術指紋辨識 |
| `http/fuzzing/` | 模糊測試（DAST） |
| `ssl/` | SSL/TLS 憑證問題 |
| `dns/` | DNS 設定問題 |
| `network/` | 網路服務漏洞 |

---

## 篩選條件

### 依嚴重性篩選（最重要）

```bash
# 只掃 critical
vulnsight -u https://example.com -s critical

# 只掃 critical 和 high（建議日常使用）
vulnsight -u https://example.com -s critical,high

# 掃 critical、high、medium
vulnsight -u https://example.com -s critical,high,medium

# 排除 info（掃除 info 以外的所有等級）
vulnsight -u https://example.com -es info
```

嚴重性等級：
| 等級 | 說明 |
|------|------|
| `critical` | 嚴重漏洞，需立即處理 |
| `high` | 高風險漏洞 |
| `medium` | 中風險漏洞 |
| `low` | 低風險漏洞 |
| `info` | 資訊性偵測，非漏洞 |

### 依標籤篩選

```bash
# 只跑 CVE 相關
vulnsight -u https://example.com -tags cve

# 跑 SQL Injection 相關
vulnsight -u https://example.com -tags sqli

# 跑多個標籤
vulnsight -u https://example.com -tags cve,sqli,xss

# 排除特定標籤
vulnsight -u https://example.com -etags dos,fuzz
```

常用標籤：
| 標籤 | 說明 |
|------|------|
| `cve` | CVE 漏洞 |
| `sqli` | SQL Injection |
| `xss` | Cross-Site Scripting |
| `rce` | 遠端程式碼執行 |
| `lfi` | 本地檔案包含 |
| `ssrf` | 伺服器端請求偽造 |
| `oast` | 外部互動偵測 |
| `exposure` | 敏感資訊洩漏 |
| `misconfig` | 設定錯誤 |
| `dos` | 拒絕服務（危險，慎用） |

---

## 輸出格式

### 輸出到終端（預設）

```bash
vulnsight -u https://example.com -s critical,high
```

### 輸出到純文字檔案

```bash
vulnsight -u https://example.com -o results.txt
```

### 輸出到 JSON 檔案

```bash
vulnsight -u https://example.com -json-export results.json
```

### 輸出到 JSONL 格式（每行一筆，適合程式處理）

```bash
vulnsight -u https://example.com -jsonl-export results.jsonl
```

### 輸出到 Markdown 報告

```bash
vulnsight -u https://example.com -markdown-export ./report/
```

### 輸出到 PDF 報告

```bash
vulnsight -u https://example.com -pdf-export report.pdf
```

### 只顯示發現的漏洞（安靜模式）

```bash
vulnsight -u https://example.com -silent
```

---

## 速度與效能調整

### 加速掃描（重要）

```bash
vulnsight -u https://example.com \
  -c 50 \        # 同時跑 50 個模板（預設 25）
  -rl 300 \      # 每秒最多 300 個請求（預設 150）
  -bs 50 \       # 每個模板同時處理 50 個目標（預設 25）
  -timeout 5     # 逾時縮短為 5 秒（預設 10）
```

### 參數說明

| 參數 | 預設值 | 說明 | 建議值 |
|------|--------|------|--------|
| `-c` | 25 | 同時執行的模板數 | 50～100 |
| `-rl` | 150 | 每秒請求數上限 | 300～500 |
| `-bs` | 25 | 每模板並行目標數 | 50 |
| `-timeout` | 10 | 請求逾時（秒） | 5～10 |
| `-retries` | 1 | 失敗重試次數 | 1～3 |

> ⚠️ 速度調太快可能觸發目標的防火牆或 WAF，請依實際情況調整。

---

## Docker 使用方式

### 基本掃描

```bash
docker run --rm laiyuci7/vulnsight:latest -u https://example.com
```

### 加速掃描（建議）

```bash
docker run --rm laiyuci7/vulnsight:latest \
  -u https://example.com \
  -s critical,high \
  -c 50 -rl 300
```

### 掃描結果儲存到本機

```bash
# Windows（PowerShell）
docker run --rm -v ${PWD}:/output laiyuci7/vulnsight:latest \
  -u https://example.com \
  -s critical,high \
  -json-export /output/results.json

# Linux / macOS
docker run --rm -v $(pwd):/output laiyuci7/vulnsight:latest \
  -u https://example.com \
  -s critical,high \
  -json-export /output/results.json
```

### 使用本機目標清單

```bash
# Windows（PowerShell）
docker run --rm -v ${PWD}:/data laiyuci7/vulnsight:latest \
  -l /data/targets.txt \
  -s critical,high \
  -o /data/results.txt

# Linux / macOS
docker run --rm -v $(pwd):/data laiyuci7/vulnsight:latest \
  -l /data/targets.txt \
  -s critical,high \
  -o /data/results.txt
```

### 使用自訂模板

```bash
docker run --rm \
  -v $(pwd)/my-templates:/templates \
  -v $(pwd)/output:/output \
  laiyuci7/vulnsight:latest \
  -u https://example.com \
  -t /templates \
  -o /output/results.txt
```

---

## 常見使用情境

### 情境 1：快速偵查（5 分鐘內完成）

```bash
vulnsight -u https://example.com \
  -s critical,high \
  -tags cve \
  -c 50 -rl 300 \
  -silent
```

### 情境 2：完整安全評估（存報告）

```bash
vulnsight -u https://example.com \
  -s critical,high,medium \
  -c 50 -rl 300 \
  -json-export full-report.json \
  -markdown-export ./report/
```

### 情境 3：批次掃描多個目標

```bash
# 建立目標清單
echo "https://example.com" >> targets.txt
echo "https://test.com" >> targets.txt
echo "https://demo.com" >> targets.txt

# 執行批次掃描
vulnsight -l targets.txt \
  -s critical,high \
  -c 50 \
  -o batch-results.txt
```

### 情境 4：偵測特定漏洞類型

```bash
# 只偵測 SQL Injection
vulnsight -u https://example.com -tags sqli

# 只偵測 XSS
vulnsight -u https://example.com -tags xss

# 偵測敏感資訊洩漏
vulnsight -u https://example.com -t http/exposures/

# SSL 憑證問題
vulnsight -u https://example.com -t ssl/
```

### 情境 5：Docker + 儲存 JSON 結果

```bash
# Windows PowerShell
docker run --rm -v ${PWD}:/output laiyuci7/vulnsight:latest `
  -u https://example.com `
  -s critical,high `
  -c 50 -rl 300 `
  -json-export /output/results.json `
  -silent
```

### 情境 6：只看有發現漏洞的結果（安靜模式）

```bash
vulnsight -u https://example.com -s critical,high -silent
```

---

## 注意事項

### ⚠️ 法律與道德規範

- **只能對你有授權的目標進行掃描**
- 未經授權的掃描可能違反法律
- 在進行滲透測試前，請確認已取得書面授權

### 📝 第一次使用

```bash
# 先更新模板庫（很重要！）
vulnsight -update-templates
```

### 🐳 Docker 模板快取

每次用 Docker 執行都會重新下載模板（約 300MB），建議掛載 volume 快取：

```bash
# Windows PowerShell
docker run --rm \
  -v vulnsight-templates:/root/nuclei-templates \
  -v ${PWD}:/output \
  laiyuci7/vulnsight:latest \
  -u https://example.com \
  -s critical,high \
  -o /output/results.txt
```

第一次會下載，之後就用快取，速度快很多。

### 🔧 Debug 模式

```bash
# 顯示所有請求和回應（除錯用）
vulnsight -u https://example.com -debug

# 只顯示請求
vulnsight -u https://example.com -debug-req

# 顯示詳細輸出
vulnsight -u https://example.com -v
```

### 📊 顯示掃描統計

```bash
vulnsight -u https://example.com -stats
```

---

## 版本資訊

```bash
# 查看 Vulnsight 版本
vulnsight -version

# 查看模板版本
vulnsight -tv
```

---

## 更多資源

- **GitHub**：https://github.com/laiyuci/vulnsight
- **Docker Hub**：https://hub.docker.com/r/laiyuci7/vulnsight
- **模板語法文件**：https://docs.nuclei.sh/templating-guide/introduction
- **完整參數說明**：`vulnsight -h`
