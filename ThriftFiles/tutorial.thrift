include "shared.thrift"

namespace go tutorial

struct Data {
    1: string text,
}

service FormatData {
    Data doFormat(1: Data data)
}
