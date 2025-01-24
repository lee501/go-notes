### 一、基础安全配置
#### 1. 隐藏版本信息
```nginx
  http {
    # 关闭在响应头中显示Nginx版本号
    # 默认响应头: Server: nginx/1.18.0
    # 关闭后响应头: Server: nginx
    server_tokens off;
   }
```
#### 2. 配置安全Headers
```nginx
# 防止网站被嵌入恶意网页中，避免点击劫持
add_header X-Frame-Options "SAMEORIGIN";

# 启用浏览器XSS防护功能，并在检测到攻击时，停止渲染页面
add_header X-XSS-Protection "1; mode=block";

# 禁止浏览器猜测（嗅探）资源的MIME类型，防止资源类型混淆攻击
add_header X-Content-Type-Options "nosniff";

# 控制引用地址信息的传递，增强隐私保护
add_header Referrer-Policy "strict-origin-origin-when-cross-origin";

# 内容安全策略，控制资源加载来源，防止XSS等攻击
# default-src 'self': 只允许加载同源资源
# http: https:: 允许通过HTTP和HTTPS加载资源
# data:: 允许data:URI的资源（如base64编码的图片）
# blob:: 允许blob:URI的资源（如视频流）
# 'unsafe-inline': 允许内联脚本和样式（根据需要配置）
add_header Content-Security-Policy "default-src 'self' http: https: data: blob: 'unsafe-inline'";
```

### 二、访问控制优化
#### 1. 限制连接数
```nginx
http {
    # 定义一个共享内存区域，用于存储IP连接数信息
    # $binary_remote_addr: 使用二进制格式存储客户端IP，节省空间
    # zone=addr:10m: 指定共享内存区域名称为addr，大小为10MB
    limit_conn_zone $binary_remote_addr zone=addr:10m;
    
    # 限制每个IP同时最多100个连接
    limit_conn addr 100;
    
    # 定义请求频率限制，每个IP每秒最多10个请求
    # rate=10r/s: 每秒10个请求
    limit_req_zone $binary_remote_addr zone=req_zone:10m rate=10r/s;
    
    # 应用请求频率限制，burst=20表示最多允许20个请求排队
    limit_req zone=req_zone burst=20 nodelay;
}
```
#### 2. 配置白名单
```nginx
  location /admin/ {
    # 允许内网IP段访问
    # 192.168.1.0/24: 允许192.168.1.x网段的所有IP
    allow 192.168.1.0/24;
    
    # 允许另一个内网IP段访问
    allow 10.0.0.0/8;
    
    # 拒绝其他所有IP访问
    deny all;
    
    # 开启基础认证
    auth_basic "Restricted Access";
    auth_basic_user_file /etc/nginx/.htpasswd;
}
```

### 三、SSL/TLS安全配置
#### 1. 启用HTTPS
```nginx
server {
    # 监听443端口，启用SSL
    listen 443 ssl;
    
    # 指定SSL证书路径
    ssl_certificate /path/to/cert.pem;
    ssl_certificate_key /path/to/key.pem;
    
    # 将所有HTTP请求重定向到HTTPS
    if ($scheme != "https") {
        return 301 https://$server_name$request_uri;
    }
    
    # 启用HSTS，强制浏览器在指定时间内使用HTTPS访问
    add_header Strict-Transport-Security "max-age=31536000" always;
}
```
#### 2. 优化SSL配置
```nginx
# 只允许TLS 1.2和1.3版本，禁用不安全的SSL和早期TLS版本
ssl_protocols TLSv1.2 TLSv1.3;

# 配置加密套件，按推荐顺序排列
# ECDHE: 使用椭圆曲线密钥交换
# AES-GCM: 使用AES-GCM加密模式
ssl_ciphers ECDHE-ECDSA-AES128-GCM-SHA256:ECDHE-RSA-AES128-GCM-SHA256:ECDHE-ECDSA-AES256-GCM-SHA384:ECDHE-RSA-AES256-GCM-SHA384;

# 优先使用服务器的加密套件
ssl_prefer_server_ciphers on;

# 配置SSL会话缓存，提高性能
# shared:SSL:10m: 所有工作进程共享的缓存，大小为10MB
ssl_session_cache shared:SSL:10m;

# SSL会话超时时间
ssl_session_timeout 10m;

# 启用OCSP Stapling，提供证书状态信息
ssl_stapling on;
ssl_stapling_verify on;
resolver 8.8.8.8 8.8.4.4 valid=300s;
resolver_timeout 5s;
```

