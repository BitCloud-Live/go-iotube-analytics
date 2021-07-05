tvl = from(bucket: "my-bucket")
  |> range(start: v.timeRangeStart)
  |> filter(fn: (r) => r["_measurement"] == "tvl")
  |> filter(fn: (r) => r["_field"] == "tvl")
  |> filter(fn: (r) => r["network"] == "ethereum")

r0 = from(bucket: "my-bucket")
  |> range(start: v.timeRangeStart)
  |> filter(fn: (r) => r["_measurement"] == "price")
  |> last()
  |> tableFind(fn: (key) => key._field == "price")
  |> getRecord(idx: 0)


tvl
  |> map(fn: (r) => ({
      _time: r._time,
      _value: r._value * r0._value,
      symbol: r.symbol
    })
  )