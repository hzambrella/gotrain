GOGO=`pwd`
cd ../../
mkdir pkg
export GOLIBS=`pwd`
cd -
export GOPATH=$GOLIBS:$GOGO
echo $GOPATH
