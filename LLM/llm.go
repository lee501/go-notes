package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// 敏感数据类型定义
type SensitiveDataType string

const (
	PersonalIdentity SensitiveDataType = "personal_identity"
	FinancialInfo    SensitiveDataType = "financial_info"
	HealthInfo       SensitiveDataType = "health_info"
	ContactInfo      SensitiveDataType = "contact_info"
	LocationInfo     SensitiveDataType = "location_info"
	CompanyInfo      SensitiveDataType = "company_info"
	CredentialInfo   SensitiveDataType = "credential_info"
	BiometricInfo    SensitiveDataType = "biometric_info"
)

// 检测结果结构
type DetectionResult struct {
	Type        SensitiveDataType `json:"type"`
	Content     string            `json:"content"`
	StartPos    int               `json:"start_pos"`
	EndPos      int               `json:"end_pos"`
	Confidence  float64           `json:"confidence"`
	Category    string            `json:"category"`
	Description string            `json:"description"`
	Severity    string            `json:"severity"`
	Context     string            `json:"context"`
}

// 大模型API接口
type LLMProvider interface {
	DetectSensitiveData(ctx context.Context, text string) ([]DetectionResult, error)
	GetProviderName() string
}

// OpenAI API实现
type OpenAIProvider struct {
	APIKey  string
	BaseURL string
	Model   string
	Client  *http.Client
}

// OpenAI API请求结构
type OpenAIRequest struct {
	Model       string    `json:"model"`
	Messages    []Message `json:"messages"`
	Temperature float64   `json:"temperature"`
	MaxTokens   int       `json:"max_tokens"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// OpenAI API响应结构
type OpenAIResponse struct {
	Choices []Choice `json:"choices"`
	Usage   Usage    `json:"usage"`
}

type Choice struct {
	Message      Message `json:"message"`
	FinishReason string  `json:"finish_reason"`
}

type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

// 创建OpenAI提供者
func NewOpenAIProvider(apiKey, baseURL, model string) *OpenAIProvider {
	if baseURL == "" {
		baseURL = "https://api.openai.com/v1"
	}
	if model == "" {
		model = "gpt-3.5-turbo"
	}

	return &OpenAIProvider{
		APIKey:  apiKey,
		BaseURL: baseURL,
		Model:   model,
		Client:  &http.Client{Timeout: 30 * time.Second},
	}
}

func (p *OpenAIProvider) GetProviderName() string {
	return "OpenAI"
}

// 构建检测提示词
func (p *OpenAIProvider) buildPrompt(text string) string {
	prompt := `你是一个专业的敏感数据检测专家。请分析以下文本，识别其中的敏感信息，并以JSON格式返回结果。

敏感数据类型包括：
1. personal_identity: 个人身份信息（姓名、身份证号、护照号等）
2. financial_info: 金融信息（银行卡号、支付账号、信用卡等）
3. health_info: 健康信息（病历、诊断、用药等）
4. contact_info: 联系信息（手机号、邮箱、QQ、微信等）
5. location_info: 位置信息（地址、GPS坐标等）
6. company_info: 企业信息（公司名称、组织架构等）
7. credential_info: 凭证信息（密码、密钥、证书等）
8. biometric_info: 生物特征信息（指纹、面部特征等）

请返回JSON数组格式，每个检测结果包含：
- type: 敏感数据类型
- content: 敏感内容
- start_pos: 起始位置（估算）
- end_pos: 结束位置（估算）
- confidence: 置信度（0-1）
- category: 具体分类
- description: 详细描述
- severity: 严重等级（low/medium/high/critical）
- context: 上下文信息

待检测文本：
` + text + `

请只返回JSON数组，不要包含其他文字说明。`

	return prompt
}

// 调用OpenAI API进行检测
func (p *OpenAIProvider) DetectSensitiveData(ctx context.Context, text string) ([]DetectionResult, error) {
	prompt := p.buildPrompt(text)

	reqBody := OpenAIRequest{
		Model: p.Model,
		Messages: []Message{
			{Role: "user", Content: prompt},
		},
		Temperature: 0.1,
		MaxTokens:   2000,
	}

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("序列化请求失败: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", p.BaseURL+"/chat/completions", bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+p.APIKey)

	resp, err := p.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("API请求失败: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API返回错误 %d: %s", resp.StatusCode, string(body))
	}

	var apiResp OpenAIResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, fmt.Errorf("解析响应失败: %w", err)
	}

	if len(apiResp.Choices) == 0 {
		return nil, fmt.Errorf("API响应为空")
	}

	// 解析返回的JSON结果
	content := apiResp.Choices[0].Message.Content
	content = strings.TrimSpace(content)

	// 移除可能的markdown代码块标记
	content = strings.TrimPrefix(content, "```json")
	content = strings.TrimPrefix(content, "```")
	content = strings.TrimSuffix(content, "```")
	content = strings.TrimSpace(content)

	var results []DetectionResult
	if err := json.Unmarshal([]byte(content), &results); err != nil {
		return nil, fmt.Errorf("解析检测结果失败: %w, 原始内容: %s", err, content)
	}

	return results, nil
}

