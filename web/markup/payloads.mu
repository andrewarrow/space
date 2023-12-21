{{ define "payloads" }}
  {{ range $i, $item := . }}
    div text-xs
      {{ $ago := index $item "created_at_ago" }}
      {{ $ago }}
    div font-mono text-xs 
      {{ $p := index $item "payload" }}
      {{ jq $p }}
  {{ end }}
  {{ end }}
