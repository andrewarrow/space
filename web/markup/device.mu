{{ define "device" }}
  {{ $guid := index . "guid" }}
  {{ $name := index . "name" }}
  div id=d{{$guid}} cursor-pointer border-2 border-black w-full md:w-1/2 text-center
    {{ $name }}
  {{ end }}