// 本地模型提供者（示例）
type LocalLLMProvider struct {
	ModelPath string
	Client    *http.Client
}

func NewLocalLLMProvider(modelPath string) *LocalLLMProvider {
	return &LocalLLMProvider{
		ModelPath: modelPath,
		Client:    &http.Client{Timeout: 60 * time.Second},
	}
}

func (p *LocalLLMProvider) GetProviderName() string {
	return "LocalLLM"
}

// 本地模型检测实现（使用ollama或其他本地服务）
func (p *LocalLLMProvider) DetectSensitiveData(ctx context.Context, text string) ([]DetectionResult, error) {
	// 这里是本地模型的调用示例（如ollama）
	prompt := p.buildLocalPrompt(text)

	reqBody := map[string]interface{}{
		"model":  "llama2",
		"prompt": prompt,
		"stream": false,
	}

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("序列化请求失败: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", "http://localhost:11434/api/generate", bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := p.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("本地模型请求失败: %w", err)
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("解析响应失败: %w", err)
	}

	response, ok := result["response"].(string)
	if !ok {
		return nil, fmt.Errorf("无效的响应格式")
	}

	// 解析检测结果
	return p.parseLocalResponse(response)
}

func (p *LocalLLMProvider) buildLocalPrompt(text string) string {
	return fmt.Sprintf(`作为敏感数据检测专家，请分析以下文本中的敏感信息：

文本：%s

请识别并返回JSON格式的检测结果。`, text)
}

func (p *LocalLLMProvider) parseLocalResponse(response string) ([]DetectionResult, error) {
	// 简化的解析逻辑
	var results []DetectionResult

	// 这里需要根据实际的本地模型输出格式进行解析
	// 示例：假设模型返回的是JSON格式
	if strings.Contains(response, "[") && strings.Contains(response, "]") {
		start := strings.Index(response, "[")
		end := strings.LastIndex(response, "]") + 1
		jsonStr := response[start:end]

		if err := json.Unmarshal([]byte(jsonStr), &results); err != nil {
			return nil, fmt.Errorf("解析本地模型响应失败: %w", err)
		}
	}

	return results, nil
}

// 敏感数据检测器
type LLMSensitiveDetector struct {
	provider            LLMProvider
	batchSize           int
	maxRetries          int
	retryDelay          time.Duration
	confidenceThreshold float64
}

// 创建LLM检测器
func NewLLMSensitiveDetector(provider LLMProvider) *LLMSensitiveDetector {
	return &LLMSensitiveDetector{
		provider:            provider,
		batchSize:           5000, // 单次处理的最大文本长度
		maxRetries:          3,
		retryDelay:          time.Second,
		confidenceThreshold: 0.5,
	}
}

// 设置置信度阈值
func (d *LLMSensitiveDetector) SetConfidenceThreshold(threshold float64) {
	d.confidenceThreshold = threshold
}

// 检测敏感数据
func (d *LLMSensitiveDetector) DetectSensitiveData(ctx context.Context, text string) ([]DetectionResult, error) {
	if len(text) == 0 {
		return []DetectionResult{}, nil
	}

	var allResults []DetectionResult

	// 如果文本太长，分批处理
	if len(text) > d.batchSize {
		chunks := d.splitText(text, d.batchSize)
		offset := 0

		for _, chunk := range chunks {
			results, err := d.detectWithRetry(ctx, chunk)
			if err != nil {
				return nil, fmt.Errorf("检测文本块失败: %w", err)
			}

			// 调整位置偏移
			for i := range results {
				results[i].StartPos += offset
				results[i].EndPos += offset
			}

			allResults = append(allResults, results...)
			offset += len(chunk)
		}
	} else {
		var err error
		allResults, err = d.detectWithRetry(ctx, text)
		if err != nil {
			return nil, err
		}
	}

	// 过滤低置信度结果
	return d.filterByConfidence(allResults), nil
}

// 带重试的检测
func (d *LLMSensitiveDetector) detectWithRetry(ctx context.Context, text string) ([]DetectionResult, error) {
	var lastErr error

	for i := 0; i < d.maxRetries; i++ {
		results, err := d.provider.DetectSensitiveData(ctx, text)
		if err == nil {
			return results, nil
		}

		lastErr = err
		if i < d.maxRetries-1 {
			select {
			case <-time.After(d.retryDelay):
			case <-ctx.Done():
				return nil, ctx.Err()
			}
		}
	}

	return nil, fmt.Errorf("重试%d次后仍然失败: %w", d.maxRetries, lastErr)
}

// 分割文本
func (d *LLMSensitiveDetector) splitText(text string, maxSize int) []string {
	var chunks []string
	runes := []rune(text)

	for i := 0; i < len(runes); i += maxSize {
		end := i + maxSize
		if end > len(runes) {
			end = len(runes)
		}
		chunks = append(chunks, string(runes[i:end]))
	}

	return chunks
}

// 按置信度过滤结果
func (d *LLMSensitiveDetector) filterByConfidence(results []DetectionResult) []DetectionResult {
	var filtered []DetectionResult

	for _, result := range results {
		if result.Confidence >= d.confidenceThreshold {
			filtered = append(filtered, result)
		}
	}

	return filtered
}

