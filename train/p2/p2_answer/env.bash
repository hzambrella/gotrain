GOGO=`pwd`
cd ../../
export GOLIBS="$(dirname `pwd`)/pkg"
mkdir -p $GOLIBS/src
cd -
echo $GOLIBS
export GOPATH=$GOLIBS:$GOGO
echo $GOPATH
