#!/bin/sh

WORK_PATH=../
echo "work path:"$WORK_PATH

cd $WORK_PATH

truffle test --network gethtest
if [ $? -ne 0 ] 
then
    echo "truffle test failed."
    exit
fi

truffle migrate --reset --network gethtest
