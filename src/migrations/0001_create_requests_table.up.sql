CREATE TABLE IF NOT EXISTS requests (
   user_id Nullable(UInt64),
   request_from Nullable(String),
   request_to Nullable(String),
   request_method Nullable(String),
   request_body Nullable(String),
   response_body Nullable(String),
   response_code Nullable(UInt32),
   response_time Nullable(Int64),
   created_at DateTime('UTC') default now()
)
ENGINE = MergeTree
ORDER BY created_at;