div p-3
  {{ $items := index . "items" }}
  table
    tr font-bold
      td pr-3
        title
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
{{ define "ticket" }}
  {{ $title := index . "title" }}
    td pr-3 whitespace-nowrap
      {{ $title }}
  {{ $description := index . "description" }}
    td pr-3 whitespace-nowrap
      {{ $description }}
  {{ $created_at := index . "created_at" }}
    td pr-3 whitespace-nowrap
      {{ $created_at }}
  {{ $updated_at := index . "updated_at" }}
    td pr-3 whitespace-nowrap
      {{ $updated_at }}
  {{ $id := index . "id" }}
    td pr-3 whitespace-nowrap
      {{ $id }}
  {{ $guid := index . "guid" }}
    td pr-3 whitespace-nowrap
      {{ $guid }}
  {{end}}
