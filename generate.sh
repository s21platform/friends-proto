protoc  --go_out=./  \
        --go-grpc_out=./ \
        friends.proto

protoc --doc_out=. --doc_opt=markdown,README.md ./friends.proto
