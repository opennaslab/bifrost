#!/usr/bin/env bash

# 1. Start docker service
systemctl start docker

# 2. Start mysql server
service mysql start

# 3. Start bifrost
