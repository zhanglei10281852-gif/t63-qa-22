# Bug Reproduction

## 包的性质

当前 test_model_fix 保存的是被测模型修复后的结果源码，不是初始含 Bug 源码。要复现原始缺陷，必须检出下面固定的 parent SHA；不要在当前修复结果源码上期待重新出现修复前失败。生成系统使用的可信验证补丁和完整验证日志仅在本地留存，不提交到结果分支。

## 问题现象

受保护请求已经被取消后，认证流程仍继续查询会话和操作员，并可能返回一个有效身份。请先不要修改代码，定位取消信号在哪一层被隔离，并给出数据库访问证据。

## 含 Bug 版本

- 仓库：zhanglei10281852-gif/t63-qa-22
- 仓库地址：https://github.com/zhanglei10281852-gif/t63-qa-22.git
- parent SHA：eadd38e1e02fd296335e194514305e799781e5e9

## 复现步骤

```bash
git clone -- https://github.com/zhanglei10281852-gif/t63-qa-22.git bug-repro
cd bug-repro
git checkout --detach eadd38e1e02fd296335e194514305e799781e5e9
go test ./internal/service/auth -run TestAuthenticationStopsAfterRequestCancellation -count=1
```

## 双架构完整错误信息

### linux/amd64

- 容器内复现预期退出码：1
- 容器内复现实际退出码：1

stdout：

```text
$ go test ./internal/service/auth -run TestAuthenticationStopsAfterRequestCancellation -count=1
--- FAIL: TestAuthenticationStopsAfterRequestCancellation (0.96s)
    cancellation_test.go:20: cancelled authentication returned principal {ID:operator_000001 Username:admin DisplayName:Administrator Role:administrator Status:active PasswordHash:$2a$10$GgadDA0tlmpvehZo7UyQHOTX7m71CBRJTYpBQsfHzYOddNA4tv42e CreatedAt:2026-08-18 00:00:00 +0000 UTC UpdatedAt:2026-08-18 00:00:00 +0000 UTC}
FAIL
FAIL	sanitation-operations/internal/service/auth	0.961s
FAIL

```

stderr：

```text
warning: internal/service/auth/cancellation_test.go has type 100755, expected 100644
warning: internal/service/auth/service_test.go has type 100755, expected 100644
warning: internal/service/auth/cancellation_test.go has type 100755, expected 100644
warning: internal/service/auth/service_test.go has type 100755, expected 100644

```

### linux/arm64

- 容器内复现预期退出码：1
- 容器内复现实际退出码：1

stdout：

```text
$ go test ./internal/service/auth -run TestAuthenticationStopsAfterRequestCancellation -count=1
--- FAIL: TestAuthenticationStopsAfterRequestCancellation (1.09s)
    cancellation_test.go:20: cancelled authentication returned principal {ID:operator_000001 Username:admin DisplayName:Administrator Role:administrator Status:active PasswordHash:$2a$10$3EM4hEngNN1zKmlMBXad9e6nBhCay0XbepdcMH4dz18Rk8CDoYnd6 CreatedAt:2026-08-18 00:00:00 +0000 UTC UpdatedAt:2026-08-18 00:00:00 +0000 UTC}
FAIL
FAIL	sanitation-operations/internal/service/auth	1.285s
FAIL

```

stderr：

```text
warning: internal/service/auth/cancellation_test.go has type 100755, expected 100644
warning: internal/service/auth/service_test.go has type 100755, expected 100644
warning: internal/service/auth/cancellation_test.go has type 100755, expected 100644
warning: internal/service/auth/service_test.go has type 100755, expected 100644

```

## 通过条件

在触发条件下，定向测试 TestAuthenticationStopsAfterRequestCancellation 应通过，相关包、全量测试、竞态测试和构建检查均通过；回退 gold 唯一修复后定向测试重新失败。
