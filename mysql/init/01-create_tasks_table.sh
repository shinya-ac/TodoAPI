#!/bin/sh

echo "### initialize start ####"
CMD_MYSQL="mysql -u${MYSQL_USER} -p${MYSQL_PASSWORD} ${MYSQL_DATABASE}"
$CMD_MYSQL -e "CREATE TABLE tasks (
    id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    content TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);"
echo "### initialize finish ####"