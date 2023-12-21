{{ define "device_show" }}
  div id=device space-y-3 w-full flex flex-col justify-center items-center
    {{ range $i, $item := . }}
      div cursor-pointer border-2 border-black w-full md:w-1/2 text-center
        {{ $item }}
    {{ end }}
  {{ end }}
