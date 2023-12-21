{{ define "payloads" }}
  {{ range $i, $item := . }}
    div font-mono text-xs 
      {{ $p := index $item "payload" }}
      {{ jq $p }}
  {{ end }}
  {{ end }}
