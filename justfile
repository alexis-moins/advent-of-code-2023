day-one:
    @-go run cmd/day_one/day_one.go

test-day-one:
    @-go test ./test/day_one -v

day-two:
    @-go run cmd/day_two/day_two.go

test-day-two:
    @-go test ./test/day_two -v

day-three:
    @-go run cmd/day_three/day_three.go

test-day-three:
    @-go test ./test/day_three -v
