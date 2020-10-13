#!/usr/bin/env bash
 if [ ! -f /etc/gu/auth.properties ]; then
        bash ./fetch-config.sh
    fi