{{ define "device_show" }}
  div id=device space-y-3 w-full flex flex-col justify-center items-center
    {{ range $i, $item := . }}
      {{ $g := index $item "guid" }}
      div space-y-3 id=d{{$g}} text-sm cursor-pointer border-2 border-black w-full md:w-1/2 text-center
        {{ $f := index $item "function" }}
        {{ $params := index $item "params" }}
        div
          {{ $f }}
        form id=h{{$g}} hidden2 w-full 
          {{ range $i, $param := $params }}
            div flex space-x-3 w-full pb-6 items-center justify-left
              div w-1/4
                {{ $param }}
              div w-3/4
                input w-full type=input name=foo border-2 border-black
          {{ end }}
    {{ end }}
  div id=payloads mt-9 space-y-3 w-full flex flex-col justify-center items-center
    hi
  {{ end }}
