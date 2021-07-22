# Health Monitor
A simple tool to monitor server health written in Go and web technologies.

## How does it work?
It stores health-check api response time in Redis storage every minute for each target (server/endpoint).
To increase the accuracy, it calls every 20 seconds and updates the result stored in the cache.
So even if 2 of the calls get failed, we still have a valid response time for that minute.

## Installation
```shell
git clone https://github.com/miladrahimi/health-monitor.git
cd health-monitor
cp .env.example .env
docker-compose up -d
docker-compose ps
```

## Configuration
Open `.env` with a text editor and change the available variables.

Available variables:
* **APP_EXPOSED_PORT**: The exposed port for web app
* **TARGETS**: The comma-separated list of health-check endpoints to call
* **TIMEZONE**: The timezone!

## Monitoring
Open your browser, surf localhost with the docker exposed port (default: 7575).

The chart is powered by [Chart.js](https://www.chartjs.org)

## Demo
<p align="center">
  <img alt="Demo" src="https://github.com/miladrahimi/health-monitor/blob/master/demo.png?raw=true">
</p>

## See Also
* **[health-monitor](https://github.com/miladrahimi/ping-monitor)**: A simple tool to monitor server pings

## License
PhpRouter is initially created by [Milad Rahimi](https://miladrahimi.com)
and released under the [MIT License](http://opensource.org/licenses/mit-license.php).
