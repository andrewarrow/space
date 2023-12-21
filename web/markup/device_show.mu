{{ define "device_show" }}
  div id=device space-y-3 w-full flex flex-col justify-center items-center
    {{ range $i, $item := . }}
      {{ $g := index $item "guid" }}
      div id=d{{$g}} text-sm cursor-pointer border-2 border-black w-full md:w-1/2 text-center
        {{ $f := index $item "function" }}
        {{ $f }}
    {{ end }}
  div id=payloads mt-9 space-y-3 w-full flex flex-col justify-center items-center
    hi
  {{ end }}
