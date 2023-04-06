alias x := xds
alias s := server
alias c := client

xds +args:
    #!/bin/zsh
    set -eux
    cd xds
    just {{args}}

server +args:
    #!/bin/zsh
    set -eux
    cd server
    just {{args}}

client +args:
    #!/bin/zsh
    set -eux
    cd client
    just {{args}}