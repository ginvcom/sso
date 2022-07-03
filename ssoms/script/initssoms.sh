#!/bin/bash
HOSTNAME="localhost"

PORT="13306"

USERNAME="root"

PASSWORD="123456"

insertUser="INSERT INTO user(uuid, avatar, name, mobile, password, birth) Values()"

mysql -h ${HOSTNAME} -P ${PORT} -u ${USERNAME} - p ${PASSWORD} -e "${insertUser}"