package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"time"
)

// ================================
// 1. 核心接口定义
// ================================

// 通知服务接口
type NotificationService interface {
	SendNotification(userID int, message string) error
	SendNotificationWithType(userID int, notificationType string, data map[string]interface{}) error
}

// 用户信息获取接口
type UserProvider interface {
	GetUser(userID int) (*User, error)
}

// ================================
// 2. 数据模型
// ================================

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type NotificationRequest struct {
	UserID  int                    `json:"user_id"`
	Type    string                 `json:"type"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
	SentAt  time.Time              `json:"sent_at"`
}

// ================================
// 3. 具体通知实现
// ================================

// 邮件通知实现
type EmailNotifier struct {
	smtpHost     string
	smtpPort     string
	smtpUsername string
	smtpPassword string
	userProvider UserProvider
}

func NewEmailNotifier(host, port, username, password string, userProvider UserProvider) *EmailNotifier {
	return &EmailNotifier{
		smtpHost:     host,
		smtpPort:     port,
		smtpUsername: username,
		smtpPassword: password,
		userProvider: userProvider,
	}
}

func (e *EmailNotifier) SendNotification(userID int, message string) error {
	user, err := e.userProvider.GetUser(userID)
	if err != nil {
		return fmt.Errorf("获取用户信息失败: %w", err)
	}

	if user.Email == "" {
		return fmt.Errorf("用户 %d 没有邮箱地址", userID)
	}

	return e.sendEmail(user.Email, "系统通知", message)
}

func (e *EmailNotifier) SendNotificationWithType(userID int, notificationType string, data map[string]interface{}) error {
	user, err := e.userProvider.GetUser(userID)
	if err != nil {
		return fmt.Errorf("获取用户信息失败: %w", err)
	}

	subject, body := e.generateEmailContent(notificationType, user.Name, data)
	return e.sendEmail(user.Email, subject, body)
}

func (e *EmailNotifier) sendEmail(to, subject, body string) error {
	// 构建邮件内容
	msg := fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\n%s", to, subject, body)

	// SMTP认证
	auth := smtp.PlainAuth("", e.smtpUsername, e.smtpPassword, e.smtpHost)

	// 发送邮件
	addr := e.smtpHost + ":" + e.smtpPort
	err := smtp.SendMail(addr, auth, e.smtpUsername, []string{to}, []byte(msg))
	if err != nil {
		return fmt.Errorf("发送邮件失败: %w", err)
	}

	log.Printf("邮件已发送至 %s: %s", to, subject)
	return nil
}

func (e *EmailNotifier) generateEmailContent(notificationType, userName string, data map[string]interface{}) (string, string) {
	switch notificationType {
	case "welcome":
		return "欢迎注册", fmt.Sprintf("亲爱的 %s，欢迎加入我们的平台！", userName)
	case "order_created":
		orderID := data["order_id"]
		amount := data["amount"]
		return "订单创建成功", fmt.Sprintf("亲爱的 %s，您的订单 #%v 已创建，金额: ¥%.2f", userName, orderID, amount)
	case "password_reset":
		token := data["reset_token"]
		return "密码重置", fmt.Sprintf("亲爱的 %s，您的密码重置令牌: %s", userName, token)
	default:
		return "系统通知", fmt.Sprintf("亲爱的 %s，您有新的通知", userName)
	}
}

// 短信通知实现
type SMSNotifier struct {
	apiURL       string
	apiKey       string
	userProvider UserProvider
}

func NewSMSNotifier(apiURL, apiKey string, userProvider UserProvider) *SMSNotifier {
	return &SMSNotifier{
		apiURL:       apiURL,
		apiKey:       apiKey,
		userProvider: userProvider,
	}
}

func (s *SMSNotifier) SendNotification(userID int, message string) error {
	user, err := s.userProvider.GetUser(userID)
	if err != nil {
		return fmt.Errorf("获取用户信息失败: %w", err)
	}

	if user.Phone == "" {
		return fmt.Errorf("用户 %d 没有手机号码", userID)
	}

	return s.sendSMS(user.Phone, message)
}

func (s *SMSNotifier) SendNotificationWithType(userID int, notificationType string, data map[string]interface{}) error {
	user, err := s.userProvider.GetUser(userID)
	if err != nil {
		return fmt.Errorf("获取用户信息失败: %w", err)
	}

	message := s.generateSMSContent(notificationType, user.Name, data)
	return s.sendSMS(user.Phone, message)
}

func (s *SMSNotifier) sendSMS(phone, message string) error {
	// 构建请求数据
	payload := map[string]interface{}{
		"phone":   phone,
		"message": message,
		"api_key": s.apiKey,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("构建请求数据失败: %w", err)
	}

	// 发送HTTP请求
	resp, err := http.Post(s.apiURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("发送短信请求失败: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("短信发送失败，状态码: %d", resp.StatusCode)
	}

	log.Printf("短信已发送至 %s: %s", phone, message)
	return nil
}

func (s *SMSNotifier) generateSMSContent(notificationType, userName string, data map[string]interface{}) string {
	switch notificationType {
	case "welcome":
		return fmt.Sprintf("【系统通知】%s，欢迎注册我们的平台！", userName)
	case "order_created":
		orderID := data["order_id"]
		return fmt.Sprintf("【订单通知】您的订单 #%v 已创建成功", orderID)
	case "verification_code":
		code := data["code"]
		return fmt.Sprintf("【验证码】您的验证码是：%s，请在5分钟内使用", code)
	default:
		return fmt.Sprintf("【系统通知】%s，您有新的通知", userName)
	}
}

// 推送通知实现
type PushNotifier struct {
	fcmServerKey string
	userProvider UserProvider
}

func NewPushNotifier(fcmServerKey string, userProvider UserProvider) *PushNotifier {
	return &PushNotifier{
		fcmServerKey: fcmServerKey,
		userProvider: userProvider,
	}
}

func (p *PushNotifier) SendNotification(userID int, message string) error {
	// 这里简化实现，实际需要获取用户的设备token
	log.Printf("推送通知给用户 %d: %s", userID, message)
	return nil
}

func (p *PushNotifier) SendNotificationWithType(userID int, notificationType string, data map[string]interface{}) error {
	message := p.generatePushContent(notificationType, data)
	return p.SendNotification(userID, message)
}

func (p *PushNotifier) generatePushContent(notificationType string, data map[string]interface{}) string {
	switch notificationType {
	case "welcome":
		return "欢迎加入我们！"
	case "order_created":
		return "您的订单已创建成功"
	case "new_message":
		return "您有新消息"
	default:
		return "您有新通知"
	}
}

// ================================
// 4. 组合通知器 - 支持多渠道同时发送
// ================================

type MultiChannelNotifier struct {
	notifiers []NotificationService
}

func NewMultiChannelNotifier(notifiers ...NotificationService) *MultiChannelNotifier {
	return &MultiChannelNotifier{
		notifiers: notifiers,
	}
}

func (m *MultiChannelNotifier) SendNotification(userID int, message string) error {
	var errors []error

	for _, notifier := range m.notifiers {
		if err := notifier.SendNotification(userID, message); err != nil {
			errors = append(errors, err)
			log.Printf("通知发送失败: %v", err)
		}
	}

	// 如果所有渠道都失败了，返回错误
	if len(errors) == len(m.notifiers) {
		return fmt.Errorf("所有通知渠道都失败了: %v", errors)
	}

	return nil
}

func (m *MultiChannelNotifier) SendNotificationWithType(userID int, notificationType string, data map[string]interface{}) error {
	var errors []error

	for _, notifier := range m.notifiers {
		if err := notifier.SendNotificationWithType(userID, notificationType, data); err != nil {
			errors = append(errors, err)
			log.Printf("通知发送失败: %v", err)
		}
	}

	if len(errors) == len(m.notifiers) {
		return fmt.Errorf("所有通知渠道都失败了: %v", errors)
	}

	return nil
}

// ================================
// 5. 用户服务实现
// ================================

type UserRepository interface {
	Save(user *User) error
	FindByID(id int) (*User, error)
}

// 内存用户仓库实现
type MemoryUserRepository struct {
	users  map[int]*User
	nextID int
}

func NewMemoryUserRepository() *MemoryUserRepository {
	return &MemoryUserRepository{
		users:  make(map[int]*User),
		nextID: 1,
	}
}

func (r *MemoryUserRepository) Save(user *User) error {
	if user.ID == 0 {
		user.ID = r.nextID
		r.nextID++
	}
	r.users[user.ID] = user
	return nil
}

func (r *MemoryUserRepository) FindByID(id int) (*User, error) {
	user, exists := r.users[id]
	if !exists {
		return nil, fmt.Errorf("用户 %d 不存在", id)
	}
	return user, nil
}

func (r *MemoryUserRepository) GetUser(userID int) (*User, error) {
	return r.FindByID(userID)
}

// 用户服务
type UserService struct {
	userRepo UserRepository
	notifier NotificationService // 依赖接口，不是具体实现
}

func NewUserService(userRepo UserRepository, notifier NotificationService) *UserService {
	return &UserService{
		userRepo: userRepo,
		notifier: notifier,
	}
}

func (s *UserService) RegisterUser(name, email, phone string) (*User, error) {
	// 创建用户
	user := &User{
		Name:  name,
		Email: email,
		Phone: phone,
	}

	// 保存到数据库
	if err := s.userRepo.Save(user); err != nil {
		return nil, fmt.Errorf("保存用户失败: %w", err)
	}

	// 发送欢迎通知 - 不关心具体通知方式
	if err := s.notifier.SendNotificationWithType(user.ID, "welcome", map[string]interface{}{
		"user_name": user.Name,
	}); err != nil {
		log.Printf("发送欢迎通知失败: %v", err)
		// 注意：通知失败不应该影响用户注册流程
	}

	return user, nil
}

func (s *UserService) ResetPassword(userID int) error {
	// 生成重置令牌
	resetToken := "abc123xyz789"

	// 发送密码重置通知
	return s.notifier.SendNotificationWithType(userID, "password_reset", map[string]interface{}{
		"reset_token": resetToken,
	})
}

// ================================
// 6. 依赖注入和配置
// ================================

type NotificationConfig struct {
	SMTPHost     string
	SMTPPort     string
	SMTPUsername string
	SMTPPassword string
	SMSApiURL    string
	SMSApiKey    string
	FCMServerKey string
}

func NewNotificationService(config *NotificationConfig, userProvider UserProvider) NotificationService {
	// 创建各种通知器
	emailNotifier := NewEmailNotifier(
		config.SMTPHost,
		config.SMTPPort,
		config.SMTPUsername,
		config.SMTPPassword,
		userProvider,
	)

	smsNotifier := NewSMSNotifier(
		config.SMSApiURL,
		config.SMSApiKey,
		userProvider,
	)

	pushNotifier := NewPushNotifier(
		config.FCMServerKey,
		userProvider,
	)

	// 根据需要选择单一通知器或多渠道通知器
	// 生产环境可能只用邮件
	// return emailNotifier

	// 或者使用多渠道通知
	return NewMultiChannelNotifier(emailNotifier, smsNotifier, pushNotifier)
}

// ================================
// 7. 完整使用示例
// ================================

func main() {
	// 1. 创建用户仓库
	userRepo := NewMemoryUserRepository()

	// 2. 配置通知服务
	config := &NotificationConfig{
		SMTPHost:     "smtp.gmail.com",
		SMTPPort:     "587",
		SMTPUsername: "your-email@gmail.com",
		SMTPPassword: "your-password",
		SMSApiURL:    "https://api.sms-provider.com/send",
		SMSApiKey:    "your-sms-api-key",
		FCMServerKey: "your-fcm-server-key",
	}

	// 3. 创建通知服务
	notificationService := NewNotificationService(config, userRepo)

	// 4. 创建用户服务（依赖注入）
	userService := NewUserService(userRepo, notificationService)

	// 5. 使用示例
	fmt.Println("=== 用户注册示例 ===")
	user, err := userService.RegisterUser("张三", "zhangsan@example.com", "13800138000")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("用户注册成功: %+v\n", user)

	// 6. 密码重置示例
	fmt.Println("\n=== 密码重置示例 ===")
	err = userService.ResetPassword(user.ID)
	if err != nil {
		log.Printf("密码重置通知发送失败: %v", err)
	}

	// 7. 直接使用通知服务
	fmt.Println("\n=== 直接通知示例 ===")
	err = notificationService.SendNotificationWithType(user.ID, "order_created", map[string]interface{}{
		"order_id": "ORD-001",
		"amount":   99.99,
	})
	if err != nil {
		log.Printf("订单通知发送失败: %v", err)
	}

	fmt.Println("示例执行完成")
}

/*
=== 关键实现要点 ===

1. **接口抽象**: NotificationService 接口让用户服务不关心具体通知方式

2. **多种实现**: EmailNotifier, SMSNotifier, PushNotifier 各自实现接口

3. **依赖注入**: 在应用启动时注入具体的通知实现

4. **错误处理**: 通知失败不影响主业务流程

5. **扩展性**: 可以轻松添加新的通知渠道

6. **测试友好**: 可以轻松mock NotificationService进行单元测试

7. **配置驱动**: 可以根据配置选择不同的通知实现

8. **组合模式**: MultiChannelNotifier 支持同时使用多种通知渠道
*/
