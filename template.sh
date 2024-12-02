#!/bin/bash

DIR=`printf d%02d $1`

mkdir "${DIR}"
touch "${DIR}/main.go"
cat > "${DIR}/main.go" << EOF
package main

func main() {

}
EOF

touch "${DIR}/main_test.go"
cat > "${DIR}/main_test.go" << EOF
package main
EOF

touch "${DIR}/input.txt"
touch "${DIR}/example1.txt"
