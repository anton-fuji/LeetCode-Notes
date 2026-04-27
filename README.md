# LeetCode Solutions in Go

## Stats

| Difficulty | Count |
|------------|-------|
| Easy       | 10    |
| Medium     | 0     |
| Hard       | 0     |

## Problems

<!-- AUTO-GENERATED: DO NOT EDIT BELOW -->
| # | Title | Difficulty | Tags |
|---|-------|------------|------|
| 100 | [Same Tree](problems/0100-same-tree/) | Easy | DFS |
| 101 | [Symmetric Tree](problems/0101-symmetric-tree/) | Easy | DFS |
| 104 | [Maximum Depth of Binary Tree](problems/0104-max-depth/) | Easy | DFS |
| 66 | [Plus One](problems/0066-plus-one/) | Easy | array, math |
| 67 | [Add Binary](problems/0067-add-binary/) | Easy |  |
| 69 | [Sqrt(x)](problems/0069-sqrtx/) | Easy | Binary Search |
| 70 | [Climbing Stairs](problems/0070-climbing-stairs/) | Easy | Fibonacci |
| 83 | [Remove Duplicates from Sorted List](problems/0083-remove-duplicates/) | Easy | Singly Linked List |
| 88 | [Merge Sorted Arrary](problems/0088-merge-sorted/) | Easy | Merge |
| 94 | [Binary Tree Inorder Traversal](problems/0094-inorder/) | Easy | DFS |
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
