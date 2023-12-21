{{ define "payloads" }}
  {{ range $i, $item := . }}
    div 
      {{ $p := index $item "payload" }}
      {{ jq $p }}
  {{ end }}
  {{ end }}
