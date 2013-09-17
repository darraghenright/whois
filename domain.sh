#!/bin/bash

less domain.txt | xargs -n 1 -P 32 sh -c 'whois -c IE $1.ie | tail +5 > whois/$1' sh