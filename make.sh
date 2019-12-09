go build -o libvarint.so -buildmode=c-shared varint.go
g++ -O3 -Wall -shared -std=c++2a -fPIC `python3 -m pybind11 --includes` test_varint.cpp -o pyvaruint`python3-config --extension-suffix` $(pkg-config --cflags --libs libbitcoin) -L. -lvarint

