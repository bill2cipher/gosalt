#!/usr/bin/env bash
Server=$1
Source=$2
Dest=$3

salt "${Server}" "${Source}" "${Dest}"
