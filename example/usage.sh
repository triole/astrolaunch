#!/bin/bash

astrolaunch calc -d 20241112
astrolaunch calc -d 20241112 -r 3

astrolaunch exec -a sun.dawn -p 1m -q 2h ls -la /tmp

astrolaunch ops -f test
