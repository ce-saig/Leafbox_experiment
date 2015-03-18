#/bin/bash
lbmysql=sudo docker run --name phpmyadmin-mysql -p 3306 -e MYSQL_ROOT_PASSWORD=password -d mysql
lbmyadmin=sudo docker run -d --link phpmyadmin-mysql:mysql -e MYSQL_USERNAME=root --name phpmyadmin -p 80 corbinu/docker-phpmyadmin
