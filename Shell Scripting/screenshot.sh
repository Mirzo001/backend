#! /usr/bin/bash

# scrot -u screen 



# read -p  "Enter your username: " username1

trap printout SIGINT
printout() {
    echo ""
    echo "Taken = $count screenshots"
    exit
}

read  -p  "Enter sleep time: " tm

# mkdir "$(date +"%Y-%m-%d")"

mkdir -p data
cd data
# cd "$(date +"%Y-%m-%d")"
while [ 1 ]
do
    mkdir -p "$(date +"%Y-%m-%d")"
   scrot ./$(date +"%Y-%m-%d")/"Screen Shot  `date +%Y-%m-%d" at "%H.%M.%S`.jpg" --delay $tm -c 
   ((count++))
done

# cd /mnt/c/Users/user/Desktop/Shelltutorial/

