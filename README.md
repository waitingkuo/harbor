Harbor: the easier way to install docker things
===============================================

Harbor provides a easier way to install docker & its ecosystem. It's under developed, we currently only support docker-machine & docker-compose. Any feedback is wellcome.


# Installation

    go install github.com/waitingkuo/harbor
    
Harbor will install binary file to ~/.harbor/bin , to add it into `$PATH`, add following to your `~/.bashrc`

    export PATH=~/.harbor/bin:$PATH

### Install pacages

    harbor install docker-machine
    harbor install docker-compose

    
## Todo

1. update packaeg
2. find a way to update the software version
3. remove command
4. list command to show all installed packages
5. search command to search packages
