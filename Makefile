serve:
	@docker compose -f build/docker-compose.yml up -d

stop:
	@docker compose -f build/docker-compose.yml stop

buildup:
	@docker compose -f build/docker-compose.yml up -d --no-deps --build api-server

test-part1:
	cd _test/part1 && go test -v

test-part2:
	cd _test/part2 && go test -v

test-part3:
	cd _test/part3 && go test -v

test-part4:
	cd _test/part4 && go test -v

test-part5:
	cd _test/part5 && go test -v
