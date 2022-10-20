备份 MySQL 至阿里云OSS

## 快速使用
若要运行本系统，请在项目根目录下创建 `config.yml` 文件，并配置以下选项：
> 为了安全，请不要将 `config.yml` 文件加入版本控制。

```yaml
DATABASE: "数据库"
MYSQL_HOST: "127.0.0.1"
MYSQL_USERNAME: "root"
MYSQL_PASSWORD: "123456"
BACK_DIR: "./backup/"

CONTAINER__NAME: "MySQL容器名:mysql"

OSS_ACCESS_KEY_ID: "OSS AccessKeyID"
OSS_ACCESS_KEY_SECRET: "OSS AccessKeySecret"
OSS_ENDPOINT: "OSS Endpoint，如：oss-cn-hangzhou.aliyuncs.com"
OSS_BUCKET: "Bucket"


```

执行 `go run . composer.yml`