#!/bin/bash

sudo docker stop test_sagara
sudo docker rm test_sagara
sudo docker rmi test_sagara
sudo docker build -t test_sagara .
sudo docker run -tid -p 8800:8800 --network="host" --name test_sagara test_sagara ./MainEndPoint