namespace go RPCThrift.thrift.RPCThrift

service RPCThrift{
    i32 Add(1:i32 addStart, 2:i32 addEnd);
}