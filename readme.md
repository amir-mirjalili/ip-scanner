# IP Scanner Web Application

A full-stack IP scanner system built with **Go** that scans IP ranges, detects active devices, and stores results for management via RESTful APIs.

---

## Features

- Scan a given IP range (CIDR notation) for active devices
- Detect devices’ IP, MAC address, hostname, and OS (when possible)
- Store scan sessions and their results in **MySQL** (via GORM)
- CRUD operations for assets (discovered devices)
- View scan history and detailed scan results
- RESTful API built with Echo framework
- Validation using `go-playground/validator`
- Designed for multi-user session support (optional)

---

## Technical Stack

- **Backend:** Go, Echo framework, GORM ORM
- **Database:** MySQL (can be replaced with other SQL databases)
- **Network scanning:** ICMP Ping and ARP requests (requires elevated permissions)

---

## Setup

### Prerequisites

- Go 1.20+
- MySQL server
- `arping` installed on your system
- (Linux) Set binary capability for raw socket usage:

```bash
sudo setcap cap_net_raw+ep ./ip-scanner
```
