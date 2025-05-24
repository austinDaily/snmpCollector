# SNMP Data Collector

A simple Go CLI tool that collects basic SNMP data (such as hostname, uptime, and system description) from a network device like a router or switch.

---

## 📦 Features

- Connects to a device using SNMP v2c.
- Queries standard OIDs like:
  - `sysDescr` (System Description)
  - `sysUpTime` (System Uptime)
  - `sysName` (Device Hostname)
- Clean command-line output.

---

