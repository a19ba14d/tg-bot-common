#!/bin/bash

# 检查是否提供了标签参数
if [ $# -eq 0 ]; then
    echo "错误: 请提供标签名称"
    echo "用法: ./update_tag.sh <标签名称> [提交哈希]"
    exit 1
fi

TAG_NAME=$1
COMMIT_HASH=$2

# 如果没有提供提交哈希，使用当前 HEAD
if [ -z "$COMMIT_HASH" ]; then
    COMMIT_HASH=$(git rev-parse HEAD)
fi

echo "开始更新标签: $TAG_NAME"

# 检查提交哈希是否有效
if ! git rev-parse --verify "$COMMIT_HASH" >/dev/null 2>&1; then
    echo "错误: 无效的提交哈希: $COMMIT_HASH"
    exit 1
fi

# 删除本地标签
echo "删除本地标签..."
git tag -d "$TAG_NAME" 2>/dev/null || true

# 删除远程标签
echo "删除远程标签..."
git push origin ":refs/tags/$TAG_NAME" 2>/dev/null || true

# 创建新标签
echo "在提交 $COMMIT_HASH 上创建新标签..."
git tag "$TAG_NAME" "$COMMIT_HASH"

# 推送新标签
echo "推送新标签到远程仓库..."
if git push origin "$TAG_NAME"; then
    echo "标签 $TAG_NAME 更新成功！"
else
    echo "错误: 推送标签失败"
    exit 1
fi 