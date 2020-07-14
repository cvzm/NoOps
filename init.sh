#!/bin/bash

# 该文件用于初始化系统，安装必要组件，保证本服务可开箱即用。

# 必备组件清单
# 1. ansible
# 2. sshpass
# 3. bash-completion


####### Start ######
cat logo.ascii

# 获取发行版
Get_Dist_Name() {
    if grep -Eqii "CentOS" /etc/issue || grep -Eq "CentOS" /etc/*-release; then
        DISTRO='CentOS'
        PM='yum'
    elif grep -Eqi "Debian" /etc/issue || grep -Eq "Debian" /etc/*-release; then
        DISTRO='Debian'
        PM='apt'
    elif grep -Eqi "Ubuntu" /etc/issue || grep -Eq "Ubuntu" /etc/*-release; then
        DISTRO='Ubuntu'
        PM='apt'
    elif grep -Eqi "Raspbian" /etc/issue || grep -Eq "Raspbian" /etc/*-release; then
        DISTRO='Raspbian'
        PM='apt'
    else
        DISTRO='unknow'
    fi
    echo $DISTRO;
    echo $PM;
}
Get_Dist_Name >/dev/null

# 具体业务逻辑
case $DISTRO in
    Ubuntu)
        echo todo ubuntu...
        ;;
    CentOS)
        echo os is $DISTRO!, Start init...
        $PM install -y epel-release
        $PM install -y ansible sshpass
        echo "NoOps system init OK!"
        ;;
    *)
        echo unknow os $DISTRO, exit! 
        echo "NoOps system init Error!"
        ;;
esac