// 生成检测报告
func (d *LLMSensitiveDetector) GenerateReport(results []DetectionResult, originalText string) string {
	report := fmt.Sprintf("=== 基于大模型的敏感数据检测报告 ===\n")
	report += fmt.Sprintf("检测提供者: %s\n", d.provider.GetProviderName())
	report += fmt.Sprintf("检测时间: %s\n", time.Now().Format("2006-01-02 15:04:05"))
	report += fmt.Sprintf("文本长度: %d 字符\n", len(originalText))
	report += fmt.Sprintf("检测到敏感项: %d 个\n", len(results))
	report += fmt.Sprintf("置信度阈值: %.2f\n\n", d.confidenceThreshold)

	// 按严重程度分类统计
	severityCount := make(map[string]int)
	typeCount := make(map[SensitiveDataType]int)

	for _, result := range results {
		severityCount[result.Severity]++
		typeCount[result.Type]++
	}

	report += "风险等级分布:\n"
	for severity, count := range severityCount {
		report += fmt.Sprintf("  %s: %d 项\n", severity, count)
	}

	report += "\n类型分布:\n"
	for dataType, count := range typeCount {
		report += fmt.Sprintf("  %s: %d 项\n", dataType, count)
	}

	if len(results) > 0 {
		report += "\n检测详情:\n"
		for i, result := range results {
			report += fmt.Sprintf("%d. [%s] %s\n", i+1, result.Type, result.Category)
			report += fmt.Sprintf("   内容: %s\n", result.Content)
			report += fmt.Sprintf("   位置: %d-%d\n", result.StartPos, result.EndPos)
			report += fmt.Sprintf("   置信度: %.2f\n", result.Confidence)
			report += fmt.Sprintf("   严重程度: %s\n", result.Severity)
			report += fmt.Sprintf("   描述: %s\n", result.Description)
			if result.Context != "" {
				report += fmt.Sprintf("   上下文: %s\n", result.Context)
			}
			report += "\n"
		}
	}

	return report
}

// 批量检测
func (d *LLMSensitiveDetector) BatchDetect(ctx context.Context, texts []string) ([][]DetectionResult, error) {
	results := make([][]DetectionResult, len(texts))

	for i, text := range texts {
		result, err := d.DetectSensitiveData(ctx, text)
		if err != nil {
			return nil, fmt.Errorf("批量检测第%d个文本失败: %w", i+1, err)
		}
		results[i] = result
	}

	return results, nil
}

// 主函数示例
func main() {
	// 创建OpenAI提供者（需要设置API密钥）
	provider := NewOpenAIProvider("your-api-key", "", "gpt-3.5-turbo")

	// 创建检测器
	detector := NewLLMSensitiveDetector(provider)
	detector.SetConfidenceThreshold(0.7)

	// 测试文本
	testTexts := []string{
		"用户张三的手机号是13812345678，身份证号是330106199001011234，居住在杭州市西湖区文三路123号",
		"我的银行卡号是6222024200012345678，密码是123456，请帮我转账到支付宝账号zhangsan@alipay.com",
		"病人李四，男，45岁，诊断为2型糖尿病，血糖值8.5mmol/L，需要每日注射胰岛素",
		"ABC科技有限公司的财务经理王五，工作邮箱wangwu@abc.com，办公电话010-12345678",
	}

	ctx := context.Background()

	// 单个文本检测示例
	fmt.Println("=== 单个文本检测示例 ===")
	testText := testTexts[0]
	fmt.Printf("检测文本: %s\n\n", testText)

	results, err := detector.DetectSensitiveData(ctx, testText)
	if err != nil {
		fmt.Printf("检测失败: %v\n", err)
		return
	}

	report := detector.GenerateReport(results, testText)
	fmt.Println(report)

	// 批量检测示例
	fmt.Println("\n=== 批量检测示例 ===")
	batchResults, err := detector.BatchDetect(ctx, testTexts)
	if err != nil {
		fmt.Printf("批量检测失败: %v\n", err)
		return
	}

	for i, results := range batchResults {
		fmt.Printf("文本 %d 检测结果: %d 个敏感项\n", i+1, len(results))
		for _, result := range results {
			fmt.Printf("  - [%s] %s (置信度: %.2f)\n", result.Type, result.Content, result.Confidence)
		}
	}

	// 使用本地模型的示例（需要本地部署ollama或其他服务）
	fmt.Println("\n=== 本地模型检测示例 ===")
	localProvider := NewLocalLLMProvider("llama2")
	localDetector := NewLLMSensitiveDetector(localProvider)

	// 注意：这个示例需要本地运行ollama服务
	// localResults, err := localDetector.DetectSensitiveData(ctx, testText)
	// if err != nil {
	// 	fmt.Printf("本地模型检测失败: %v\n", err)
	// } else {
	// 	fmt.Printf("本地模型检测到 %d 个敏感项\n", len(localResults))
	// }

	_ = localDetector // 避免未使用变量警告
}
