#include <iostream>
#include <pybind11/pybind11.h>
#include <bitcoin/bitcoin.hpp>
#include "libvarint.h"

namespace py = pybind11;

py::bytes encode_varint(const uint64_t value)
{
    bc::data_chunk data(10);
    GoSlice slice = { data.data(), 10, 10 };
    GoInt n = PutVarint(slice, value);
    data.resize(n);
    return py::bytes(
        reinterpret_cast<const char*>(data.data()), data.size());
}

uint64_t decode_varint(py::bytes data)
{
    std::string data_string = data;
    auto data_size = data_string.size();

    GoSlice slice2 = { data_string.data(), data_size, data_size };
    GoInt64 result = ReadVarint(slice2);
    return result;
}

PYBIND11_MODULE(pyvaruint, m)
{
    m.def("encode_varint", &encode_varint);
    m.def("decode_varint", &decode_varint);
}

