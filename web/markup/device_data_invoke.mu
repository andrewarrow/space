{{ define "device_data_invoke" }}
  div my-3
    {{ $item := index . "item" }}
    {{ $item }}
  div id=device-data space-y-3 w-full flex flex-col justify-center items-center
    div
      release_food(bay, grams)
  div space-y-3
    div
      bay
    div
      grams
    div
      execute
  {{ end }}
