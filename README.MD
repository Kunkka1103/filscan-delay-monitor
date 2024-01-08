# Filscan Delay Monitor

This is a simple monitoring tool for Filscan.

## Features

- Monitor Filscan main-net and calibration
- Send metrics to a Pushgateway

## Usage

```text
Usage of ./filscan-delay-monitor:
  -calibration-url string
        Calibration url (default "https://calibration.filscan.io")
  -interval int
        Interval in minutes (default 5)
  -mainnet-url string
        Mainnet url (default "https://filscan.io")
  -push-address string
        Pushgateway address (default "http://localhost:9091")
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
