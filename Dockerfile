#Version:0.0.1
#publicIP
FROM registry.cn-hangzhou.aliyuncs.com/shanchain/nodejs-image
#FROM registry-internal.cn-hangzhou.aliyuncs.com/shanchain/nodejs-image
MAINTAINER snow "xuefeng.zhao@shanchain.com"
ADD logger /logger
CMD ["node", "/logger/bin/www", "daemon off;"]
EXPOSE 80