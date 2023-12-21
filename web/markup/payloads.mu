{{ define "payloads" }}
  {{ range $i, $item := . }}
    div text-xs
      {{ $ago := index $item "created_at_ago" }}
      {{ $ago }}
    div font-mono text-xs 
      {{ $p := index $item "payload" }}
      {{ $full_url_photo := index $p "photo" }}
      {{ jq $p }}
      {{ if $full_url_photo }}
        img src={{$full_url_photo}}
      {{ end }}
  {{ end }}
  {{ end }}
