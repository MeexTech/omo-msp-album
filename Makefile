.PHONY: proto
proto:
	protoc --proto_path=. --micro_out=. --go_out=. proto/album/common.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/album/album.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/album/collective.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/album/exhibit.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/album/panorama.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/album/photocopy.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/album/style.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/album/frame.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/album/composition.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/album/page.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/album/sheet.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/album/certificate.proto
