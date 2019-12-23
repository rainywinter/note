# git




## git url 非标准端口
git支持ssh/git/http/https（也支持ftp，但不推荐），主要有以下几个url格式：
- ssh://[user@]host.xz[:port]/path/to/repo.git/
- git://host.xz[:port]/path/to/repo.git/
- http[s]://host.xz[:port]/path/to/repo.git/
- ftp[s]://host.xz[:port]/path/to/repo.git/

还有另外一种scp风格的语法（An alternative scp-like syntax may also be used with the ssh protocol）：
- [user@]host.xz:path/to/repo.git/

问题来了，当使用ssh协议，且是scp风格的语法时，怎么指定端口？例如：git@xx.yy.com:rainywinter/foo.git,端口号18080

方案1 显式指定ssh协议：git clone ssh://git@xx.yy.com:18080/rainywinter/foo.git
方案2 命令不做修改 git clone git@xx.yy.com:rainywinter/foo.git，但是需要配置ssh config文件

        Host xx.yy.com
	        HostName xx.yy.com
	        User git
	        Port 18080