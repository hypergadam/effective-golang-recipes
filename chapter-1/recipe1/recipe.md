## Using In-Memory Readers and Writers to Support `[]byte`

Keep the existing stream-based serialization and adapt the input and output:

- `bytes.NewReader(data)` wraps a `[]byte` in an `io.Reader`, so `UnmarshalRide` can reuse `DecodeRide`.
- `bytes.Buffer` implements `io.Writer`. Pass `&buf` to the encoder, then return `buf.Bytes()` from `MarshalRide`.

`Decoder` and `Encoder` wrap the standard JSON types. Construct those types with
`json.NewDecoder` and `json.NewEncoder`; their internal fields are private.

`json.Encoder.Encode` adds a trailing newline, which the decoder accepts.
Encoding and decoding errors are returned to the caller.

These adapters prepare bytes for a shared-memory transport; they do not allocate
or synchronize shared memory between processes.
