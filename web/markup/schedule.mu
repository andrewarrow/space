{{ define "schedule" }}
  {{ $guid := index . "guid" }}
  {{ $name := index . "name" }}
  div flex w-full id=w{{$guid}}
    div w-full
      {{ $name }}
    div mx-auto mr-9
      a href=# id=x{{$guid}}
        X
  {{ end }}
