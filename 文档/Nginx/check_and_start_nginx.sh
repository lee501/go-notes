#!/bin/bash
# �鿴nginx�����Ƿ��������У�Ϊ0���ʾ�Ѿ�down��
nginx_result=$(ps -C nginx --no-heading|wc -l)
echo nginx_result > /app/demo/a.txt
if [ "${nginx_result}" = "0" ]; then
	# ǰ����Ҫ��nginx����Ϊservice�� ����ֱ��ʹ��nginx���ھ���·�����磺 /usr/local/nginx/sbin/nginx
    echo '��������' > /app/demo.a.txt
	service nginx start
    sleep 1
    nginx_result=$(ps -C nginx --no-heading|wc -l)
    if [ "${nginx_result}" = "0" ]; then
		# �������nginx�����ǲ��еĻ����Ͱ�keepalivedҲֹͣ��
		# ��������ʱ�keepalived���Ӷ���֤����nginx�ҵ���Ҳ��ʹ��
        /etc/rc.d/init.d/keepalived stop
    fi
fi
