FROM postgres

COPY scripts/* /tmp

# скачивание скрипта заполнения бд
RUN apt-get update && apt-get install wget -y
RUN wget "https://drive.google.com/u/0/uc?id=1gG5fi8OEA7z9ZwzKho4jrKfmPecUdi9P&export=download" -O /tmp/fill_db.sql

# объединение скриптов создания и заполнения бд
# скрипт в директории docker-entrypoint-initdb.d выполнится при запуске контейнера автоматически 
RUN cat /tmp/create_db.sql /tmp/fill_db.sql > /docker-entrypoint-initdb.d/init.sql



