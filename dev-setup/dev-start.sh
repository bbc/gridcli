#!/usr/bin/env bash
set -eou
CONF_FILE=$(pwd)/dev-setup/application.conf
LOGBACK_FILE=$(pwd)/dev-setup/logback.xml
GRID_DIR=/Users/fadhilijuma/bbc/grid
red='\x1B[0;31m'
plain='\x1B[0m' # No Color

setupImgops() {
    if [ ! -f ./imgops/dev/nginx.conf ]; then
        bucket=$(bash get-stack-resource.sh ImageBucket)
        if [ -z "$bucket" ]; then
            echo -e "${red}[CANNOT GET ImageBucket] This may be because your default region for the media-service profile has not been set.${plain}"
            exit 1
        fi

        sed -e 's/{{BUCKET}}/'"${bucket}"'/g' ./imgops/dev/nginx.conf.template > ./imgops/dev/nginx.conf
    fi

    netLocation=$(networksetup -getcurrentlocation)
    if [[ $netLocation = "BBC On Network" ]]; then
        ## Use Reith resolver for Nginx
        sed -i'.bak' 's/resolver.*/resolver 10.162.8.1;/g' imgops/dev/nginx.conf
    else
        ## Use Default resolver
        sed -i'.bak' 's/resolver.*/resolver 8.8.8.8 8.8.4.4;/g' imgops/dev/nginx.conf
    fi
    rm imgops/dev/nginx.conf.bak
}

startDockerContainers() {
    docker-compose up -d imgops
}

buildJs() {
  pushd kahuna || echo "cannot find path kahuna" && exit 1
  npm install
  npm run build-dev
  popd kahuna || echo "cannot find path kahuna" && exit 1
}

startPlayApps() {
    if [ "$IS_DEBUG" = true ] ; then
        sbt -jvm-debug 5005 "runAll -Dconfig.file=$CONF_FILE -Dlogger.file=$LOGBACK_FILE"
    else
        sbt "runAll -Dconfig.file=$CONF_FILE -Dlogger.file=$LOGBACK_FILE"
    fi
}

# We use auth.properties as a proxy for whether all the configuration files have been downloaded given the implementation of `fetchConfig.sh`.
downloadApplicationConfig() {
    if [ ! -f /etc/gu/auth.properties ]; then
        bash ./fetch-config.sh
    fi
}

checkGridDir() {
    if [ -z "$GRID_DIR" ]; then
        echo -e "It looks like you haven't set the directory for where the Grid source code is set yet. Store this in your bash profile with: export GRID_DIR=<dir/path> and then reload your bash profile."
        exit 1
    fi
}

# Runs pingler in background without printing output
startPinging() {
    echo ""
    echo "Running pingler in the background"
    echo ""
    ./the_pingler.sh > /dev/null &
}

main() {
    for i in "$@" ; do
        if [[ $i == "js" ]] ; then
             BUILD_JS=true
        fi
        if [[ $i == "debug" ]] ; then
             IS_DEBUG=true
        fi
    done

    checkGridDir
    echo "Starting The Grid using location: $GRID_DIR and using conf file: $CONF_FILE"
    pushd "${GRID_DIR}" || echo "cannot find path ${GRID_DIR}" && exit 1
    startPinging
    setupImgops
    downloadApplicationConfig
    startDockerContainers

    if [[ "$BUILD_JS" = true ]] ; then
        echo "Building JavaScript"
        buildJs
    fi

    startPlayApps
    popd || echo "cannot find path ${GRID_DIR}" && exit 1
}

main "$@"
kill $! # Needed to kill pingler