{{ define "device_show" }}
  div my-3
    {{ $item := index . "item" }}
    {{ $cats := index . "cats" }}
    {{ $item }}
  div id=device-cats space-y-3 w-full flex flex-col justify-center items-center
    {{ range $i, $item := $cats }}
      div id=cat-{{$item}} cursor-pointer border-2 border-black w-full md:w-1/2 text-center
        {{ $item }}
    {{ end }}
  {{ end }}
