# Vulnsight 使用說明書

> 快速、基於模板的漏洞掃描工具｜支援 HTTP / DNS / TCP / SSL 等多種協定

---

## 目錄

1. [安裝](#安裝)
2. [快速開始](#快速開始)
3. [目標設定](#目標設定)
4. [模板選擇](#模板選擇)
5. [篩選條件](#篩選條件)
6. [輸出格式](#輸出格式)
7. [速度調整](#速度調整)
8. [完整指令範例](#完整指令範例)
9. [注意事項](#注意事項)

---

## 安裝

### 前置需求

安裝 Docker Desktop：https://www.docker.com/products/docker-desktop/

### 下載 Vulnsight

```bash
docker pull laiyuci7/vulnsight:latest
```

---

## 快速開始

```bash
# 掃描單一目標（只掃高危，約 5～10 分鐘）
docker run --rm laiyuci7/vulnsight:latest -u https://example.com -s critical,high

# 掃描並將結果存到本機（Windows PowerShell）
docker run --rm -v ${PWD}:/output laiyuci7/vulnsight:latest `
  -u https://example.com `
  -s critical,high `
  -o /output/results.txt

# 掃描並將結果存到本機（Linux / macOS）
docker run --rm -v $(pwd):/output laiyuci7/vulnsight:latest \
  -u https://example.com \
  -s critical,high \
  -o /output/results.txt
```

---

## 目標設定

### 單一目標

```bash
docker run --rm laiyuci7/vulnsight:latest -u https://example.com
docker run --rm laiyuci7/vulnsight:latest -u 192.168.1.1
docker run --rm laiyuci7/vulnsight:latest -u 192.168.1.0/24
```

### 多個目標（逗號分隔）

```bash
docker run --rm laiyuci7/vulnsight:latest \
  -u https://example.com,https://test.com,https://demo.com \
  -s critical,high
```

### 從清單檔案批次掃描

先建立 `targets.txt`，每行一個目標：
```
https://example.com
https://test.example.com
192.168.1.1
```

執行批次掃描：
```bash
# Windows PowerShell
docker run --rm -v ${PWD}:/data laiyuci7/vulnsight:latest `
  -l /data/targets.txt `
  -s critical,high `
  -o /data/results.txt

# Linux / macOS
docker run --rm -v $(pwd):/data laiyuci7/vulnsight:latest \
  -l /data/targets.txt \
  -s critical,high \
  -o /data/results.txt
```

---

## 模板選擇

Vulnsight 共有約 10,000 個偵測模板，**建議指定範圍**以加快速度。

### 依目錄指定

```bash
# 只偵測已知 CVE 漏洞（最常用）
docker run --rm laiyuci7/vulnsight:latest -u https://example.com -t http/cves/

# 偵測敏感資訊洩漏
docker run --rm laiyuci7/vulnsight:latest -u https://example.com -t http/exposures/

# 偵測設定錯誤
docker run --rm laiyuci7/vulnsight:latest -u https://example.com -t http/misconfiguration/

# 偵測 SSL 憑證問題
docker run --rm laiyuci7/vulnsight:latest -u https://example.com -t ssl/

# 偵測 DNS 問題
docker run --rm laiyuci7/vulnsight:latest -u https://example.com -t dns/

# 同時指定多個目錄
docker run --rm laiyuci7/vulnsight:latest -u https://example.com \
  -t http/cves/ -t http/exposures/
```

### 常用模板目錄

| 目錄 | 內容 | 速度 |
|------|------|------|
| `http/cves/` | 已知 CVE 漏洞 | 中 |
| `http/exposures/` | 敏感資訊洩漏 | 快 |
| `http/misconfiguration/` | 設定錯誤 | 快 |
| `http/technologies/` | 技術指紋辨識 | 快 |
| `http/takeovers/` | 子網域接管 | 快 |
| `ssl/` | SSL/TLS 問題 | 快 |
| `dns/` | DNS 問題 | 快 |
| `network/` | 網路服務漏洞 | 慢 |

---

## 篩選條件

### 依嚴重性篩選（建議必用）

```bash
# 只掃最嚴重的（最快）
docker run --rm laiyuci7/vulnsight:latest -u https://example.com -s critical

# 掃高危以上（建議日常使用）
docker run --rm laiyuci7/vulnsight:latest -u https://example.com -s critical,high

# 掃中危以上
docker run --rm laiyuci7/vulnsight:latest -u https://example.com -s critical,high,medium
```

| 等級 | 說明 |
|------|------|
| `critical` | 嚴重漏洞，需立即處理 |
| `high` | 高風險漏洞 |
| `medium` | 中風險漏洞 |
| `low` | 低風險漏洞 |
| `info` | 資訊偵測，非漏洞（數量最多，速度最慢）|

### 依標籤篩選

```bash
# 只偵測 SQL Injection
docker run --rm laiyuci7/vulnsight:latest -u https://example.com -tags sqli

# 只偵測 XSS
docker run --rm laiyuci7/vulnsight:latest -u https://example.com -tags xss

# 只偵測 RCE（遠端程式碼執行）
docker run --rm laiyuci7/vulnsight:latest -u https://example.com -tags rce

# 組合多個標籤
docker run --rm laiyuci7/vulnsight:latest -u https://example.com -tags sqli,xss,rce
```

| 標籤 | 說明 |
|------|------|
| `cve` | CVE 漏洞 |
| `sqli` | SQL Injection |
| `xss` | 跨站腳本攻擊 |
| `rce` | 遠端程式碼執行 |
| `lfi` | 本地檔案包含 |
| `ssrf` | 伺服器端請求偽造 |
| `exposure` | 敏感資訊洩漏 |
| `misconfig` | 設定錯誤 |

---

## 輸出格式

### 純文字（最簡單）

```bash
# Windows PowerShell
docker run --rm -v ${PWD}:/output laiyuci7/vulnsight:latest `
  -u https://example.com -s critical,high `
  -o /output/results.txt
```

### JSON（可用程式讀取）

```bash
# Windows PowerShell
docker run --rm -v ${PWD}:/output laiyuci7/vulnsight:latest `
  -u https://example.com -s critical,high `
  -json-export /output/results.json
```

### Markdown 報告

```bash
# Windows PowerShell
docker run --rm -v ${PWD}:/output laiyuci7/vulnsight:latest `
  -u https://example.com -s critical,high `
  -markdown-export /output/report/
```

### PDF 報告

```bash
# Windows PowerShell
docker run --rm -v ${PWD}:/output laiyuci7/vulnsight:latest `
  -u https://example.com -s critical,high `
  -pdf-export /output/report.pdf
```

---

## 速度調整

預設掃描很慢是因為模板數量龐大。以下參數可大幅加速：

| 參數 | 說明 | 預設值 | 建議值 |
|------|------|--------|--------|
| `-c` | 同時執行的模板數 | 25 | 50 |
| `-rl` | 每秒最多請求數 | 150 | 300 |
| `-timeout` | 請求逾時秒數 | 10 | 5 |

```bash
# 加速版（速度快 3～5 倍）
docker run --rm laiyuci7/vulnsight:latest \
  -u https://example.com \
  -s critical,high \
  -c 50 -rl 300 -timeout 5
```

> ⚠️ 速度設太高可能觸發目標的防火牆，請依情況調整。

### Docker 模板快取（避免每次重新下載）

每次啟動 Docker 都會重新下載約 300MB 的模板，加上快取可大幅加速後續掃描：

```bash
# 第一次：建立快取 volume
docker run --rm \
  -v vulnsight-templates:/root/nuclei-templates \
  laiyuci7/vulnsight:latest -update-templates

# 之後每次使用（速度快很多）
docker run --rm \
  -v vulnsight-templates:/root/nuclei-templates \
  -v ${PWD}:/output \
  laiyuci7/vulnsight:latest \
  -u https://example.com \
  -s critical,high \
  -o /output/results.txt
```

---

## 完整指令範例

### 情境 1：快速偵查（5 分鐘內完成）

```bash
docker run --rm laiyuci7/vulnsight:latest \
  -u https://example.com \
  -s critical,high \
  -tags cve \
  -c 50 -rl 300 \
  -silent
```

### 情境 2：完整安全評估（存 JSON 報告）

```bash
# Windows PowerShell
docker run --rm -v ${PWD}:/output laiyuci7/vulnsight:latest `
  -u https://example.com `
  -s critical,high,medium `
  -c 50 -rl 300 `
  -json-export /output/report.json `
  -markdown-export /output/report/
```

### 情境 3：批次掃描多目標（存結果）

```bash
# Windows PowerShell
docker run --rm -v ${PWD}:/data laiyuci7/vulnsight:latest `
  -l /data/targets.txt `
  -s critical,high `
  -c 50 -rl 300 `
  -o /data/results.txt
```

### 情境 4：只偵測特定漏洞類型

```bash
# 偵測 SQL Injection + XSS（存結果）
docker run --rm -v ${PWD}:/output laiyuci7/vulnsight:latest `
  -u https://example.com `
  -tags sqli,xss `
  -o /output/results.txt

# 偵測敏感資訊洩漏
docker run --rm laiyuci7/vulnsight:latest \
  -u https://example.com \
  -t http/exposures/ \
  -silent
```

### 情境 5：更新模板庫

```bash
docker run --rm \
  -v vulnsight-templates:/root/nuclei-templates \
  laiyuci7/vulnsight:latest -update-templates
```

---

## 注意事項

### ⚠️ 法律規範

- **只能對你有授權的目標進行掃描**
- 未經授權的掃描屬於違法行為
- 進行滲透測試前，請確認已取得書面授權

### 🔧 顯示全部說明

```bash
docker run --rm laiyuci7/vulnsight:latest -h
```

### 📊 顯示即時掃描統計

```bash
docker run --rm laiyuci7/vulnsight:latest \
  -u https://example.com -s critical,high -stats
```

### 🔍 Debug 模式（除錯用）

```bash
# 顯示所有請求與回應
docker run --rm laiyuci7/vulnsight:latest \
  -u https://example.com -debug

# 顯示詳細輸出
docker run --rm laiyuci7/vulnsight:latest \
  -u https://example.com -v
```
