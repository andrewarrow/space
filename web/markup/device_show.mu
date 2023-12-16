div p-3 
  div py-3 text-center font-bold
    {{ $name := index . "name" }}
    {{ $data := index . "data" }}
    {{ $name }}
  div space-y-3 w-full flex flex-col justify-center items-center
    div id=json space-y-3
      {{ range $k,$v := $data }}
        div
          div cursor-pointer id={{$k}}
            {{ $k }}
          div font-mono
            {{ $v }}
          {{ if hasImage $v }}
            div
              {{ $full_url_photo := $v }}
              img src={{$full_url_photo}}
          {{ end }}
      {{ end }}
