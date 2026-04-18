### 1. 両方のノードが存在しない場合

```go
if p == nil && q == nil {
    return true
}
```
両方ともノードが存在しなければ、一致しているとみなす。

### 2. どちらか一方だけ存在する場合
```go
if p == nil || q == nil {
    return false
}
```
片方だけノードが存在する場合、構造が異なるため不一致。

### 3. ノードの値が異なる場合
```go
if p.Val != q.Val {
    return false
}
```
同じ位置にあるノードでも、値が違えば別の木

### 4. 子ノードを再帰的に比較する
```go
return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
```
- 左の子同士が同じ
- 右の子同士が同じ
この両方を満たす場合のみ「同じ木」と判定する


## 計算量
- 時間計算量
   - O(n) 
   - 全てノードを一回ずつ訪問する

