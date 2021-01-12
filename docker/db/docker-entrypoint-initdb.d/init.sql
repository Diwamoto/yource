grant all privileges on *.* to user@"%" identified by 'password' with grant option;
FLUSH PRIVILEGES;
SET GLOBAL sql_mode = '';