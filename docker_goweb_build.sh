sudo docker rmi -f goweb
sudo docker container prune
sudo docker build . -t goweb
