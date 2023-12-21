{{ define "payloads" }}
  {{ range $i, $item := . }}
    div 
      {{ $p := index $item "payload" }}
      {{ $p }}
  {{ end }}
  {{ end }}
