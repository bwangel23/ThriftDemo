namespace go example

struct Data {
    1: string text,
}

service FormatData {
    Data doFormat(1: Data data)
}
