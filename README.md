# goimportstidy

This tool updates your Go import lines, grouping it into three groups: 
 - stdlib,
 - external libraries,
 - local libraries (optional).
 
Installation: 

     $ go install github.com/kwinata/goimportstidy@latest
     
Usage:

    $ goimportstidy -w -local github.com/shipwallet main.go -current github.com/shipwallet/core .
