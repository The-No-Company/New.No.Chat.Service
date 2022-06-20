#!/bin/sh

#crontab -u www-data /var/www/cron/tasks
#crontab -u www-data -l
#crond -b -l 8
/go/bin/CompileDaemon --build="go build main.go" --command=./main