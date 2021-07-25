# Health Monitor
A simple tool to monitor server health written in Go and web technologies.

## How does it work?
It stores health-check API response time in Redis storage every minute for each API.
To increase the accuracy, it calls each API every 20 seconds and updates the result stored in the storage.
So even if two of the calls failed, there would be a response time for that minute.

## Installation
```shell
git clone https://github.com/miladrahimi/health-monitor.git
cd health-monitor
cp .env.example .env
docker-compose up -d
docker-compose ps
```

## Configuration
Open `.env` with a text editor and change the environment variables.

Environment variables:
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
* **[Ping Monitor](https://github.com/miladrahimi/ping-monitor)**: A simple tool to monitor server pings

## License
Health Monitor is initially created by [Milad Rahimi](https://miladrahimi.com)
and released under the [MIT License](http://opensource.org/licenses/mit-license.php).
