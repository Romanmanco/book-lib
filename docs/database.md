# Запуск локально бд и проверка статуса
sudo systemctl start mysql
sudo systemctl status mysql

# Пример создания базы данных:
CREATE DATABASE mydatabase CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

# Пример создания пользователя:
CREATE USER 'myuser'@'localhost' IDENTIFIED BY 'mypassword';

# Пример установки разрешения пользователю:
GRANT ALL PRIVILEGES ON mydatabase.* TO 'myuser'@'localhost';
FLUSH PRIVILEGES;

# Пример подключения к бд:
mysql -u myuser -p mydatabase


