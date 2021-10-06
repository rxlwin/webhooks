# 到项目目录，执行 pull
cd /www/wwwroot/pingo/test/
git pull origin 1.0
# 找到程序进程
arr=($(netstat -ltnp | grep 8080))
pidstr=${arr[6]}
pidid=(${pidstr//\// })
pidid=${pidid[0]}
# 杀死进程,然后重新运行
(kill $pidid|go run main.go&)