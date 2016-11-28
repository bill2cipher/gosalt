#!/bin/bash
CodeDir=$1
Version=$2
ReleaseDir=$3/${Version}

shift 3
Servers=$@
CodeDir=${CodeDir}/${Version}/Server

# check if CodeDir is a directory
if [ -f ${CodeDir} ] ; then
  printf "dir %s is not a directory, but a normal file\n" ${ReleaseDir}
  exit 1
fi

# make dir CodeDir or ReleaseDir if not exist
if [ ! -e ${CodeDir} ] ; then 
  mkdir -p ${CodeDir}
fi

if [ ! -e ${ReleaseDir} ] ; then
  mkdir -p ${CodeDir}
fi

# sync code
cd ${CodeDir}
p4 sync ./...
if [ ! $? -eq 0 ] ; then
  printf "p4 sync code failed"
	exit 0
fi

# test if it is all, if it is, convert to others
if [ ${#Servers[@]} -eq 1 -a ${Servers[0]} = 'all' ]; then
  Servers=('game' 'battle' 'gateway' 'eorm')
fi

function build_release() {
  cd ${CodeDir}/$1
  rm -rf _build

  # clean release dir content if exist
  if [ -e ${ReleaseDir}/$1 ] ; then
    rm -rf ${ReleaseDir}/$1
  fi 

  rebar3 as prod release

  # copy new release content into target
  if [ $? -eq 0 ] ; then
    cp -r _build/prod/rel/$1 ${ReleaseDir}
  else
    printf "make release for %s failed, exiting\n" $1
    exit 1
  fi
}

function tar_release() {
  cd ${ReleaseDir}
  tar -czf mgame_${Version}_`date "+%Y%m%d%H%M%S"`.tar.gz ${Servers}
}

for SvrName in ${Servers[@]}; do
  build_release ${SvrName}
done
tar_release
