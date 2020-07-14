#!/bin/bash
#********************************************************************
# Author: C.v.
# Date: 2020-06-29
# Description: 用于将参数当成命令执行
#   主要用于在程序中利用该shell执行命令，直接使用命令字符串，无需拆分
#********************************************************************

# 执行命令
echo "echo shell.sh => " $@
eval $@