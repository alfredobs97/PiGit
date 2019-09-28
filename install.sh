#!/bin/bash

sudo apt install -y git

sudo useradd git
read -p "Password for git user" pass 

(echo $pass ; echo $pass ) | passwd git

# Inicialize repos in /repos

sudo mkdir /repos/

echo "To init a repo-server cd into folder and write:"
echo "git init --bare"

echo "The repo will be made"
echo ""
echo "---------------------------------"
echo ""

echo "To add this repo to other location write:"
echo "git remote add $nameOfLocalRepository ssh://git@$IpRaspberry:$DirectoryOfRepository"

echo ""
echo "---------------------------------"
echo ""

echo "To clone this repository git clone ssh://git@$IpRaspberry:$DirectoryOfRepo"
echo ""


