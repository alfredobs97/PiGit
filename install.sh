#!/bin/bash

sudo useradd git
read -p "Password for git user" pass 

(echo $pass ; echo $pass ) | passwd git

# Inicialize repos in /repos

sudo mkdir /repos/

echo "To add this repo to other location write:"
echo "git remote add $nameOfLocalRepository ssh://git@$IpRaspberry:$DirectoryOfRepository"

echo ""
echo "---------------------------------"
echo ""

echo "To clone this repository git clone ssh://git@$IpRaspberry:$DirectoryOfRepo"
echo ""


