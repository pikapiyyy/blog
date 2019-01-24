#!/bin/bash
app=blog
p=$(pidof $app)
if [ $p ]
then
    echo "kill $app pid $p"
    kill -9 $p
fi
echo "$app building..."
go build
echo "$app start..."
nohup ./$app > nohup.out 2>&1 &
echo "new $app is running"
echo "end...."
exit
