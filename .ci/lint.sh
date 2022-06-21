DIR=$PWD

if [ ! -d $DIR/cmd ] 
then
    echo "Please run lint.sh from root :)"
    exit 1
else
    golangci-lint run $DIR/cmd $DIR/core $DIR/cmdutil
fi