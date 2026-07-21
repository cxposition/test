# GitHub Actions + ArgoCD 自动化流程说明

## 1. 目录说明
- .github/workflows/go-ci.yaml: Go 代码质量与测试
- .github/workflows/cd-argocd.yaml: 构建镜像、推送 GHCR、更新 Kubernetes 清单
- k8s/base: ArgoCD 监听的部署清单
- argocd/application.yaml: ArgoCD Application 资源

## 2. 流程
1. 开发代码推送到 main/master。
2. GitHub Actions 构建镜像并推送到 GHCR。
3. GitHub Actions 将 k8s/base/deployment.yaml 中镜像 tag 更新为当前 commit SHA 并提交回仓库。
4. ArgoCD 监听到清单变化后自动同步到集群。

## 3. 首次使用要改的地方
1. 修改 argocd/application.yaml 中的 repoURL。
2. 确保仓库默认分支与 targetRevision 一致。
3. 确保 ArgoCD 已安装在集群，并且存在 argocd 命名空间。
4. 如果仓库是私有的，需要在 ArgoCD 中配置仓库访问凭据。

## 4. GHCR 权限说明
- 工作流使用 GITHUB_TOKEN 推送镜像到 ghcr.io。
- 仓库设置里需要允许 GitHub Actions 写 Packages。

## 5. 应用到集群
kubectl apply -f argocd/application.yaml
