#!/bin/sh
#用来导入内容到tmp_master.api,tmp_master.api用于生成swagger的json文件

> script/swagger/tmp_master.api
> script/swagger/tmp_global.api

cat api/cluster.api >> script/swagger/tmp_master.api
echo "" >> script/swagger/tmp_master.api
cat api/container.api >> script/swagger/tmp_master.api
echo "" >> script/swagger/tmp_master.api
cat api/cronjob.api >> script/swagger/tmp_master.api
echo "" >> script/swagger/tmp_master.api
cat api/deployment.api >> script/swagger/tmp_master.api
echo "" >> script/swagger/tmp_master.api
cat api/namespace.api >> script/swagger/tmp_master.api
echo "" >> script/swagger/tmp_master.api
cat api/job.api >> script/swagger/tmp_master.api
echo "" >> script/swagger/tmp_master.api
cat api/node.api >> script/swagger/tmp_master.api
echo "" >> script/swagger/tmp_master.api
cat api/token.api >> script/swagger/tmp_master.api
echo "" >> script/swagger/tmp_master.api
cat api/global.api >> script/swagger/tmp_global.api

sed -i "/^import (/,/^)$/s/.*/\n/g" script/swagger/tmp_master.api
echo -e '\nimport (\n	"tmp_global.api"\n)' >> script/swagger/tmp_master.api