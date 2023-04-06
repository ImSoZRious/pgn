alias x := xds
alias s := server
alias c := client

xds +args:
    #!/bin/zsh
    set -eu
    cd xds
    just {{args}}

server +args:
    #!/bin/zsh
    set -eu
    cd server
    just {{args}}

client +args:
    #!/bin/zsh
    set -eu
    cd client
    just {{args}}

k3d +args:
    #!/bin/zsh
    set -eu
    cd k3d
    just {{args}}

build:
    #!/bin/bash
    set -eu
    just client build
    just server build
    just xds build

deploy:
    #!/bin/bash
    set -eu
    just client deploy
    just server deploy
    just xds deploy