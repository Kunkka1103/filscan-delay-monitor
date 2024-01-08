# Filscan Delay Monitor

This is a simple monitoring tool for Filscan.

## Features

- Monitor Filscan main-net and calibration
- Send metrics to a Pushgateway

## Usage

```text
Usage of ./filscan-delay-monitor:
  -mainnet-url string
        URL for main-net
  -calibration-url string
        URL for calibration
  -interval int
        Interval in minutes to check the URLs (default 1)
  -push-address string
        Address of the Pushgateway to send metrics
```

## supervisor config
```text
[program:filscan-delay-monitor]
command=/opt/filscan-delay-monitor/filscan-delay-monitor  -push-address=<pushgateway_address> -mainnet-url=<mainnet_url> -calibration-url=<calibration_url> -interval=<interval_in_minutes>
directory=/opt/filscan-delay-monitor/
autostart=true
autorestart=true
user=root
numprocs=1
redirect_stderr=true
stdout_logfile=/opt/filscan-delay-monitor/logs/filscan-delay-monitor.log
stderr_logfile=/opt/filscan-delay-monitor/logs/filscan-delay-monitor.log
minfds=655350
minprocs=655350
```
