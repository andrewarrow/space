{{ define "ticket" }}
  {{ $title := index . "title" }}
    td pr-3
      {{ $title }}
  {{ $description := index . "description" }}
    td pr-3
      {{ $description }}
  {{ $created_at := index . "created_at_ago" }}
    td pr-3 whitespace-nowrap
      {{ $created_at }}
  {{ $updated_at := index . "updated_at_ago" }}
    td pr-3 whitespace-nowrap
      {{ $updated_at }}
  {{ $id := index . "id" }}
    td pr-3
      {{ $id }}
  {{ $guid := index . "guid" }}
    td whitespace-nowrap
      {{ $guid }}
  {{ end }}
