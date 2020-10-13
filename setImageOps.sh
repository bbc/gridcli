#!/usr/bin/env bash
red='\x1B[0;31m'
plain='\x1B[0m' # No Color
 if [ ! -f ./imgops/dev/nginx.conf ]; then
        bucket=$(bash get-stack-resource.sh ImageBucket)
        if [ -z "$bucket" ]; then
            echo -e "${red}[CANNOT GET ImageBucket] This may be because your default region for the media-service profile has not been set.${plain}"
            exit 1
        fi

        sed -e 's/{{BUCKET}}/'"${bucket}"'/g' ./imgops/dev/nginx.conf.template > ./imgops/dev/nginx.conf
    fi