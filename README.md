# Making zombcoin by golang

## How to start this.  

This repository is based on nomadcoder's lectures. If you want to find the course, you can find [here](https://nomadcoders.co/nomadcoin).  
   
- Install Go
  
```
# download the latest version
wget https://go.dev/dl/go1.19.2.linux-amd64.tar.gz

# remove old version (if any)
sudo rm -rf /usr/local/go

# install the new version
sudo tar -C /usr/local -xzf go1.19.2.linux-amd64.tar.gz
```
  
- Configure Environmental Variables
  
```
# run these commands
cat <<EOF >> ~/.profile
export GOROOT=/usr/local/go
export GOPATH=$HOME/go
export GO111MODULE=on
export PATH=$PATH:/usr/local/go/bin:$HOME/go/bin
EOF

source ~/.profile

go version
```
  
- Git clone this repository and install dependencies by go module.  

## Steps of Course  
- I followed steps with commit messages  

1. Make an module of blockchain
2. Build a explorer server for viewing blockchain's data
3. Pre-Build a rest API server for communicate with blockchain modules. (Use gorilla-mux dependency)
4. Make for command line interface for api server & explorer
