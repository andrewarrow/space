{{ define "cat" }}
  {{ $count := index . "count" }}
  {{ $name := index . "name" }}
  {{ $id := index . "id" }}
  div id={{$id}} cursor-pointer border-2 border-black w-full md:w-1/2 text-center
    {{ $name }} {{ $count }}
  {{ end }}
