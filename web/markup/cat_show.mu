{{ define "cat_show" }}
  div id=devices space-y-3 w-full flex flex-col justify-center items-center
    {{ range $i, $item := . }}
      {{ $guid := index $item "guid" }}
      {{ $name := index $item "name" }}
      {{ $ago := index $item "last_connection_at_ago" }}
      div a1 border-2 border-black w-full md:w-1/2 text-center id=w{{$guid}}
        div 
          {{ $name }}
        div a2 flex items-center justify-center
          div pl-3 data-click cursor-pointer id=d{{$guid}}
            data
          div w-full text-xs
            {{ $ago }}
          div ml-auto pr-3 cat-click cursor-pointer id=c{{$guid}}
            category
    {{ end }}
  {{ end }}
