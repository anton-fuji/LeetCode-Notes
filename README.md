# LeetCode Solutions in Go

## Stats

| Difficulty | Count |
|------------|-------|
| Easy       | 0     |
| Medium     | 0     |
| Hard       | 0     |

## Problems

<!-- AUTO-GENERATED: DO NOT EDIT BELOW -->
| # | Title | Difficulty | Tags |
|---|-------|------------|------|
<!-- AUTO-GENERATED: END -->

## Usage

```bash
# 新しい問題を追加
mise run new -- 42

# README を再生成
mise run readme

# テスト実行
mise run test -- 0001-two-sum

# 全テスト実行
mise run test-all
```

## Setup (Neovim)

Telescope で `problems/` 配下を live_grep すればタグ・難易度で絞り込めます。

```lua
-- キーマップ例
vim.keymap.set("n", "<leader>lc", function()
  require("telescope.builtin").live_grep({ search_dirs = { "problems" } })
end, { desc = "LeetCode grep" })
```
