#!/bin/bash

CodeDir=$1
Version=$2
ReleaseDir=$3/${Version}

echo "sync server code starting..."
echo "sync server code finished"

echo "build server code starting..."
dd if=/dev/zero of=/tmp/test.dat bs=512 count=1024
if [ ! $? -eq 0 ] ; then
  echo "build server code failed"
  exit 1
fi
echo "build server code finished"

echo "build start script starting..."
echo "${ReleaseDir}/start.sh"
echo "echo 'starting servers'" >> ${ReleaseDir}/start.sh
echo "build start script finished"

echo "tar servers starting..."
dest=mgame_${Version}_`date "+%Y%m%d%H%M%S"`.tar
cd /tmp
tar -cvf ${dest} test.dat
mv ${dest} ${ReleaseDir}
cd ${ReleaseDir}
tar -r -f ${dest} start.sh
echo "tar servers finished"
