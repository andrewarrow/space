{{ define "device_data_show" }}
  div my-3
    {{ $item := index . "item" }}
    {{ $item }}
  div id=device-data space-y-3 w-full flex flex-col justify-center items-center
    data
  {{ end }}
