echo "hello"
cd apim || exit
[ ! -e go.mod ] || rm go.mod
[ ! -e go.sum ] || rm go.sum