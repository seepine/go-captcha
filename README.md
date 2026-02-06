# go-captcha

一个基于 [wenlng/go-captcha](https://github.com/wenlng/go-captcha) 的轻量级验证码微服务。提供开箱即用的 HTTP 接口，用于生成滑动拼图验证码，支持 Docker 部署，易于集成到 Node.js、Java、Python 等多种后端系统中。

## ✨ 特性

- **轻量高效**：基于 Go 语言开发，资源占用低，启动速度快。
- **开箱即用**：提供标准 HTTP 接口，无需了解底层图形生成逻辑。
- **易于部署**：支持 Docker 一键启动，也支持二进制直接运行。
- **安全机制**：支持 API Key 鉴权，后端缓存坐标校验，防止暴力破解。
- **滑动拼图**：默认提供滑动拼图验证码，体验友好，安全性高。

## 🚀 快速开始

### 方式一：Docker 启动（推荐）

直接拉取镜像并运行：

```bash
docker run -d \
  -p 8080:8080 \
  --name go-captcha \
  -e API_KEYS=secret-key-1,secret-key-2 \
  seepine/go-captcha
```

### 方式二：本地运行

需要 Go 1.25+ 环境。

```bash
# 克隆项目
git clone https://github.com/seepine/go-captcha.git
cd go-captcha

# 安装依赖
go mod download

# 运行
# 可选设置环境变量：PORT=8080 API_KEYS=mysecret
go run .
```

服务启动后，默认监听在 `:8080` 端口。

## ⚙️ 配置说明

可以通过环境变量配置服务：

| 变量名     | 描述                                                      | 默认值 | 示例            |
| :--------- | :-------------------------------------------------------- | :----- | :-------------- |
| `PORT`     | 服务监听端口                                              | `8080` | `9000`          |
| `API_KEYS` | API 访问密钥，多个密钥用逗号 `,` 分隔。留空则不开启鉴权。 | (空)   | `sk-123,sk-456` |

## 🔌 接口文档

### 生成验证码

获取滑动拼图验证码的背景图、拼图块及坐标数据。

- **接口地址**: `/api/captcha/gen`
- **请求方式**: `GET` 或 `POST`
- **鉴权**: 若配置了 `API_KEYS`，需在 Header 中携带 `ApiKey`。

#### 请求示例

```bash
curl -X POST http://localhost:8080/api/captcha/gen \
  -H "ApiKey: your-secret-key"
```

#### 响应示例

```json
{
  // 前端展示需要的数据
  "slideData": {
    "thumbX": 0, // 拼图块在背景图中的 X 轴起始位置（前端渲染用）
    "thumbY": 0, // 拼图块在背景图中的 Y 轴起始位置
    "thumbWidth": 0, // 拼图块宽度
    "thumbHeight": 0, // 拼图块高度
    "image": "base64...", // 背景图片 Base64
    "thumb": "base64..." // 拼图块图片 Base64
  },
  // 后端验证需要的核心数据（严禁发送给前端）
  "slideVerifyData": {
    "x": 208, // 真实的缺口 X 坐标
    "y": 64 // 真实的缺口 Y 坐标
  }
}
```

## 📖 对接流程

完整的验证码校验流程通常涉及 **前端**、**业务后端** 和 **验证码服务** 三方。

1.  **前端请求验证码**：
    - 用户在前端触发验证码。
    - 前端请求 **你的业务后端** 接口。

2.  **业务后端获取数据**：
    - 你的业务后端调用本服务 `/api/captcha/gen`。
    - 获取到 `slideData` (图片数据) 和 `slideVerifyData` (真实坐标)。

3.  **缓存与下发**：
    - 业务后端生成一个唯一标识 `uniqueId`。
    - 将 `uniqueId` 与 `slideVerifyData` ( `x/y` 坐标) 存入缓存（如 Redis），并设置短暂过期时间（如 5 分钟）。
    - 将 `slideData` 和 `uniqueId` 返回给前端。**注意：千万不要把 `slideVerifyData` 返回给前端！**

4.  **用户交互**：
    - 前端使用 `slideData` 渲染验证码（参考 [wenlng/go-captcha 前端实现](https://github.com/wenlng/go-captcha)）。
    - 用户拖动滑块，前端获取拖动的最终 `x/y` 坐标。

5.  **校验**：
    - 前端将 `uniqueId` 和用户拖动的 `x/y` 坐标提交给 **你的业务后端**。
    - 业务后端根据 `uniqueId` 从缓存取出真实的 `x/y` 坐标。
    - 比较用户提交的 `x/y` 与真实 `x/y`，也可允许一定的误差范围（例如 y 必须一致，x 允许 ±5 像素）
    - 校验通过则执行后续业务逻辑，失败则拒绝。

## 📂 项目结构

```
.
├── common           # 公共逻辑 (验证码初始化、日志等)
├── routes           # HTTP 路由处理
├── utils            # 工具函数
├── main.go          # 程序入口
├── Dockerfile       # Docker 构建文件
└── README.md        # 说明文档
```

## 🛠️ 二次开发

本项目核心依赖 [wenlng/go-captcha](https://github.com/wenlng/go-captcha)。如果你需要增加点击验证码、旋转验证码等功能，可以修改 `common/captcha.go` 和 `routes/captcha.go` 进行扩展。

## 📄 License

MIT
