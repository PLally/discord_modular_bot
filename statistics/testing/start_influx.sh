# start testing database used for stats
docker run -p 8086:8086 -v influxdb:/var/lib/influxdb influxdb