FROM nginx:1.18 AS NGINX


FROM golang:1.18 AS GOLANG

# 安装nginx
COPY --from=NGINX /usr/sbin/nginx /usr/sbin/nginx
COPY --from=NGINX /etc/nginx /etc/nginx
COPY --from=NGINX /usr/share/nginx/html /usr/share/nginx/html
COPY --from=NGINX /var/log/nginx /var/log/nginx
COPY --from=NGINX /var/cache/nginx /var/cache/nginx
# 切换用户
RUN sed -i 's/user  nginx;/user root;/g' /etc/nginx/nginx.conf


ENV TZ=Asia/Shanghai
ENV GOPROXY=https://goproxy.cn

COPY entrypoint.sh /entrypoint.sh

RUN chmod +x /entrypoint.sh && \
    echo 'export PS1="\e[1m\e[31m[\h] \e[32m(rango) \e[34m\u@$(hostname -i)\e[35m \w\e[0m\n$ "' >> ~/.bashrc

WORKDIR /app

ENTRYPOINT ["/entrypoint.sh"]

EXPOSE 80
EXPOSE 443

CMD ["nginx"]
