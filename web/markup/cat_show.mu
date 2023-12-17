{{ define "cat_show" }}
  div id=devices space-y-3 w-full flex flex-col justify-center items-center
    {{ range $i, $item := . }}
      {{ $id := index $item "id" }}
      {{ $name := index $item "name" }}
      div id={{$name}} cursor-pointer border-2 border-black w-full md:w-1/2 text-center
        {{ $name }}
    {{ end }}
  {{ end }}
