{{ define "schedules" }}
  div space-y-3 w-full flex flex-col justify-center items-center
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
      div flex space-x-3 items-center
        div
          input type=text id=name placeholder=name border-2 border-grey-400 p-1 rounded
        div text-xs
          input type=submit value=Add border rounded bg-blue-600 text-white py-2 px-2
    div id=schedule-list w-full
      {{ . }}
  {{ end }}
