div p-3 
  div py-3 text-center
    Device
  div space-y-3 w-full flex flex-col justify-center items-center
    div id=json
      {{ range $k,$v := . }}
        div cursor-pointer id={{$k}}
          {{ $k }}
      {{ end }}
