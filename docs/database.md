# Запуск локально бд и проверка статуса
sudo systemctl start mysql
sudo systemctl status mysql

# Пример подключения к бд:
mysql -u myuser -p

# Пример создания базы данных:
CREATE DATABASE mydatabase CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

# Пример создания пользователя:
CREATE USER 'myuser'@'localhost' IDENTIFIED WITH mysql_native_password BY 'mypassword';

# Пример установки разрешения пользователю:
GRANT ALL PRIVILEGES ON mydatabase.* TO 'myuser'@'localhost';
FLUSH PRIVILEGES;

## Тестовая база данных
CREATE DATABASE test_booklib;
CREATE USER 'myuser'@'%' IDENTIFIED WITH mysql_native_password BY 'mypassword';
GRANT ALL PRIVILEGES ON test_booklib.* TO 'myuser'@'%';
FLUSH PRIVILEGES;


