想测试一下 `git branch -f` 的用法


过程

1. 我在 `git-branch--f` 分支 试了一下 `git branch -f git-branch--f-branched`，发现没变化，但是查看分支会发现多了一个分支。
2. 上网搜索了一下，发现需要两个参数，`git branch -f [branch_name] [object_pos]`，解释是 **移动分支到目标位置**。
3. 在 `git-branch--f` 分支又提交了一次之后， `git-branch--f` 分支已经有 2 个 commit 了，而 `git-branch--f-branched` 分支只有一个 commit。于是我尝试交换两个分支的进度。
4. 两个 commit 的 ID 分别为 `f1406bdd` `b1d777f5`
5. 第一次直接交换，想让落后分支先切到最新的提交上，发现报错，`fatal: Cannot force update the current branch.`。所以只能先切到其他分支上进行操作，还是直接交换，但是不知道为啥看不到结果。所以就 checkout 到其中一个分支上，发现两个都在最新的分支上。这应该是说明了落后分支确实移到最新的提交上了。
6. 然后尝试将 `git-branch--f` 分支移到最早的提交上，因为最早的提交已经没有分支在那里了，所以只能用提交的 ID。
7. 总结一下，其实只要有 ID，就可以直接切换，只要当前不是在修改的分支上：
8. `git branch -f git-branch--f f1406bdd`
9. `git branch -f git-branch--f-branched b1d777f5`
10. 这样就可以交换了两个分支了


参考

* [git branch -f的作用_滕青山YYDS的博客-CSDN博客](https://blog.csdn.net/qq_34626094/article/details/120381912)