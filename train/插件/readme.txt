install vim-go vim-jsbeautity(debian x64)



--------
install
--------

scp vimrc.tar.gz shu@192.168.0.117:/tmp //复制此vimrc.tar.gz至服务器

ssh shu@192.168.0.117 // 登录服务器或电脑

cd ~

mv /tmp/vimrc.tar.gz

mkdir vim.bak // 备份原vimrc文件

mv .vimrc vim.bak(若存在)

mv .vim vim.bak(若存在)

mv bin bin.bak(若存在)
tar -xzf vim.tar.gz

[sudo]cp bin/* $GOROOT/bin 



// 重新打开vim即可使用以下特性
----------------vim-go key
-------------------

打开go文件，光标对着go方法:
"\dv"，打开关联方法的文件(ctrl+w跳转窗口)。
"\e",重命名文法。

----------------
vim-jsbeautity key
-------------------

打开html,javacript,css文件，使用"ctrl+f" 可格式化源码.




参考资料

https://github.com/fatih/vim-go

http://studygolang.com/articles/1785

