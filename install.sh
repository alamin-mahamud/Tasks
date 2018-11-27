echo "changing directory to Tasks"
cd $GOPATH/src/github.com/alamin-mahamud/go-tasks

# echo "creating table"
# cat schema.sql | sqlite3 tasks.db

echo "building the go binary"
mkdir -p build
go build -o build/tasks

echo "starting the binary"
./build/tasks
