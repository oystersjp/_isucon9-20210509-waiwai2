gogo:
	sudo systemctl stop nginx
	sudo systemctl stop isucari.golang.service
	sudo systemctl stop mysql
	sudo truncate --size 0 /var/log/nginx/access.log
	sudo truncate --size 0 /var/log/nginx/error.log
	sudo truncate --size 0 /tmp/mysql-slow.log
	$(MAKE) build
	sudo systemctl start mysql
	sudo systemctl start isucari.golang.service 
	sudo systemctl start nginx
	sleep 6
	$(MAKE) benchmark
build:
	 cd /home/isucon/isucari/webapp/go && go build -o isucari && cd ~/

benchmark:
	ssh isucon-bench "cd /home/isucon/isucari &&  make  benchmark"

alp:
	cd  ../ && sudo cat /var/log/nginx/access.log | alp json -m "/new_items/.+,/items/.+,/users/.+,/transactions/.+,/upload/.+" --sort=sum --format=md
