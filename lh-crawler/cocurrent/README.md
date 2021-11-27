# 并发爬虫

```
├── engine 处理器
├── fetcher 下载器
├── model mapping
├── scheduler 调度器
└── zhenai 真爱网
    └── parser 解析器
```

```shell
docker pull docker.elastic.co/elasticsearch/elasticsearch:7.15.2
docker run -p 127.0.0.1:9200:9200 -p 127.0.0.1:9300:9300 -e "discovery.type=single-node" docker.elastic.co/elasticsearch/elasticsearch:7.15.2
```