div p-3
  {{ $items := index . "items" }}
  table
    tr font-bold
      td pr-3
        Title
      td pr-3
        description
      td pr-3
        created_at
      td pr-3
        updated_at
      td pr-3
        id
      td pr-3
        guid
    {{ range $i, $item := $items }}
      tr 
        {{ template "ticket" $item }}
    {{end}}
