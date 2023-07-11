namespace go base

struct NilResponse {}

struct CommomResponse {
    1: i32 code
    2: string message
}