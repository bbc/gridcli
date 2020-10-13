#!/usr/bin/env bash
  pushd kahuna || echo "unable to open folder kahuna" && exit
  npm install
  npm run build-dev
  popd || echo "unable to exit folder kahuna" && exit

if [ "$IS_DEBUG" = true ] ; then
        sbt -jvm-debug 5005 runAll
    else
        sbt runAll
    fi