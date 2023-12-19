{{ define "schedule" }}
  {{ $guid := index . "guid" }}
  {{ $name := index . "name" }}
  div flex w-full id=w{{$guid}} items-center
    div w-full mb-3
      div
        {{ $name }}
      div text-gray-400
        Every day at 09:00
    div mx-auto mr-9
      a href=# id=x{{$guid}}
        X
  {{ end }}
