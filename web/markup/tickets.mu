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
      {{ $title := index $item "title" }}
        td pr-3
          {{ $title }}
      {{ $description := index $item "description" }}
        td pr-3
          {{ $description }}
      {{ $created_at := index $item "created_at_ago" }}
        td pr-3 whitespace-nowrap
          {{ $created_at }}
      {{ $updated_at := index $item "updated_at_ago" }}
        td pr-3 whitespace-nowrap
          {{ $updated_at }}
      {{ $id := index $item "id" }}
        td pr-3
          {{ $id }}
      {{ $guid := index $item "guid" }}
        td whitespace-nowrap
          {{ $guid }}
    {{end}}
