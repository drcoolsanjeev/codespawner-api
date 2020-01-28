#! /usr/bin/env bash

for dir in code-runner hurdle root; do
  echo `pwd`/$dir
  cd $dir

  if [ ! -f conf.json ]; then
    echo "Created conf.json"
    cp config/conf.json.sample config/conf.json
  fi

  if [ -f nginx.conf.sample ] && [ ! -f nginx.conf ]; then
    echo "Created nginx.conf"
    cp nginx.conf.sample nginx.conf
  fi
  if [ -f .air.conf.sample ] && [ ! -f .air.conf ]; then
    echo "Created .air.conf"
    cp .air.conf.sample .air.conf
  fi
  cd ..
done

cp run-all.py.sample run-all.py
