div p-3
  {{ $items := index . "items" }}
  table
    {{ range $i, $item := $items }}
      {{ $title := index $item "title" }}
      tr
        td
          {{ $title }}
    {{ end }}
