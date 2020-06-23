elasticsearch6.8 安装指南

注意事项：
1. 使用docker-compose安装时, 请将 `elasticsearch.yml` 文件copy至 `/docker/elasticsearch6.8/config/` 目录下
2. 设置虚拟内存区数量
```
echo "vm.max_map_count=655300" > /etc/sysctl.conf
sysctl -p
```