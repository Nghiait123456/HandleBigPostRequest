RERUN=`which reflex1`
if [[ "$RERUN" == "" ]]; then
    echo "install reflex"
    go get -u github.com/cespare/reflex
fi

echo "running app..."
reflex -c ./reflex.conf