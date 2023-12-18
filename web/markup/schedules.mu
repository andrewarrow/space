{{ define "schedules" }}
  div id=schedule-list space-y-3 w-full flex flex-col justify-center items-center
    {{ range $i, $item := . }}
      div
        {{ $item }}
    {{ end }}
  {{ end }}
