# svn 使用笔记






## revert to revision
回滚版本库
1.  更新到最新：svn update。记录最新的版本号v_latest
2.  使用svn log命令找到需要回溯的版本号。记为v_target
3.  回溯版本：svn merge -r v_latest:v_target . （注意最后的点表示当前路径）
4.  确认无误后svn commit提交。


## 恢复到指定版本（working copy）

svn update revision