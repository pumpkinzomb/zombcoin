# Making zombcoin by golang

## How to start this.  
  
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
  