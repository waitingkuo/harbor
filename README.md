Harbor: the easier way to install docker things
===============================================


# Installation

    go install github.com/waitingkuo/harbor
    
Harbor will install binary file to ~/.harbor/bin , to add it into `$PATH`, add following to your `~/.bashrc`

    export PATH=~/.harbor/bin:$PATH

### Install pacages

    harbor install docker-machine
    harbor install docker-compose

    
