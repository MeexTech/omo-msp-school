.PHONY: proto
proto:
	protoc --proto_path=. --micro_out=. --go_out=. proto/school/common.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/school/classes.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/school/teacher.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/school/student.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/school/school.proto