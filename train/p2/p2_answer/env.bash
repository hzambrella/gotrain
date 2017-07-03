GOGO=`pwd`
cd ../../
mkdir pkg
GOLIBS=`pwd`
cd -
export GOPATH=$GOLIBS:$GOGO
echo $GOPATH
