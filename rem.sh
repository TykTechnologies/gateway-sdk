echo "hello"
cd gate || exit
[ ! -e go.mod ] || rm go.mod
[ ! -e go.sum ] || rm go.sum