{{ define "cat_show" }}
  div id=devices space-y-3 w-full flex flex-col justify-center items-center
    {{ range $i, $item := . }}
      {{ $guid := index $item "guid" }}
      {{ $name := index $item "name" }}
      div id=d{{$guid}} cursor-pointer border-2 border-black w-full md:w-1/2 text-center
        div
          {{ $name }}
        div flex items-center justify-center
          div pl-3
            data
          div w-full 
            31 days ago
          div ml-auto pr-3
            category
    {{ end }}
  {{ end }}