### 四、文件上传安全
#### 1. 限制文件上传大小
```nginx
# 限制请求体大小，即上传文件的最大大小为10MB
client_max_body_size 10m;

# 设置请求体缓冲区大小为128KB
# 超过此大小的请求体会被写入临时文件
client_body_buffer_size 128k;

# 配置临时文件存储路径
client_body_temp_path /var/nginx/client_body_temp;
```
#### 2. 配置上传目录权限
```nginx
location /uploads/ {
    # 指定上传根目录
    root /var/www/uploads;
    
    # 指定临时文件目录
    client_body_temp_path /var/www/tmp;
    
    # 允许的WebDAV方法
    dav_methods PUT DELETE MKCOL COPY MOVE;
    
    # 自动创建上传目录
    create_full_put_path on;
    
    # 设置目录访问权限
    # user:rw - 文件所有者可读写
    # group:rw - 组用户可读写
    # all:r - 其他用户只读
    dav_access user:rw group:rw all:r;
    
    # 限制上传文件类型
    if ($request_filename ~* ^.*?\.(php|php5|sh|pl|py)$) {
        return 403;
    }
}
```

### 五、防止常见攻击
#### 1. 防止SQL注入
```nginx
location / {
    # 检查URL中是否包含特殊字符
    # 如果包含分号、单引号、尖括号等字符，返回444状态码
    # 444是Nginx特殊状态码，表示关闭连接而不发送响应头
    if ($request_uri ~* [;'<>] ) {
        return 444;
    }
    
    # 检查查询字符串中的特殊字符
    if ($args ~* [;'<>] ) {
        return 444;
    }
    
    # 保护敏感URI
    location ~* /(admin|backup|config|db|src)/ {
        deny all;
    }
}
```
#### 2. 防止目录遍历
```nginx
# 禁止访问所有以点开头的隐藏文件和目录
location ~ /\. {
    # 拒绝所有请求
    deny all;
    
    # 禁止记录访问日志
    access_log off;
    
    # 禁止记录404错误日志
    log_not_found off;
}

# 禁止访问特定目录
location ~* ^/(uploads|images)/.*\.(php|php5|sh|pl|py|asp|aspx|jsp)$ {
    deny all;
}

# 防止目录列表
location / {
    autoindex off;
}
```

### 六、日志安全
#### 1. 配置访问日志
```nginx
# 定义详细的日志格式
log_format detailed '$remote_addr - $remote_user [$time_local] '  # 记录客户端IP和访问时间
                    '"$request" $status $body_bytes_sent '        # 记录请求信息、状态码和发送字节数
                    '"$http_referer" "$http_user_agent" '        # 记录来源页面和用户代理
                    '$request_time $upstream_response_time';      # 记录请求处理时间和上游响应时间

# 配置访问日志
# buffer=32k: 使用32KB缓冲区
# flush=5s: 每5秒刷新一次日志
access_log /var/log/nginx/access.log detailed buffer=32k flush=5s;

# 对于静态资源，可以关闭访问日志以提高性能
location /static/ {
    access_log off;
}
```
#### 2. 配置错误日志
```nginx

# 设置错误日志级别为warn
# 可选级别: debug, info, notice, warn, error, crit, alert, emerg
error_log /var/log/nginx/error.log warn;

# 对于开发环境，可以使用debug级别获取更多信息
# error_log /var/log/nginx/error.log debug;
```

### 七、其他安全措施
#### 1. 禁止执行脚本
```nginx

location /static/ {
    # 禁止执行PHP文件
    location ~ \.(php|php5)$ {
        deny all;
    }
    
    # 只允许特定文件类型
    location ~* \.(css|js|jpg|jpeg|png|gif|ico|svg|woff|woff2|ttf|eot)$ {
        expires 30d;  # 设置缓存时间
        add_header Cache-Control "public, no-transform";
    }
}
```
#### 2. 配置超时时间
```nginx

# 客户端请求体超时时间，单位秒
client_body_timeout 10;

# 客户端请求头超时时间
client_header_timeout 10;

# 客户端保持连接超时时间
# 第一个参数是客户端超时时间
# 第二个参数是在响应头中的Keep-Alive超时时间
keepalive_timeout 5 5;

# 向客户端发送响应的超时时间
send_timeout 10;

# 读取代理服务器响应的超时时间
proxy_read_timeout 10;

# 连接代理服务器的超时时间
proxy_connect_timeout 10;
```