{{ define "schedules" }}
  div id=schedule-list space-y-3 w-full flex flex-col justify-center items-center
    form space-y-3 id=add-schedule-form
      div flex space-x-3
        div
          Every day at
        div
          select id=value
            option value=9 
              9AM
            option value=10
              10AM
      div flex space-x-3 bg-r items-center
        div
          input type=text id=name placeholder=name
        div text-xs
          input type=submit value=Add border rounded bg-blue-600 text-white py-2 px-2
    {{ range $i, $item := . }}
      {{ $guid := index $item "guid" }}
      {{ $name := index $item "name" }}
      div flex w-full id=g{{$guid}}
        div w-full
          {{ $name }}
        div mx-auto mr-9
          x
    {{ end }}
  {{ end }}
