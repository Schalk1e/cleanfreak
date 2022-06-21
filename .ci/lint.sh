DIR=$PWD

if [ ! -d $DIR/cmd ] 
then
    echo "Please run lint.sh from root :)"
    exit 1
else
    go mod verify
    go mod download
    go mod tidy
    golangci-lint run $DIR/cmd $DIR/core $DIR/cmdutil
fi