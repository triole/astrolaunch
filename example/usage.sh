#!/bin/bash

# calc diff and print
astrolaunch calc -d 20241112
astrolaunch calc -d 20241112 -r 3

# execute command given as cli arg
astrolaunch exec -a sun.dawn -p 1m -q 2h ls -la /tmp
astrolaunch exec --log-level debug -w -a sun.dawn -p 1m -q 1m ls -la /tmp

# process config file and run operations defined in there
astrolaunch ops -f test
