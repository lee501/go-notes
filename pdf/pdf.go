package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

type MetaResult struct {
	meta map[string]string
	err  error
}

type BodyResult struct {
	body string
	err  error
}

// 检查必要的命令是否存在
func checkCommands() error {
	commands := []string{"pdfinfo", "pdftotext"}
	for _, cmd := range commands {
		if _, err := exec.LookPath(cmd); err != nil {
			return fmt.Errorf("command '%s' not found: %v", cmd, err)
		}
	}
	return nil
}

// 检查文件是否存在和可读
func checkFile(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return fmt.Errorf("file does not exist: %s", path)
	}

	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("cannot open file: %v", err)
	}
	file.Close()

	return nil
}

// Convert PDF with better error handling
func ConvertPDFText(path string) (BodyResult, MetaResult, error) {
	// 预检查
	if err := checkCommands(); err != nil {
		return BodyResult{}, MetaResult{}, err
	}

	if err := checkFile(path); err != nil {
		return BodyResult{}, MetaResult{}, err
	}

	metaResult := MetaResult{meta: make(map[string]string)}
	bodyResult := BodyResult{}
	mr := make(chan MetaResult, 1)

	go func() {
		defer func() { mr <- metaResult }()

		cmd := exec.Command("pdfinfo", path)
		metaStr, err := cmd.Output()
		if err != nil {
			// 获取详细错误信息
			if exitError, ok := err.(*exec.ExitError); ok {
				metaResult.err = fmt.Errorf("pdfinfo failed with exit status %d: %s",
					exitError.ExitCode(), string(exitError.Stderr))
			} else {
				metaResult.err = fmt.Errorf("pdfinfo failed: %v", err)
			}
			return
		}

		// Parse meta output
		for _, line := range strings.Split(string(metaStr), "\n") {
			if parts := strings.SplitN(line, ":", 2); len(parts) > 1 {
				metaResult.meta[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
			}
		}

		// Convert parsed meta
		if x, ok := metaResult.meta["ModDate"]; ok {
			if t, ok := pdfTimeLayouts.Parse(x); ok {
				metaResult.meta["ModifiedDate"] = fmt.Sprintf("%d", t.Unix())
			}
		}
		if x, ok := metaResult.meta["CreationDate"]; ok {
			if t, ok := pdfTimeLayouts.Parse(x); ok {
				metaResult.meta["CreatedDate"] = fmt.Sprintf("%d", t.Unix())
			}
		}
	}()

	br := make(chan BodyResult, 1)
	go func() {
		defer func() { br <- bodyResult }()

		cmd := exec.Command("pdftotext", "-q", "-nopgbrk", "-enc", "UTF-8", "-eol", "unix", path, "-")
		body, err := cmd.Output()
		if err != nil {
			// 获取详细错误信息
			if exitError, ok := err.(*exec.ExitError); ok {
				bodyResult.err = fmt.Errorf("pdftotext failed with exit status %d: %s",
					exitError.ExitCode(), string(exitError.Stderr))
			} else {
				bodyResult.err = fmt.Errorf("pdftotext failed: %v", err)
			}
			return
		}

		bodyResult.body = string(body)
	}()

	bodyRes := <-br
	metaRes := <-mr

	// 检查是否有错误
	if bodyRes.err != nil && metaRes.err != nil {
		return bodyRes, metaRes, fmt.Errorf("both operations failed - body: %v, meta: %v", bodyRes.err, metaRes.err)
	}

	return bodyRes, metaRes, nil
}

var pdfTimeLayouts = timeLayouts{
	time.ANSIC,
	"Mon Jan _2 15:04:05 2006 MST",
	"D:20060102150405-07'00'", // PDF standard format
	"D:20060102150405",        // PDF format without timezone
}

type timeLayouts []string

func (tl timeLayouts) Parse(x string) (time.Time, bool) {
	// 处理PDF标准时间格式
	if strings.HasPrefix(x, "D:") {
		x = strings.TrimPrefix(x, "D:")
		// Remove quotes if present
		x = strings.Trim(x, "'")
	}

	for _, layout := range tl {
		t, err := time.Parse(layout, x)
		if err == nil {
			return t, true
		}
	}
	return time.Time{}, false
}

// 测试函数
func main() {
	bodyResult, metaResult, err := ConvertPDFText("example.pdf")

	if err != nil {
		fmt.Printf("General error: %v\n", err)
	}

	if bodyResult.err != nil {
		fmt.Printf("Body extraction error: %v\n", bodyResult.err)
	} else {
		fmt.Printf("Body length: %d characters\n", len(bodyResult.body))
	}

	if metaResult.err != nil {
		fmt.Printf("Meta extraction error: %v\n", metaResult.err)
	} else {
		fmt.Printf("Meta fields: %d\n", len(metaResult.meta))
		for k, v := range metaResult.meta {
			fmt.Printf("  %s: %s\n", k, v)
		}
	}
}
