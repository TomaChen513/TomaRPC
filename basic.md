邮箱号：

fixed-term.Xinyu.CHEN2@cn.bosch.com

数据库

boschIntern         NLIBosch~xinyu

服务器

BoschIntern        Boschpro321!

Linux常用命令：
* ps -ef | grep extra   查看具体进程号
* df -h  查看磁盘情况
* cp file1 file2 复制文件

vim批量替换

:%s/source_pattern/target_pattern/g

docker复制

docker cp ef41d:/home/ /home/mysqlBackup/

1.导出结构不导出数据

复制代码代码如下:

mysqldump　--opt　-d　数据库名　-u　root　-p　>　xxx.sql

2.导出数据不导出结构

复制代码代码如下:

mysqldump　-t　数据库名　-uroot　-p　>　xxx.sql

3.导出数据和表结构

复制代码代码如下:

mysqldump　数据库名　-uroot　-p　>　xxx.sql
