#! /usr/bin/env bash

for dir in code-runner hurdle root; do
  echo `pwd`/$dir
  cd $dir

  if [ ! -f conf.json ]; then
    echo "Created conf.json..."
    cp conf.json.sample conf.json
  fi

  if [ -f nginx.conf.sample ] && [ ! -f nginx.conf ]; then
    echo "Created nginx.conf..."
    cp nginx.conf.sample nginx.conf
  fi

  cd ..
done

cp run-all.py.sample run-all.py
cd static/web/scripts && npm i && cd -
