{{ define "device" }}
  {{ $id := index . "id" }}
  {{ $name := index . "name" }}
  div id=d{{$id}} cursor-pointer border-2 border-black w-full md:w-1/2 text-center
    {{ $name }}
  {{ end }}
